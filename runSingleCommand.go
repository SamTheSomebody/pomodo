package main

import (
	"log"
	"os"
)

func runSingleCommand(s *state) {
	cmds := commands{
		values: make(map[string]func(*state, command) error),
	}
	cmds.register("list", listTaskHandler)
	cmds.register("add", addTaskHandler)           // taskname -optional params: -d due at, -t time estimate, -s description, -e enthusiasm, -p priority
	cmds.register("done", completeTaskHandler)     // taskname -optional params: time taken
	cmds.register("edit", editTaskHandler)         // taskname -optional params
	cmds.register("start", startTaskHandler)       // taskname or duration
	cmds.register("allocate", allocateTimeHandler) // duration
	c := createCommandFromInput()
	err := cmds.run(c, s)
	if err != nil {
		log.Fatal(err)
	}
}

func createCommandFromInput() command {
	if len(os.Args) < 2 {
		log.Fatal("[Fatal] Less than 2 arguments provided")
	}
	name := os.Args[1]
	args := processArguments()
	c := command{
		name:      name,
		arguments: args,
	}
	return c
}

func processArguments() map[string]string {
	args := make(map[string]string)
	var name, value string
	for _, a := range os.Args[2:] {
		if a[0] == '-' {
			addToArgumentsMap(name, value, args)
			value = ""
			name = a
		} else {
			value += a + " "
		}
	}
	addToArgumentsMap(name, value, args)
	return args
}

func addToArgumentsMap(name, value string, args map[string]string) {
	if len(name) > 0 {
		if len(value) == 0 {
			log.Fatal("No value specified for: " + name) // TODO in the future this should refer to the help section
		}
		args[name] = value
	} else {
		args["default"] = value
	}
}
