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

type UserRegisteredEvent struct {
	Username string
}

func (e *UserRegisteredEvent) Type() event.Type {
	return "user.registered"
}

func main() {
	// Create an event dispatcher
	events := event.NewDispatcher()

	// Register event listener
	events.Listen("user.registered", event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("New user: %s\n", e.(*UserRegisteredEvent).Username)
		return nil
	}), 0)

	// Dispatch event
	events.Dispatch(&UserRegisteredEvent{
		Username: "kodeyeen",
	})
}
```
