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
	events := event.NewDispatcher()

	// Register an event listener.
	events.Listen("user.registered", event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("New user: %s\n", e.Payload().(*UserRegisteredEvent).Username)
		return nil
	}), 0)

	// Create an event to be dispatched
	userRegisteredEvt := event.New("user.registered", &UserRegisteredEvent{
		Username: "kodeyeen",
	})

	// Dispatch the event.
	events.Dispatch(userRegisteredEvt)
}
```
