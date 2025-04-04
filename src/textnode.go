package src

import "fmt"

type TextType int

const (
	NormalText TextType = iota
	BoldText
	ItalicText
	CodeText
	LinkText
	ImageText
)

type TextNode struct {
	Type TextType
	Text string
	URL  string
}

// Equals method to compare two TextNode objects
func (tn TextNode) Equals(other TextNode) bool {
	return tn.Type == other.Type && tn.Text == other.Text && tn.URL == other.URL
}

// String method to return a string representation of the TextNode object
func (tn TextNode) String() string {
	// Convert TextType to its string representation
	typeStr := ""
	switch tn.Type {
	case NormalText:
		typeStr = "text"
	case BoldText:
		typeStr = "bold"
	case ItalicText:
		typeStr = "italic"
	case CodeText:
		typeStr = "code"
	case LinkText:
		typeStr = "link"
	case ImageText:
		typeStr = "image"
	default:
		typeStr = "unknown"
	}
	return fmt.Sprintf("TextNode(%q, %q, %q)", tn.Text, typeStr, tn.URL)
}
