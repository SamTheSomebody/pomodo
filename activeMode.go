package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func activeMode(s *state) {
	if s.CFG.IsDebugMode {
		fmt.Println("[DEBUG] Entered active mode")
	}

	t := timer{
		Duration: time.Second * 3,
		State:    s,
	}
	t.start()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if s.CFG.IsDebugMode {
			fmt.Printf("[DEBUG] Received input: %v\n", scanner.Text())
		}
	}
}
