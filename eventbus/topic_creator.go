package eventbus

type TopicCreator interface {
	CreateTopic(topicName string) error
}
