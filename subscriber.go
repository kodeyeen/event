package event

type Subscriber interface {
	SubscribedEvents() map[Type]map[int][]Listener
}

type SubscriberFunc func() map[Type]map[int][]Listener

func (f SubscriberFunc) SubscribedEvents() map[Type]map[int][]Listener {
	return f()
}
