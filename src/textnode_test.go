package src

import (
	"testing"
)

func TestTextNodeEquals(t *testing.T) {
	node1 := TextNode{
		Type: LinkText,
		Text: "This is some anchor text",
		URL:  "https://www.boot.dev",
	}

	node2 := TextNode{
		Type: BoldText,
		Text: "This is some bold text",
		URL:  "",
	}

	node3 := TextNode{
		Type: LinkText,
		Text: "This is some anchor text",
		URL:  "https://www.boot.dev",
	}

	if !node1.Equals(node3) {
		t.Errorf("Expected nodes to be equal")
	}

	if node1.Equals(node2) {
		t.Errorf("Expected nodes to not be equal")
	}
}

func TestTextNodeRepr(t *testing.T) {
	node := TextNode{
		Type: LinkText,
		Text: "Anchor text",
		URL:  "https://www.boot.dev",
	}

	expected := `TextNode("Anchor text", "LinkText", "https://www.boot.dev")`
	if node.Repr() != expected {
		t.Errorf("Expected %s, got %s", expected, node.Repr())
	}
}
