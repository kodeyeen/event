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
