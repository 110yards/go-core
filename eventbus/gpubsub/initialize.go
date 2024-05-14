package gpubsub

import (
	"context"

	"110yards.ca/libs/go/core/eventbus"
	"cloud.google.com/go/pubsub"
)

var client *pubsub.Client
var projectId string

func InitializeGoogle(gcloudProjectId string) error {
	c, err := pubsub.NewClient(context.Background(), gcloudProjectId)

	if err != nil {
		return err
	}

	client = c
	projectId = gcloudProjectId

	return nil
}

func InitializePublishers(isDev bool, gcloudProjectId string, topicNames []string) error {
	if isDev {
		return eventbus.InitializeDevPublishers(topicNames)
	}

	if client == nil {
		err := InitializeGoogle(gcloudProjectId)

		if err != nil {
			return err
		}
	}

	for _, topicName := range topicNames {
		publisher := NewPublisher(topicName)
		eventbus.AddPublisher(topicName, publisher)
	}

	return nil

}
