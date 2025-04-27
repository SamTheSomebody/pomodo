package main

func runSingleCommand(s *state) {
		cmds := commands{
			values: make(map[string]func(*state, command) error)
		}	
		cmds.register("add", addTaskHandler) //taskname -optional params:
		cmds.register("done", completeTaskHandler) //taskname -optional params: time taken
		cmds.register("edit", editTaskHandler) //taskname -optional params
		cmds.register("start", startTaskHandler) //taskname or duration
		cmds.register("allocate", allocateTimeHandler) //duration
		c := createCommandFromInput()

}

func createCommandFromInput() command {
  if len(os.Args) < 2 {
    log.Fatal("[Fatal] Less than 2 arguments provided")
  }
  name := os.Args[1]
  var args []string
  if len(os.Args) > 2 {
    args = os.Args[2:]
  }
  c := command{
    name: name,
    arguments: args,
  }
  return c
}
