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
	*event.Base
	Username string
}

func main() {
	// Create an event dispatcher
	events := event.NewDispatcher()

	// Register an event listener
	events.Listen("user.registered", event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("New user: %s\n", e.(*UserRegisteredEvent).Username)
		return nil
	}), 0)

	// Dispatch an event
	events.Dispatch(&UserRegisteredEvent{
		Base: event.NewBase("user.registered")
		Username: "kodeyeen",
	})
}
```
