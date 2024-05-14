package eventbus

import (
	"110yards.ca/libs/go/core/logger"
)

var publishers = make(map[string]Publisher)

func AddPublisher(topicName string, publisher Publisher) {
	publishers[topicName] = publisher
}

func InitializeDevPublishers(topicNames []string) error {

	for _, topicName := range topicNames {
		publisher := NewVirtualPublisher(topicName)
		AddPublisher(topicName, publisher)
	}

	return nil
}

func GetPublisher(topicName string) Publisher {
	publisher, exists := publishers[topicName]

	if !exists {
		logger.Warnf("publisher for topic %s not found", topicName)
		return nil
	}

	return publisher
}
