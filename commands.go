package main

type command struct {
	name      string
	arguments []string
}

type commands struct {
	values map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.values[name] = f
}
