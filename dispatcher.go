package event

type Dispatcher struct {
	listenerQueues map[Type]ListenerQueue
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		listenerQueues: make(map[Type]ListenerQueue),
	}
}

func (d *Dispatcher) Dispatch(e Event) error {
	lq := d.listenerQueues[e.Type()]

	lq.handleEvent(e)

	return nil
}

func (d *Dispatcher) Listen(_type Type, listener Listener, priority int) {
	lq := d.listenerQueues[_type]

	lq.enqueue(&ListenerQueueItem{
		Listener: listener,
		Priority: priority,
	})

	d.listenerQueues[_type] = lq
}

func (d *Dispatcher) ListenFunc(_type Type, listener func(Event) error, priority int) {
	d.Listen(_type, ListenerFunc(listener), priority)
}

func (d *Dispatcher) HasListener(_type Type) bool {
	_, ok := d.listenerQueues[_type]
	return ok
}

func (d *Dispatcher) Subscribe(subscriber Subscriber) {
	events := subscriber.SubscribedEvents()

	for _type, listeners := range events {
		lq := d.listenerQueues[_type]
		lq.enqueue(listeners...)
		d.listenerQueues[_type] = lq
	}
}

func (d *Dispatcher) SubscribeFunc(subscriber func() map[Type]ListenerQueue) {
	d.Subscribe(SubscriberFunc(subscriber))
}
