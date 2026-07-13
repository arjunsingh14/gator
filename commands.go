package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
	handlerMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	commandHandler, exists := c.handlerMap[cmd.name]
	if !exists {
		return errors.New("Invalid command")
	}
	return commandHandler(s, cmd)
}

func (c * commands) register(name string, f func(*state, command) error ) error {
	c.handlerMap[name] = f
	return nil
}

