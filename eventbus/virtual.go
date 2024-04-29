package eventbus

import (
	"encoding/json"

	"110yards.ca/libs/go/core/logger"
)

type VirtualPublisher struct {
	topicName string
}

// CreateSubscription implements Publisher.
func (v *VirtualPublisher) CreateSubscription(subscriptionName string, topicName string, endpoint string, config SubscriptionConfig) error {
	logger.Infof("Created virtual subscription %s on topic %s to endpoint %s", subscriptionName, topicName, endpoint)
	return nil
}

// CreateTopic implements Publisher.
func (v *VirtualPublisher) CreateTopic(topicName string) error {
	logger.Infof("Created virtual topic %s", topicName)
	return nil
}

// Publish implements Publisher.
func (v *VirtualPublisher) Publish(message interface{}) error {
	// TODO: repo for messages
	j, err := json.Marshal(message)
	if err != nil {
		return err
	}
	logger.Infof("Published message %s", string(j))

	return nil
}

func NewVirtualPublisher(topicName string) Publisher {
	return &VirtualPublisher{topicName: topicName}
}
