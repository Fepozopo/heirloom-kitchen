package inline

import (
	"reflect"
	"testing"

	"github.com/Fepozopo/heirloom-kitchen/src/nodes"
)

func TestSplitNodesDelimiter(t *testing.T) {
	// Simple code block delimiter
	textTypeText := nodes.NormalText
	textTypeCode := nodes.CodeText

	node := nodes.TextNode{
		Type: textTypeText,
		Text: "This is text with a `code block` word",
		URL:  "",
	}
	result, err := SplitNodesDelimiter([]nodes.TextNode{node}, "`", textTypeCode)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := []nodes.TextNode{
		{Type: textTypeText, Text: "This is text with a ", URL: ""},
		{Type: textTypeCode, Text: "code block", URL: ""},
		{Type: textTypeText, Text: " word", URL: ""},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Bold text delimiter
	textTypeBold := nodes.BoldText

	node = nodes.TextNode{
		Type: textTypeText,
		Text: "This is **bold** text",
		URL:  "",
	}
	result, err = SplitNodesDelimiter([]nodes.TextNode{node}, "**", textTypeBold)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = []nodes.TextNode{
		{Type: textTypeText, Text: "This is ", URL: ""},
		{Type: textTypeBold, Text: "bold", URL: ""},
		{Type: textTypeText, Text: " text", URL: ""},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Italic text delimiter
	textTypeItalic := nodes.ItalicText

	node = nodes.TextNode{
		Type: textTypeText,
		Text: "This is *italic* text",
		URL:  "",
	}
	result, err = SplitNodesDelimiter([]nodes.TextNode{node}, "*", textTypeItalic)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = []nodes.TextNode{
		{Type: textTypeText, Text: "This is ", URL: ""},
		{Type: textTypeItalic, Text: "italic", URL: ""},
		{Type: textTypeText, Text: " text", URL: ""},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Unmatched delimiter (should raise an error)
	node = nodes.TextNode{
		Type: textTypeText,
		Text: "This is unmatched `code block text",
		URL:  "",
	}
	_, err = SplitNodesDelimiter([]nodes.TextNode{node}, "`", textTypeCode)
	if err == nil {
		t.Errorf("Expected error, got none")
	}

	// No splitting needed (no delimiter present)
	node = nodes.TextNode{
		Type: textTypeText,
		Text: "No delimiter here",
		URL:  "",
	}
	result, err = SplitNodesDelimiter([]nodes.TextNode{node}, "`", textTypeCode)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = []nodes.TextNode{node}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Mixed nodes (only "text" nodes are split)
	node1 := nodes.TextNode{
		Type: textTypeText,
		Text: "Normal text",
		URL:  "",
	}
	node2 := nodes.TextNode{
		Type: textTypeBold,
		Text: "**bold** text",
		URL:  "",
	}
	node3 := nodes.TextNode{
		Type: textTypeText,
		Text: "Code `inline` text",
		URL:  "",
	}
	result, err = SplitNodesDelimiter([]nodes.TextNode{node1, node2, node3}, "`", textTypeCode)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = []nodes.TextNode{
		node1,
		node2,
		{Type: textTypeText, Text: "Code ", URL: ""},
		{Type: textTypeCode, Text: "inline", URL: ""},
		{Type: textTypeText, Text: " text", URL: ""},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
