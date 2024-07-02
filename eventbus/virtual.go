package eventbus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

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

	if v.pushTarget != "" {

	} else {
		j, err := json.Marshal(message)
		if err != nil {
			return err
		}
		logger.Infof("Published virtual message %s (no push target set)", string(j))
	}

	return nil
}

func (v *VirtualPublisher) PushMessage(message interface{}) error {
	type PubsubMessage struct {
		Message struct {
			Data interface{} `json:"data"`
		} `json:"message"`
	}

	payload := PubsubMessage{
		Message: struct {
			Data interface{} `json:"data"`
		}{
			Data: message,
		},
	}

	// use http client to POST to target
	body, err := json.Marshal(payload)
	request, err := http.NewRequest("POST", v.pushTarget, strings.NewReader(string(body)))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}
}

func NewVirtualPublisher(topicName string) Publisher {
	virtualPushKey := fmt.Sprintf("VIRTUAL_PUSH_%s", topicName)

	// replace - with _ in key and convert to upper case
	virtualPushKey = strings.ToUpper(strings.Replace(virtualPushKey, "-", "_", -1))

	pushTarget := os.Getenv(virtualPushKey)

	if pushTarget == "" {
		logger.Warnf("No virtual push target found for topic %s, configure %s to enable push in dev", topicName, virtualPushKey)
	} else {
		logger.Infof("Virtual push target for topic %s is %s", topicName, pushTarget)
	}

	return &VirtualPublisher{topicName: topicName, pushTarget: pushTarget}
}
