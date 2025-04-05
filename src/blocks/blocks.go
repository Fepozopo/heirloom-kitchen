package blocks

import (
	"regexp"
	"strings"
)

// MarkdownToBlocks takes a raw markdown string and returns a list of "block" strings.
func MarkdownToBlocks(markdown string) []string {
	// Strip any leading or trailing whitespace from the entire document
	markdown = strings.TrimSpace(markdown)

	// Normalize the markdown by removing leading spaces from each line
	lines := strings.Split(markdown, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimLeft(line, " ")
	}
	markdown = strings.Join(lines, "\n")

	// Remove any "empty" blocks due to excessive newlines
	re := regexp.MustCompile(`\n{3,}`)
	markdown = re.ReplaceAllString(markdown, "\n\n")

	// Split the markdown into blocks based on double newlines
	blocks := strings.Split(markdown, "\n\n")

	// Strip leading/trailing whitespace from each block individually
	for i, block := range blocks {
		blocks[i] = strings.TrimSpace(block)
	}

	return blocks
}
