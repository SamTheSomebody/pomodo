package main

import (
	"bufio"
	"errors"
	"os"

	"github.com/gen2brain/beeep"
	"golang.org/x/term"
)

type cli struct {
	oldState *term.State
	isRaw    bool
}

func (c *cli) MakeRaw() error {
	if c.isRaw {
		return errors.New("terminal is already raw")
	}
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	c.oldState = oldState
	c.isRaw = true
	return nil
}

const (
	moveUpTerminalLine    = "\033[1A"
	moveFrontTerminalLine = "\033[G"
	clearTerminalLine     = "\033[K"
)

func (c *cli) RewriteLine(s string) {
	c.WriteLine(moveFrontTerminalLine + clearTerminalLine + s)
}

func (c *cli) WriteLine(s string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(s)
	writer.Flush()
}

func (c *cli) Notify() error {
	return beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
}

func (c *cli) Restore() error {
	if !c.isRaw {
		return errors.New("terminal isn't raw")
	}
	err := term.Restore(int(os.Stdin.Fd()), c.oldState)
	if err != nil {
		return err
	}
	c.isRaw = false
	return nil
}
