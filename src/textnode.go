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

// Repr method to return a string representation of the TextNode object
func (tn TextNode) Repr() string {
	// Convert TextType to its string representation
	typeStr := ""
	switch tn.Type {
	case NormalText:
		typeStr = "NormalText"
	case BoldText:
		typeStr = "BoldText"
	case ItalicText:
		typeStr = "ItalicText"
	case CodeText:
		typeStr = "CodeText"
	case LinkText:
		typeStr = "LinkText"
	case ImageText:
		typeStr = "ImageText"
	}
	return fmt.Sprintf("TextNode(%q, %q, %q)", tn.Text, typeStr, tn.URL)
}
