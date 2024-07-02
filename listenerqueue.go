package event

import (
	"cmp"
	"slices"
)

type ListenerQueueItem struct {
	Listener Listener
	Priority int
}

type ListenerQueue []*ListenerQueueItem

func (lq *ListenerQueue) enqueue(listeners ...*ListenerQueueItem) {
	*lq = append(*lq, listeners...)

	// Most frequently we listen to some events on start up
	// and the rest of the work is done inside event listener handlers when events get occured.
	// So sorting is performed only when listeners are being set up.
	slices.SortStableFunc(*lq, func(a, b *ListenerQueueItem) int {
		return cmp.Compare(a.Priority, b.Priority)
	})
}

func (lq ListenerQueue) handleEvent(e Event) error {
	for _, item := range lq {
		err := item.Listener.HandleEvent(e)
		if err != nil {
			return err
		}
	}

	return nil
}
