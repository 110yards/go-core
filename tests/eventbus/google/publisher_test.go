package google

import (
	"testing"

	"110yards.ca/libs/go/core/eventbus/gpubsub"
	"github.com/stretchr/testify/assert"
)

func TestPublisher(t *testing.T) {
	setup()

	err := gpubsub.InitializeGoogle(projectId)

	assert.Nil(t, err)

	publisher := gpubsub.NewPublisher("test-topic")

	data := map[string]string{
		"key": "value",
	}
	err = publisher.Publish(data)

	assert.Nil(t, err)
}
