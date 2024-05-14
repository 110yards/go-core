package gpubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

var client *pubsub.Client
var projectId string

func InitializeGooglePublisher(gcloudProjectId string) error {
	c, err := pubsub.NewClient(context.Background(), gcloudProjectId)

	if err != nil {
		return err
	}

	client = c
	projectId = gcloudProjectId

	return nil
}
