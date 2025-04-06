package nodes

import "fmt"

type TextType int

const (
	Normal TextType = iota
	Bold
	Italic
	Code
	Link
	Image
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
	case Normal:
		typeStr = "text"
	case Bold:
		typeStr = "bold"
	case Italic:
		typeStr = "italic"
	case Code:
		typeStr = "code"
	case Link:
		typeStr = "link"
	case Image:
		typeStr = "image"
	default:
		typeStr = "unknown"
	}
	return fmt.Sprintf("TextNode(%q, %q, %q)", tn.Text, typeStr, tn.URL)
}
