package async

type ChannelProducer interface {

	Produce(channel chan<- interface{}) error
	Close() error
}

type ChannelConsumer interface {

	Consume(channel <- chan interface{}) error
	Close() error
}

type ChannelRegistryItem struct {

	Channel chan interface{}
	Consumers ChannelConsumer
	Producers []ChannelProducer
}

type ChannelRegistry struct {

	Items []ChannelRegistryItem
}

func (r *ChannelRegistry) Close() {

}