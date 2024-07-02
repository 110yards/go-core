package eventbus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"110yards.ca/libs/go/core/logger"
)

type VirtualPublisher struct {
	topicName  string
	pushTarget string
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
	j, err := json.Marshal(message)
	if err != nil {
		return err
	}

	if v.pushTarget != "" {
		// use http client to POST to target
		_, err := http.NewRequest("POST", v.pushTarget, nil)
		if err != nil {
			return err
		}
	} else {
		logger.Infof("Published virtual message %s (no push target set)", string(j))
	}

	return nil
}

func NewVirtualPublisher(topicName string) Publisher {
	virtualPushKey := fmt.Sprintf("VIRTUAL_PUSH_%s", topicName)
	pushTarget := os.Getenv(virtualPushKey)

	if pushTarget == "" {
		logger.Warnf("No virtual push target found for topic %s, configure %s to enable push in dev", topicName, virtualPushKey)
	}

	return &VirtualPublisher{topicName: topicName, pushTarget: pushTarget}
}
