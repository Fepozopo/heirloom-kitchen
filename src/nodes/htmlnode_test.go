package nodes

import (
	"testing"
)

func TestLeafNodeToHTML(t *testing.T) {
	l := &LeafNode{
		Tag:   "p",
		Value: "This is a paragraph of text",
		Props: nil,
	}

	html, err := l.ToHTML()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := "<p>This is a paragraph of text</p>"
	if html != expected {
		t.Errorf("Expected %s, got %s", expected, html)
	}
}

func TestParentNodeToHTML(t *testing.T) {
	p := &ParentNode{
		Tag: "div",
		Children: []HTMLNode{
			&LeafNode{
				Tag:   "p",
				Value: "Hello, world!",
				Props: nil,
			},
		},
		Props: nil,
	}

	html, err := p.ToHTML()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := "<div><p>Hello, world!</p></div>"
	if html != expected {
		t.Errorf("Expected %s, got %s", expected, html)
	}

	p2 := &ParentNode{
		Tag: "div",
		Children: []HTMLNode{
			&LeafNode{
				Tag:   "p",
				Value: "Hello, world!",
				Props: nil,
			},
			&LeafNode{
				Tag:   "p",
				Value: "Goodbye, world!",
				Props: nil,
			},
		},
		Props: nil,
	}

	html, err = p2.ToHTML()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected = "<div><p>Hello, world!</p><p>Goodbye, world!</p></div>"
	if html != expected {
		t.Errorf("Expected %s, got %s", expected, html)
	}

	p3 := &ParentNode{
		Tag: "div",
		Children: []HTMLNode{
			&LeafNode{
				Tag:   "b",
				Value: "Bold text",
				Props: nil,
			},
			&ParentNode{
				Tag: "p",
				Children: []HTMLNode{
					&LeafNode{
						Tag:   "i",
						Value: "Italic text",
						Props: nil,
					},
					&LeafNode{
						Tag:   "",
						Value: "Normal text",
						Props: nil,
					},
				},
				Props: nil,
			},
		},
		Props: nil,
	}

	html, err = p3.ToHTML()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected = "<div><b>Bold text</b><p><i>Italic text</i>Normal text</p></div>"
	if html != expected {
		t.Errorf("Expected %s, got %s", expected, html)
	}
}

func TestPropsToHTML(t *testing.T) {
	props := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	expected := `key1="value1" key2="value2"`
	if PropsToHTML(props) != expected {
		t.Errorf("Expected %s, got %s", expected, PropsToHTML(props))
	}
}

func TestTextNodeToHTMLNode(t *testing.T) {
	tn := TextNode{
		Type: Link,
		Text: "This is some anchor text",
		URL:  "https://www.boot.dev",
	}

	node, err := TextNodeToHTMLNode(tn)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := `<a href="https://www.boot.dev">This is some anchor text</a>`
	html, err := node.ToHTML()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if html != expected {
		t.Errorf("Expected %s, got %s", expected, html)
	}
}
