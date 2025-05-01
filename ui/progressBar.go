package ui

import (
	"strings"

	"github.com/spf13/viper"
)

func GenerateProgressBar(current float32, target float32) string {
	if !viper.GetBool("progressbar.isVisible") {
		return ""
	}

	size := viper.GetInt("progressBar.size")
	fillSegment := viper.GetString("progressBar.fillSegment")
	emptySegment := viper.GetString("progressBar.emptySegment")
	progress := int(current / target * float32(size))

	if viper.GetBool("progressbar.isDecreasing") {
		return strings.Repeat(emptySegment, size-progress) + strings.Repeat(fillSegment, progress)
	}
	return strings.Repeat(fillSegment, progress) + strings.Repeat(emptySegment, size-progress)
}
