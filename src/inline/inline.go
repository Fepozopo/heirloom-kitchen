package inline

import (
	"errors"
	"regexp"
	"strings"

	"github.com/Fepozopo/pantry-of-the-past/src/nodes"
)

// SplitNodesDelimiter splits a list of nodes based on a delimiter
func SplitNodesDelimiter(oldNodes []nodes.TextNode, delimiter string, textType nodes.TextType) ([]nodes.TextNode, error) {
	newNodes := []nodes.TextNode{}

	for _, node := range oldNodes {
		// Only process nodes that are of the "text" type
		if node.Type != nodes.Normal {
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
						Type: nodes.Normal,
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

// SplitNodesImage splits a list of nodes based on image markdown syntax
// It takes a list of TextNodes and returns a new list of TextNodes where
// image markdown syntax is split into ImageText nodes.
func SplitNodesImage(oldNodes []nodes.TextNode) []nodes.TextNode {
	newNodes := []nodes.TextNode{}
	// Regex to match image markdown syntax: ![alt text](image URL)
	imagePattern := regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)

	for _, node := range oldNodes {
		// Skip nodes that are not NormalText
		if node.Type != nodes.Normal {
			newNodes = append(newNodes, node)
			continue
		}

		parts := imagePattern.Split(node.Text, -1)
		matches := imagePattern.FindAllStringSubmatch(node.Text, -1)

		for i, part := range parts {
			// Add normal text parts
			if part != "" {
				newNodes = append(newNodes, nodes.TextNode{
					Type: nodes.Normal,
					Text: part,
					URL:  "",
				})
			}
			// Add image nodes for matches
			if i < len(matches) {
				newNodes = append(newNodes, nodes.TextNode{
					Type: nodes.Image,
					Text: matches[i][1],
					URL:  matches[i][2],
				})
			}
		}
	}

	return newNodes
}

// SplitNodesLink splits a list of nodes based on link markdown syntax
// It takes a list of TextNodes and returns a new list of TextNodes where
// link markdown syntax is split into LinkText nodes.
func SplitNodesLink(oldNodes []nodes.TextNode) []nodes.TextNode {
	newNodes := []nodes.TextNode{}
	// Regex to match link markdown syntax: [anchor text](link URL)
	linkPattern := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

	for _, node := range oldNodes {
		// Skip nodes that are not NormalText
		if node.Type != nodes.Normal {
			newNodes = append(newNodes, node)
			continue
		}

		parts := linkPattern.Split(node.Text, -1)
		matches := linkPattern.FindAllStringSubmatch(node.Text, -1)

		for i, part := range parts {
			// Add normal text parts
			if part != "" {
				newNodes = append(newNodes, nodes.TextNode{
					Type: nodes.Normal,
					Text: part,
					URL:  "",
				})
			}
			// Add link nodes for matches
			if i < len(matches) {
				newNodes = append(newNodes, nodes.TextNode{
					Type: nodes.Link,
					Text: matches[i][1],
					URL:  matches[i][2],
				})
			}
		}
	}

	return newNodes
}

// TextToTextNodes converts a raw string of markdown text into a list of TextNodes.
func TextToTextNodes(text string) []nodes.TextNode {
	// Create an initial TextNode for the entire text
	nodesList := []nodes.TextNode{
		{Type: nodes.Normal, Text: text, URL: ""},
	}

	// Process the Markdown syntax for each type: images, links, code blocks, bold, and italic text in that order
	nodesList = SplitNodesImage(nodesList)                           // Process images
	nodesList = SplitNodesLink(nodesList)                            // Process links
	nodesList, _ = SplitNodesDelimiter(nodesList, "`", nodes.Code)   // Process code blocks
	nodesList, _ = SplitNodesDelimiter(nodesList, "**", nodes.Bold)  // Process bold text
	nodesList, _ = SplitNodesDelimiter(nodesList, "*", nodes.Italic) // Process italic text

	return nodesList
}
