package ui

import "strings"

type ProgressBarSettings struct {
	IsVisible    bool   // true
	IsDecreasing bool   // false
	Size         int    // 10
	FillSegment  string // "◼"
	EmptySegment string // "◻"
}

func NewProgressBarSettings() ProgressBarSettings {
	return ProgressBarSettings{
		IsVisible:    true,
		Size:         10,
		FillSegment:  "◼",
		EmptySegment: "◻",
	}
}

func GenerateProgressBar(current float32, target float32, settings ProgressBarSettings) string {
	if !settings.IsVisible {
		return ""
	}
	progress := int(current / target * float32(settings.Size))
	if settings.IsDecreasing {
		return strings.Repeat(settings.FillSegment, settings.Size-progress) + strings.Repeat(settings.EmptySegment, progress)
	}
	return strings.Repeat(settings.FillSegment, progress) + strings.Repeat(settings.EmptySegment, settings.Size-progress)
}
