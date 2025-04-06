package nodes

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// HTMLNode is an interface defining a node that can be rendered to HTML.
type HTMLNode interface {
	ToHTML() (string, error)
}

// propsToHTML converts a map of attributes to an HTML attribute string.
func propsToHTML(props map[string]string) string {
	keys := make([]string, 0, len(props))
	for key := range props {
		keys = append(keys, key)
	}
	sort.Strings(keys) // Sort keys for consistent order

	result := ""
	for _, key := range keys {
		result += key + `="` + props[key] + `" `
	}
	return strings.TrimSpace(result)
}

// LeafNode represents an HTML node with a tag and a value.
// For a node like a text node (without a tag) tag can be empty.
type LeafNode struct {
	Tag   string
	Value string
	Props map[string]string
}

// ToHTML renders a LeafNode as an HTML string.
func (l *LeafNode) ToHTML() (string, error) {
	// If this is a text node without any tag, return its value
	if l.Tag == "" {
		return l.Value, nil
	}
	// Ensure that value is not empty
	if l.Value == "" {
		return "", errors.New("the value of a leaf node cannot be empty")
	}
	attrStr := propsToHTML(l.Props)
	if attrStr != "" {
		return fmt.Sprintf("<%s %s>%s</%s>", l.Tag, attrStr, l.Value, l.Tag), nil
	}
	return fmt.Sprintf("<%s>%s</%s>", l.Tag, l.Value, l.Tag), nil
}

// ParentNode represents an HTML node that contains children.
type ParentNode struct {
	Tag      string
	Children []HTMLNode
	Props    map[string]string
}

// ToHTML renders a ParentNode as an HTML string.
func (p *ParentNode) ToHTML() (string, error) {
	if p.Tag == "" {
		return "", errors.New("the tag of a parent node cannot be empty")
	}
	if len(p.Children) == 0 {
		return "", errors.New("the children of a parent node cannot be nil or empty")
	}

	attrStr := propsToHTML(p.Props)
	var result string
	if attrStr != "" {
		result = fmt.Sprintf("<%s %s>", p.Tag, attrStr)
	} else {
		result = fmt.Sprintf("<%s>", p.Tag)
	}

	for _, child := range p.Children {
		childHTML, err := child.ToHTML()
		if err != nil {
			return "", err
		}
		result += childHTML
	}

	result += fmt.Sprintf("</%s>", p.Tag)
	return result, nil
}

// textNodeToHTMLNode converts a TextNode to an HTMLNode.
func textNodeToHTMLNode(tn TextNode) (HTMLNode, error) {
	switch tn.Type {
	case Normal:
		// Leaf node without a tag, just plain text.
		return &LeafNode{
			Tag:   "",
			Value: tn.Text,
			Props: nil,
		}, nil
	case Bold:
		return &LeafNode{
			Tag:   "b",
			Value: tn.Text,
			Props: nil,
		}, nil
	case Italic:
		return &LeafNode{
			Tag:   "i",
			Value: tn.Text,
			Props: nil,
		}, nil
	case Code:
		return &LeafNode{
			Tag:   "code",
			Value: tn.Text,
			Props: nil,
		}, nil
	case Link:
		return &LeafNode{
			Tag:   "a",
			Value: tn.Text,
			Props: map[string]string{
				"href": tn.URL,
			},
		}, nil
	case Image:
		// For an image, the value is empty (or could be used as inner text for fallback)
		return &LeafNode{
			Tag:   "img",
			Value: "",
			Props: map[string]string{
				"src": tn.URL,
				"alt": tn.Text,
			},
		}, nil
	default:
		return nil, fmt.Errorf("invalid text type: %v", tn.Type)
	}
}
