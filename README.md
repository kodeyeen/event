## Installation

```shell
go get github.com/kodeyeen/event
```

## Quick start

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
	evtDispatcher := event.NewDispatcher()

	// Register event listener
	evtDispatcher.On(EventTypeUserRegistered, func(evt *UserRegisteredEvent) bool {
		fmt.Printf("New user: %s\n", evt.Username)
		return true
	})

	// Dispatch event
	event.Dispatch(evtDispatcher, EventTypeUserRegistered, &UserRegisteredEvent{
		Username: "kodeyeen",
	})
}
```
