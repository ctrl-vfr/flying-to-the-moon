package main

import (
	"fmt"
	"time"
)

func main() {
	// Reference new moon
	referenceNewMoon := time.Date(2022, time.December, 23, 11, 16, 0, 0, time.UTC)

	// Current time
	now := time.Now().UTC()

	// Difference in days
	days := now.Sub(referenceNewMoon).Hours() / 24

	// Length of the lunar cycle
	lunarCycle := 29.53058867 // Average length in days

	// Current phase of the moon
	currentPhase := days / lunarCycle
	phase := currentPhase - float64(int64(currentPhase)) // Fractional part

	// Determine the moon phase emoji
	var moonPhaseEmoji string
	switch {
	case phase <= 0.03 || phase > 0.97:
		moonPhaseEmoji = "🌑"
	case phase <= 0.22:
		moonPhaseEmoji = "🌒"
	case phase <= 0.35:
		moonPhaseEmoji = "🌓"
	case phase <= 0.47:
		moonPhaseEmoji = "🌔"
	case phase <= 0.63:
		moonPhaseEmoji = "🌕"
	case phase <= 0.72:
		moonPhaseEmoji = "🌖"
	case phase <= 0.78:
		moonPhaseEmoji = "🌗"
	default:
		moonPhaseEmoji = "🌘"
	}

	fmt.Printf("%s", moonPhaseEmoji)
}
