package server

import (
	"net"
	"time"

	"github.com/unistack-org/micro/v3/metadata"
	"github.com/unistack-org/micro/v3/registry"
	"github.com/unistack-org/micro/v3/util/addr"
	"github.com/unistack-org/micro/v3/util/backoff"
)

var (
	// DefaultRegisterFunc uses backoff to register service
	DefaultRegisterFunc = func(service *registry.Service, config Options) error {
		var err error

		opts := []registry.RegisterOption{
			registry.RegisterTTL(config.RegisterTTL),
			registry.RegisterDomain(config.Namespace),
		}

		for i := 0; i <= config.RegisterAttempts; i++ {
			err = config.Registry.Register(service, opts...)
			if err == nil {
				break
			}
			// backoff then retry
			time.Sleep(backoff.Do(i + 1))
			continue
		}
		return err
	}
	// DefaultDeregisterFunc uses backoff to deregister service
	DefaultDeregisterFunc = func(service *registry.Service, config Options) error {
		var err error

		opts := []registry.DeregisterOption{
			registry.DeregisterDomain(config.Namespace),
		}

		for i := 0; i <= config.DeregisterAttempts; i++ {
			err = config.Registry.Deregister(service, opts...)
			if err == nil {
				break
			}
			// backoff then retry
			time.Sleep(backoff.Do(i + 1))
			continue
		}
		return err
	}
)

func NewRegistryService(s Server) (*registry.Service, error) {
	opts := s.Options()

	advt := opts.Address
	if len(opts.Advertise) > 0 {
		advt = opts.Advertise
	}

	host, port, err := net.SplitHostPort(advt)
	if err != nil {
		return nil, err
	}

	addr, err := addr.Extract(host)
	if err != nil {
		addr = host
	}

	node := &registry.Node{
		Id:      opts.Name + "-" + opts.Id,
		Address: net.JoinHostPort(addr, port),
	}
	node.Metadata = metadata.Copy(opts.Metadata)

	node.Metadata["server"] = s.String()
	node.Metadata["broker"] = opts.Broker.String()
	node.Metadata["registry"] = opts.Registry.String()

	return &registry.Service{
		Name:     opts.Name,
		Version:  opts.Version,
		Nodes:    []*registry.Node{node},
		Metadata: metadata.New(0),
	}, nil
}