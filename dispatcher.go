package event

type Dispatcher struct {
	listeners     map[Type]map[int][]Listener
	minPriorities map[Type]int
	maxPriorities map[Type]int
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		listeners:     make(map[Type]map[int][]Listener),
		minPriorities: make(map[Type]int),
		maxPriorities: make(map[Type]int),
	}
}

func (d *Dispatcher) Dispatch(e Event) error {
	minPriority := d.minPriorities[e.Type()]
	maxPriority := d.maxPriorities[e.Type()]
	listenerPriorities := d.listeners[e.Type()]

	for i := maxPriority; i >= minPriority; i-- {
		listeners := listenerPriorities[i]

		for _, listener := range listeners {
			err := listener.HandleEvent(e)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (d *Dispatcher) Listen(_type Type, listener Listener, priority int) {
	if priority < d.minPriorities[_type] {
		d.minPriorities[_type] = priority
	} else if priority > d.maxPriorities[_type] {
		d.maxPriorities[_type] = priority
	}

	listeners, ok := d.listeners[_type]
	if !ok {
		listeners = make(map[int][]Listener)
	}

	listeners[priority] = append(listeners[priority], listener)

	d.listeners[_type] = listeners
}

func (d *Dispatcher) ListenFunc(_type Type, listener func(Event) error, priority int) {
	d.Listen(_type, ListenerFunc(listener), priority)
}

func (d *Dispatcher) HasListener(_type Type) bool {
	_, ok := d.listeners[_type]
	return ok
}

func (d *Dispatcher) Subscribe(subscriber Subscriber) {
	events := subscriber.SubscribedEvents()

	for _type, listenerPriorities := range events {
		for priority, listeners := range listenerPriorities {
			for _, listener := range listeners {
				d.Listen(_type, listener, priority)
			}
		}
	}
}

func (d *Dispatcher) SubscribeFunc(subscriber func() map[Type]map[int][]Listener) {
	d.Subscribe(SubscriberFunc(subscriber))
}
