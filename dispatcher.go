package event

import (
	"reflect"
	"slices"
)

type Dispatcher struct {
	listeners map[Type][]listener
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		listeners: make(map[Type][]listener),
	}
}

func Dispatch[T any](d *Dispatcher, _type Type, evt T) bool {
	listeners := d.listeners[_type]

	for _, lis := range listeners {
		handler, ok := lis.handler.(func(T) bool)
		if !ok {
			continue
		}

		callNext := handler(evt)

		if lis.once {
			d.Off(_type, lis.handler)
		}

		if !callNext {
			return false
		}
	}

	return true
}

func (d *Dispatcher) Listen(_type Type, handler any) {
	listeners := d.listeners[_type]

	listeners = append(listeners, listener{
		handler: handler,
		once:    false,
	})

	d.listeners[_type] = listeners
}

func (d *Dispatcher) On(_type Type, handler any) {
	d.Listen(_type, handler)
}

func (d *Dispatcher) ListenOnce(_type Type, handler any) {
	listeners := d.listeners[_type]

	listeners = append(listeners, listener{
		handler: handler,
		once:    true,
	})

	d.listeners[_type] = listeners
}

func (d *Dispatcher) Once(_type Type, handler any) {
	d.ListenOnce(_type, handler)
}

func (d *Dispatcher) Unlisten(_type Type, handler any) {
	listeners := d.listeners[_type]

	idx := slices.IndexFunc(listeners, func(l listener) bool {
		return reflect.ValueOf(l.handler).Pointer() == reflect.ValueOf(handler).Pointer()
	})

	if idx == -1 {
		return
	}

	d.listeners[_type] = slices.Delete(listeners, idx, idx+1)
}

func (d *Dispatcher) Off(_type Type, handler any) {
	d.Unlisten(_type, handler)
}

func (d *Dispatcher) Has(_type Type) bool {
	_, ok := d.listeners[_type]
	return ok
}
