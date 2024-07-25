package event

type Subscriber interface {
	SubscribedEvents() any
}

type SubscriberFunc func() any

func (f SubscriberFunc) SubscribedEvents() any {
	return f()
}
