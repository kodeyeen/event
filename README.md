# event

A Go event management package.

## Installation

```shell
go get github.com/kodeyeen/event
```

## Quickstart

```go
package main

import (
	"context"
	"fmt"

	"github.com/kodeyeen/event"
)

// Declare event payload.
// It's ok to name it with Event suffix though it's not the event itself, but its payload.
type UserRegisteredEvent struct {
	Username string
}

func main() {
	// Create an event dispatcher.
	dispr := event.NewDispatcher()

	// Register an event listener.
	dispr.Listen("user.registered", event.ListenerFunc(func(e event.Event) error {
		ep := e.Payload().(*UserRegisteredEvent)

		fmt.Printf("New user: %s\n", ep.Username)

		return nil
	}))

	// Create an event to be dispatched
	userRegisteredEvt := event.New("user.registered", &UserRegisteredEvent{
		Username: "kodeyeen",
	})

	// Dispatch the event.
	dispr.HandleEvent(context.Background(), userRegisteredEvt)
}
```
