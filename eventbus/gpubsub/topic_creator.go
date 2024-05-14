package gpubsub

import (
	"context"

	"110yards.ca/libs/go/core/eventbus"
	"110yards.ca/libs/go/core/logger"
)

type topicCreator struct {
}

func NewTopicCreator() eventbus.TopicCreator {
	return &topicCreator{}
}

// CreateTopic implements eventbus.TopicCreator.
func (t *topicCreator) CreateTopic(topicName string) error {
	topic := client.Topic(topicName)
	exists, err := topic.Exists(context.Background())

	if err != nil {
		return err
	}

	if exists {
		logger.Warnf("topic %s already exists", topicName)
		return nil
	}

	_, err = client.CreateTopic(context.Background(), topicName)

	return err
}
