package inline

import (
	"errors"
	"regexp"
	"strings"

	"github.com/Fepozopo/heirloom-kitchen/src/nodes"
)

// SplitNodesDelimiter splits a list of nodes based on a delimiter
func SplitNodesDelimiter(oldNodes []nodes.TextNode, delimiter string, textType nodes.TextType) ([]nodes.TextNode, error) {
	newNodes := []nodes.TextNode{}

	for _, node := range oldNodes {
		// Only process nodes that are of the "text" type
		if node.Type != nodes.NormalText {
			newNodes = append(newNodes, node)
			continue
		}

		// Split the text by the delimiter
		parts := strings.Split(node.Text, delimiter)

		// If there's an unmatched delimiter, we raise an error
		if len(parts)%2 == 0 {
			return nil, errors.New("unmatched delimiter '" + delimiter + "' found in text: " + node.Text)
		}

		// Alternate between regular text and the delimited text
		for i, part := range parts {
			if i%2 == 0 {
				// Even indices are regular text (non-delimited)
				if part != "" {
					newNodes = append(newNodes, nodes.TextNode{
						Type: nodes.NormalText,
						Text: part,
						URL:  node.URL,
					})
				}
			} else {
				// Odd indices are delimited text (like code, bold, italic, etc.)
				newNodes = append(newNodes, nodes.TextNode{
					Type: textType,
					Text: part,
					URL:  node.URL,
				})
			}
		}
	}

	return newNodes, nil
}

// extractMarkdownImages extracts image URLs and alt text from markdown text.
func ExtractMarkdownImages(text string) [][2]string {
	// Create a slice to store the image URLs and alt text.
	imageURLs := [][2]string{}

	// Find all image URLs in the text using regex.
	re := regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)
	matches := re.FindAllStringSubmatch(text, -1)

	if len(matches) > 0 {
		for _, match := range matches {
			// Append the alt text and image URL to the slice.
			imageURLs = append(imageURLs, [2]string{match[1], match[2]})
		}
	}

	return imageURLs
}

// extractMarkdownLinks extracts link URLs and anchor text from markdown text.
func ExtractMarkdownLinks(text string) [][2]string {
	// Create a slice to store the link URLs and anchor text.
	linkURLs := [][2]string{}

	// Find all link URLs in the text using regex.
	re := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	matches := re.FindAllStringSubmatch(text, -1)

	if len(matches) > 0 {
		for _, match := range matches {
			// Append the anchor text and link URL to the slice.
			linkURLs = append(linkURLs, [2]string{match[1], match[2]})
		}
	}

	return linkURLs
}
