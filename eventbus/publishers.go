package eventbus

import (
	"errors"

	"110yards.ca/libs/go/core/logger"
)

var publishers = make(map[string]Publisher)

func createPublisher(isDev bool, topicName string) (Publisher, error) {
	if isDev {
		return NewVirtualPublisher(topicName), nil
	} else {
		return nil, errors.New("not implemented")
	}
}

func InitializePublishers(isDev bool, topicNames []string) error {

	for _, topicName := range topicNames {
		publisher, err := createPublisher(isDev, topicName)

		if err != nil {
			return err
		}

		publishers[topicName] = publisher
	}

	return nil
}

func GetPublisher(topicName string) Publisher {
	publisher, exists := publishers[topicName]

	if !exists {
		logger.Warn("publisher for topic %s not found", topicName)
		return nil
	}

	return publisher
}
