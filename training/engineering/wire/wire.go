// The build tag makes sure the stub is not built in the final build.
//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// InitializeEvent creates an Event. It will error if the Event is staffed with
// a grumpy greeter.
func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
