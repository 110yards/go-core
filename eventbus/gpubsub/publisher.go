package gpubsub

import (
	"context"
	"encoding/json"

	"110yards.ca/libs/go/core/eventbus"
	"110yards.ca/libs/go/core/logger"
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

	r := g.topic.Publish(context.Background(), &pubsub.Message{
		Data: data,
	})

	go func() error {
		_, err := r.Get(context.Background())
		if err != nil {
			logger.Errorf("Could not publish to topic %s: %s", g.topic, err)
		}
		logger.Infof("Published message to topic %s", g.topic)
		return err
	}()

	return nil
}
