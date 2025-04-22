package main

type state struct {
	CFG      *Config
	Settings *Settings
	CLI      *cli
}

func InitializeState() *state {
	return &state{
		CFG: &Config{
			IsDebugMode: true,
		},
		Settings: NewSettings(), // TODO Read settings from file
		CLI:      &cli{},
	}
}
