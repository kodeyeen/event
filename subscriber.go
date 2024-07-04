package event

type Subscriber interface {
	SubscribedEvents() map[Type][]Listener
}

type SubscriberFunc func() map[Type][]Listener

func (f SubscriberFunc) SubscribedEvents() map[Type][]Listener {
	return f()
}
