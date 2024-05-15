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

const (
	EventTypeUserRegistered event.Type = "user.registered"
)

type UserRegisteredEvent struct {
	Username string
}

func main() {
	// Create an event dispatcher
	events := event.NewDispatcher()

	// Register event listener
	events.Listen(EventTypeUserRegistered, func(e *UserRegisteredEvent) bool {
		fmt.Printf("New user: %s\n", e.Username)
		return true
	})

	// Dispatch event
	event.Dispatch(events, EventTypeUserRegistered, &UserRegisteredEvent{
		Username: "kodeyeen",
	})
}
```

## TODO

- Priorities
- Subscribe to multiple events
