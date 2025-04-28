package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type Config struct {
	IsDebugMode bool
}

func main() {
	checkDatabase()
	s := InitializeState()
	if len(os.Args) == 0 {
		activeMode(s)
	} else {
		runSingleCommand(s)
	}
}

func activeMode(s *state) {
	t := timer{
		Duration: time.Second * 3,
		State:    s,
	}
	t.start()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if s.CFG.IsDebugMode {
			fmt.Println(scanner.Text())
		}
	}
}

func checkDatabase() {
	if _, err := os.Stat("./tasks.db"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("Database not found! Creating new one")
		f, err := os.Create("./tasks.db")
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}
}
