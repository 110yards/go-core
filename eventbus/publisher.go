package eventbus

type Publisher interface {
	Publish(message interface{}) error
}
