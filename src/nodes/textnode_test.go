package nodes

import (
	"testing"
)

func TestTextNodeEquals(t *testing.T) {
	node1 := TextNode{
		Type: Link,
		Text: "This is some anchor text",
		URL:  "https://www.boot.dev",
	}

	node2 := TextNode{
		Type: Bold,
		Text: "This is some bold text",
		URL:  "",
	}

	node3 := TextNode{
		Type: Link,
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

func TestTextNodeString(t *testing.T) {
	node := TextNode{
		Type: Link,
		Text: "Anchor text",
		URL:  "https://www.boot.dev",
	}

	expected := `TextNode("Anchor text", "link", "https://www.boot.dev")`
	if node.String() != expected {
		t.Errorf("Expected %s, got %s", expected, node.String())
	}
}
