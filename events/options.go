package events

import (
	"time"

	"github.com/unistack-org/micro/v3/metadata"
)

// PublishOptions contains all the options which can be provided when publishing an event
type PublishOptions struct {
	// Metadata contains any keys which can be used to query the data, for example a customer id
	Metadata metadata.Metadata
	// Timestamp to set for the event, if the timestamp is a zero value, the current time will be used
	Timestamp time.Time
}

// PublishOption sets attributes on PublishOptions
type PublishOption func(o *PublishOptions)

// WithMetadata sets the Metadata field on PublishOptions
func WithMetadata(md metadata.Metadata) PublishOption {
	return func(o *PublishOptions) {
		o.Metadata = metadata.Copy(md)
	}
}

// WithTimestamp sets the timestamp field on PublishOptions
func WithTimestamp(t time.Time) PublishOption {
	return func(o *PublishOptions) {
		o.Timestamp = t
	}
}

// SubscribeOptions contains all the options which can be provided when subscribing to a topic
type SubscribeOptions struct {
	// Queue is the name of the subscribers queue, if two subscribers have the same queue the message
	// should only be published to one of them
	Queue string
	// StartAtTime is the time from which the messages should be consumed from. If not provided then
	// the messages will be consumed starting from the moment the Subscription starts.
	StartAtTime time.Time
}

// SubscribeOption sets attributes on SubscribeOptions
type SubscribeOption func(o *SubscribeOptions)

// WithQueue sets the Queue fielf on SubscribeOptions to the value provided
func WithQueue(q string) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Queue = q
	}
}

// WithStartAtTime sets the StartAtTime field on SubscribeOptions to the value provided
func WithStartAtTime(t time.Time) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.StartAtTime = t
	}
}

// WriteOptions contains all the options which can be provided when writing an event to a store
type WriteOptions struct {
	// TTL is the duration the event should be recorded for, a zero value TTL indicates the event should
	// be stored indefinitely
	TTL time.Duration
}

// WriteOption sets attributes on WriteOptions
type WriteOption func(o *WriteOptions)

// WithTTL sets the TTL attribute on WriteOptions
func WithTTL(d time.Duration) WriteOption {
	return func(o *WriteOptions) {
		o.TTL = d
	}
}

// ReadOptions contains all the options which can be provided when reading events from a store
type ReadOptions struct {
	// Topic to read events from, if no topic is provided events from all topics will be returned
	Topic string
	// Query to filter the results using. The store will query the metadata provided when the event
	// was written to the store
	Query map[string]string
	// Limit the number of results to return
	Limit int
	// Offset the results by this number, useful for paginated queries
	Offset int
}

// ReadOption sets attributes on ReadOptions
type ReadOption func(o *ReadOptions)

// ReadTopic sets the topic attribute on ReadOptions
func ReadTopic(t string) ReadOption {
	return func(o *ReadOptions) {
		o.Topic = t
	}
}

// ReadFilter sets a key and value in the query
func ReadFilter(key, value string) ReadOption {
	return func(o *ReadOptions) {
		if o.Query == nil {
			o.Query = map[string]string{key: value}
		} else {
			o.Query[key] = value
		}
	}
}

// ReadLimit sets the limit attribute on ReadOptions
func ReadLimit(l int) ReadOption {
	return func(o *ReadOptions) {
		o.Limit = 1
	}
}

// ReadOffset sets the offset attribute on ReadOptions
func ReadOffset(l int) ReadOption {
	return func(o *ReadOptions) {
		o.Offset = 1
	}
}
