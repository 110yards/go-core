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
	topicName   string
	pushTargets []string
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

	if len(v.pushTargets) > 0 {
		for _, target := range v.pushTargets {
			err := pushMessage(target, message)
			if err != nil {
				return err
			}
		}
	} else {
		j, err := json.Marshal(message)
		if err != nil {
			return err
		}
		logger.Infof("Published virtual message to %s: %s (no push target set)", v.topicName, string(j))
	}

	return nil
}

func pushMessage(target string, message interface{}) error {
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

	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", target, strings.NewReader(string(body)))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(request)

	return err
}

func NewVirtualPublisher(topicName string) Publisher {
	virtualPushKey := fmt.Sprintf("VIRTUAL_PUSH_%s", topicName)
	// replace - with _ in key and convert to upper case
	virtualPushKey = strings.ToUpper(strings.Replace(virtualPushKey, "-", "_", -1))

	pushTarget := os.Getenv(virtualPushKey)

	// split by comma
	if pushTarget != "" {
		pushTargets := strings.Split(pushTarget, ",")
		logger.Infof("Virtual push targets for topic %s are %s", topicName, pushTargets)
		return &VirtualPublisher{topicName: topicName, pushTargets: pushTargets}
	}

	return &VirtualPublisher{topicName: topicName, pushTargets: []string{}}
}
