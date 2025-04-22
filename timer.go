package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"pomodo/ui"
)

type timer struct {
	Duration      time.Duration
	Remaining     time.Duration
	Ticker        *time.Ticker
	State         *state
	Done          chan (bool)
	Command       chan (byte)
	IsRunning     bool
	IsInitialized bool
}

const helpString = "('p': Pause, 'r': Restart, s': Stop, 'd': Done)"

func (t *timer) initialze() {
	t.IsInitialized = true
	t.Remaining = t.Duration
	t.Done = make(chan bool, 1)
	t.Command = make(chan byte, 1)
}

func (t *timer) start() {
	if !t.IsInitialized {
		t.initialze()
	}

	t.IsRunning = true
	t.Ticker = time.NewTicker(time.Second)

	err := t.State.CLI.MakeRaw()
	if err != nil {
		log.Fatal(err)
	}

	t.basicTickFunction()
}

func (t *timer) pause() error {
	if t.Ticker == nil {
		return fmt.Errorf("timer hasn't been started")
	}
	t.IsRunning = false
	t.generatePausedString()
	err := t.State.CLI.Restore()
	if err != nil {
		return err
	}
	t.Ticker.Stop()
	return nil
}

func (t *timer) stop() error {
	err := t.pause()
	if err != nil {
		return err
	}

	t.generateFinishedString()

	err = t.State.CLI.Notify()
	if err != nil {
		return err
	}

	return nil
}

func (t *timer) basicTickFunction() {
	t.generateTimerString()
	go t.waitForInput()
	for {
		select {
		case cmd := <-t.Command:
			switch cmd {
			case byte('r'):
				t.pause()
				t.initialze()
				t.start()
			case byte('p'):
				t.pause()
			case byte('s'):
				t.stop()
			case byte('d'):
				t.stop()
			}
		case <-t.Done:
			err := t.stop()
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

func (t *timer) waitForInput() {
	// TODO Add cancellation token
	reader := bufio.NewReader(os.Stdin)
	for {
		select {}
		b, err := reader.ReadByte()
		if err != nil {
			log.Fatal(err)
		}
		t.Command <- b
	}
}

func (t *timer) generateTimerString() {
	output := "Remaining: "
	output += t.Remaining.String() + " "
	output += ui.GenerateProgressBar(float32(t.Duration-t.Remaining), float32(t.Duration), t.State.Settings.Timer.ProgressBar)
	if t.State.Settings.Timer.IsHelpVisible {
		output += " " + helpString
	}
	output += ": "
	t.State.CLI.RewriteLine(output)
}

func (t *timer) generatePausedString() {
	output := fmt.Sprintf("[Paused] (%v) %v", t.Duration, ui.GenerateProgressBar(float32(t.Duration-t.Remaining), float32(t.Duration), t.State.Settings.Timer.ProgressBar))
	t.State.CLI.RewriteLine(output)
}

func (t *timer) generateFinishedString() {
	output := fmt.Sprintf("Done! (%v) %v\n", t.Duration, ui.GenerateProgressBar(1.0, 1.0, t.State.Settings.Timer.ProgressBar))
	t.State.CLI.RewriteLine(output)
}
