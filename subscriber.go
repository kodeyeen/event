package event

type Subscriber interface {
	SubscribedEvents() map[Type]ListenerQueue
}

type SubscriberFunc func() map[Type]ListenerQueue

func (f SubscriberFunc) SubscribedEvents() map[Type]ListenerQueue {
	return f()
}
