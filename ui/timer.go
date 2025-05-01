package ui

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Timer struct {
	Duration      time.Duration
	Remaining     time.Duration
	Ticker        *time.Ticker
	Done          chan (bool)
	Command       chan (byte)
	IsRunning     bool
	IsInitialized bool
}

const helpString = "('p': Pause, 'r': Restart, s': Stop, 'd': Done)"

func (t *Timer) initialze() {
	t.IsInitialized = true
	t.Remaining = t.Duration
	t.Done = make(chan bool, 1)
	t.Command = make(chan byte, 1)
}

func (t *Timer) Start() {
	if !t.IsInitialized {
		t.initialze()
	}

	t.IsRunning = true
	t.Ticker = time.NewTicker(time.Second)

	/*err := t.State.CLI.MakeRaw()
	if err != nil {
		log.Fatal(err)
	}*/

	t.basicTickFunction()
}

func (t *Timer) Pause() error {
	if t.Ticker == nil {
		return fmt.Errorf("Timer hasn't been started")
	}
	t.IsRunning = false
	t.generatePausedString()
	/*err := t.State.CLI.Restore()
	if err != nil {
		return err
	}*/
	t.Ticker.Stop()
	return nil
}

func (t *Timer) Stop() error {
	err := t.Pause()
	if err != nil {
		return err
	}

	t.generateFinishedString()

	/*err = t.State.CLI.Notify()
	if err != nil {
		return err
	}*/

	return nil
}

func (t *Timer) basicTickFunction() {
	t.generateTimerString()
	go t.waitForInput()
	for {
		select {
		case cmd := <-t.Command:
			switch cmd {
			case byte('r'):
				t.Pause()
				t.initialze()
				t.Start()
			case byte('p'):
				t.Pause()
			case byte('s'):
				t.Stop()
			case byte('d'):
				t.Stop()
			}
		case <-t.Done:
			err := t.Stop()
			if err != nil {
				fmt.Println(err)
			}
			return
		case <-t.Ticker.C:
			t.Remaining -= time.Second
			t.generateTimerString()
			if t.Remaining <= time.Millisecond {
				t.Done <- true
			}
		}
	}
}

func (t *Timer) waitForInput() {
	// TODO Add cancellation token
	reader := bufio.NewReader(os.Stdin)
	for {
		// select {}
		b, err := reader.ReadByte()
		if err != nil {
			log.Fatal(err)
		}
		t.Command <- b
	}
}

func (t *Timer) generateTimerString() {
	output := "Remaining: "
	output += t.Remaining.String() + " "
	output += GenerateProgressBar(float32(t.Duration-t.Remaining), float32(t.Duration))
	if viper.GetBool("timer.isHelpVisible") {
		output += " " + helpString
	}
	// output += ": "
	fmt.Println(output)
	// t.State.CLI.RewriteLine(output)
}

func (t *Timer) generatePausedString() {
	output := fmt.Sprintf("[Paused] (%v) %v", t.Duration, GenerateProgressBar(float32(t.Duration-t.Remaining), float32(t.Duration)))
	fmt.Println(output)
	// t.State.CLI.RewriteLine(output)
}

func (t *Timer) generateFinishedString() {
	output := fmt.Sprintf("Done! (%v) %v\n", t.Duration, GenerateProgressBar(1.0, 1.0))
	fmt.Println(output)
	// t.State.CLI.RewriteLine(output)
}
