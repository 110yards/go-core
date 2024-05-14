package eventbus

type SubscriptionConfig struct {
	ExpirationDays     int
	RetentionDays      int
	AckDeadlineSeconds int
}

func DefaultSubscriptionConfig() SubscriptionConfig {
	return SubscriptionConfig{
		ExpirationDays:     31,
		RetentionDays:      7,
		AckDeadlineSeconds: 60,
	}
}

type SubscriptionCreator interface {
	CreateSubscription(subscriptionName, topicName, endpoint string, config SubscriptionConfig) error
}
