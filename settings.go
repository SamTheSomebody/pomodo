package main

import "pomodo/ui"

type Settings struct {
	Timer timerSettings
}

type timerSettings struct {
	// TODO time display settings
	IsHelpVisible bool // true
	ProgressBar   ui.ProgressBarSettings
}

func NewSettings() *Settings {
	return &Settings{
		Timer: timerSettings{
			IsHelpVisible: true,
			ProgressBar:   ui.NewProgressBarSettings(),
		},
	}
}
