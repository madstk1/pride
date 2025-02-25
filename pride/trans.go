package pride

import (
	"fmt"
)

// PrintTransPrideFlag prints a trans pride flag in your terminal.
func PrintTransPrideFlag() {
	_, height := windowSize()
	segmentHeight := height / 5
	for segment := 1; segment <= 5; segment++ {
		line := "\n"
		for i := 0; i < segmentHeight-1; i++ {
			line += "\n"
		}
		if segment == 5 {
			line += "\n"
		}

		var colorCode string
		if segment == 1 || segment == 5 {
			colorCode = "106"
		}
		if segment == 2 || segment == 4 {
			colorCode = "105"
		}
		if segment == 3 {
			colorCode = "107"
		}
		fmt.Printf("\x1B[%sm%s\x1B[0m", colorCode, line)
	}
}
