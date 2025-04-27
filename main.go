package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Config struct {
	IsDebugMode bool
}

func main() {
	s := InitializeState()
	t := timer{
		Duration: time.Second * 3,
		State:    s,
	}
	t.start()

	if len(os.Args) == 0 {
		activeMode(s)
	} else {
		runSingleCommand(s)
	}
}

func activeMode(s *state) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if s.CFG.IsDebugMode {
			fmt.Println(scanner.Text())
		}
	}
}
