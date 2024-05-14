package gpubsub

import (
	"context"
	"encoding/json"

	"110yards.ca/libs/go/core/eventbus"
	"cloud.google.com/go/pubsub"
)

type publisher struct {
	topic  *pubsub.Topic
	client *pubsub.Client
}

// Topics have goroutines so we want to avoid re-initializing them if we are publishing to them
var publisherInstance publisher

func NewPublisher(topicName string) eventbus.Publisher {
	if publisherInstance.client != nil {
		return &publisherInstance
	}

	topic := client.Topic(topicName)

	return &publisher{
		topic:  topic,
		client: client,
	}
}

// Publish implements Publisher.
func (g *publisher) Publish(message interface{}) error {
	data, err := json.Marshal(message)

	if err != nil {
		return err
	}

	g.topic.Publish(context.Background(), &pubsub.Message{
		Data: data,
	})

	return nil
}
