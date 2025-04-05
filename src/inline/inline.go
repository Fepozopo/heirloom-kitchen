package inline

import (
	"errors"
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
