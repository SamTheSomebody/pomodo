package main

import (
	"errors"
	"fmt"
)

type command struct {
	name      string
	arguments map[string]string
}

type commands struct {
	values map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.values[name] = f
}

func (c *commands) run(cmd command, s *state) error {
	f, ok := c.values[cmd.name]
	if !ok {
		return errors.New(fmt.Sprintf("Command: [%v] does not exist", cmd.name))
	}
	return f(s, cmd)
}
