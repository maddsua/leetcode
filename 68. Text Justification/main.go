package main

import (
	"fmt"
	"math"
	"strings"
)

type LineCtx struct {
	words        []string
	length       int
	spacedLength int
}

func breakIntoLines(words []string, maxWidth int) []LineCtx {

	lines := make([]LineCtx, 0)

	for i := 0; i < len(words); {

		line := LineCtx{}

		for ; i < len(words); i++ {

			wordLength := len(words[i])
			addLength := wordLength
			if len(line.words) > 0 {
				addLength++
			}

			if line.spacedLength+addLength <= maxWidth {
				line.words = append(line.words, words[i])
				line.spacedLength += addLength
				line.length += wordLength
				continue
			}

			break
		}

		lines = append(lines, line)
	}

	return lines
}

func calculateSpacings(line LineCtx, maxWidth int) []int {

	totalSpacing := maxWidth - line.length

	if len(line.words) < 2 {
		return []int{totalSpacing}
	}

	spacings := len(line.words) - 1
	spaceWidth := int(math.Ceil(float64(totalSpacing) / float64(spacings)))

	spaces := make([]int, spacings)
	for i := 0; i < len(spaces); i++ {
		spaces[i] = spaceWidth
	}

	var getMappedSpacingsTotal = func() int {

		sum := 0

		for _, v := range spaces {
			sum += v
		}

		return sum
	}

	for i := len(spaces) - 1; getMappedSpacingsTotal() > totalSpacing && i >= 0; i-- {
		spaces[i]--
	}

	return spaces
}

func fillSpace(width int) string {

	result := make([]byte, width)

	for i := 0; i < width; i++ {
		result[i] = ' '
	}

	return string(result)
}

func applySpacing(line LineCtx, maxWidth int) string {

	spaces := calculateSpacings(line, maxWidth)

	var mergedLine string

	for i, v := range line.words {

		mergedLine += v

		if i < len(spaces) {
			mergedLine += fillSpace(spaces[i])
		}
	}

	return mergedLine
}

func fullJustify(words []string, maxWidth int) []string {

	lines := breakIntoLines(words, maxWidth)

	result := make([]string, 0)

	for idx, line := range lines {

		if idx < len(lines)-1 {
			lineSpaced := applySpacing(line, maxWidth)
			result = append(result, lineSpaced)
		} else {

			lastLine := strings.Join(line.words, " ")

			padSize := maxWidth - len(lastLine)
			if padSize > 0 {
				lastLine += fillSpace(padSize)
			}

			result = append(result, lastLine)
		}

	}

	return result
}

func justifyAndPrint(input []string, width int) {

	lines := fullJustify(input, width)

	for _, v := range lines {
		fmt.Printf("'%s'\n", v)
	}

	print("\n")
}

func main() {

	justifyAndPrint([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)

	justifyAndPrint([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16)

	justifyAndPrint([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20)

	justifyAndPrint([]string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}, 16)
}
