package google

import (
	"testing"

	"110yards.ca/libs/go/core/eventbus/gpubsub"
	"github.com/stretchr/testify/assert"
)

func TestTopicCreator(t *testing.T) {
	setup()

	err := gpubsub.InitializeGoogle(projectId)

	assert.Nil(t, err)

	topicCreator := gpubsub.NewTopicCreator()

	err = topicCreator.CreateTopic("test-topic")

	assert.Nil(t, err)
}
