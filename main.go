package main

import (
	"log"
	"math"
	"strings"
)

func main() {
	//words := []string{"What", "must", "be", "shall", "be."}
	words := []string{"Listen", "to", "many,", "speak", "to", "a", "few."}
	//words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth := 12

	result := fullJustify(words, maxWidth)
	for _, line := range result {
		log.Println(line, "|", len(line))
	}
}

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	line := ""
	space := " "

	for i, word := range words {
		if len(word) == maxWidth {
			if line != "" {
				result = append(result, justify(line, maxWidth, false))
				line = ""
			}
			result = append(result, word)
			continue
		}

		if maxWidth >= len(line)+len(word)+1 {
			if len(line) != 0 {
				line += space
			}
			line = line + word
		} else {
			result = append(result, justify(line, maxWidth, false))
			line = word
		}

		if i == len(words)-1 {
			result = append(result, justify(line, maxWidth, true))
		}

	}

	return result
}

func justify(line string, maxWidth int, lastLine bool) string {
	result := ""

	split := strings.Split(line, " ")
	availableSpace := len(split) - 1
	missing := maxWidth - len(line) + availableSpace
	step := int(math.Ceil(float64(missing) / float64(availableSpace)))

	if availableSpace == 0 || lastLine {
		result = line + strings.Repeat(" ", missing-availableSpace)
		return result
	}

	for i := 0; i < len(split); i++ {
		y := ""
		if i != 0 {
			y = strings.Repeat(" ", step)
			missing = missing - step
			availableSpace = availableSpace - 1
			step = int(math.Ceil(float64(missing) / float64(availableSpace)))
		}
		result = result + y + split[i]
	}

	return result
}
