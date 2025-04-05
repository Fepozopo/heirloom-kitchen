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

func TestExtractMarkdownImages(t *testing.T) {
	t.Run("extract_markdown_images", func(t *testing.T) {
		text := "This is text with a ![rick roll](https://i.imgur.com/aKaOqIh.gif) and ![obi wan](https://i.imgur.com/fJRm4Vk.jpeg)"
		result := ExtractMarkdownImages(text)
		expected := [][2]string{{"rick roll", "https://i.imgur.com/aKaOqIh.gif"}, {"obi wan", "https://i.imgur.com/fJRm4Vk.jpeg"}}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("extractMarkdownImages() = %v, want %v", result, expected)
		}
	})

	t.Run("extract_markdown_images_with_no_images", func(t *testing.T) {
		text := "This is text with no images"
		result := ExtractMarkdownImages(text)
		expected := [][2]string{}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("extractMarkdownImages() = %v, want %v", result, expected)
		}
	})
}

func TestExtractMarkdownLinks(t *testing.T) {
	t.Run("extract_markdown_links", func(t *testing.T) {
		text := "This is text with a link [to boot dev](https://www.boot.dev) and [to youtube](https://www.youtube.com/@bootdotdev)"
		result := ExtractMarkdownLinks(text)
		expected := [][2]string{{"to boot dev", "https://www.boot.dev"}, {"to youtube", "https://www.youtube.com/@bootdotdev"}}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("extractMarkdownLinks() = %v, want %v", result, expected)
		}
	})

	t.Run("extract_markdown_links_with_no_links", func(t *testing.T) {
		text := "This is text with no links"
		result := ExtractMarkdownLinks(text)
		expected := [][2]string{}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("extractMarkdownLinks() = %v, want %v", result, expected)
		}
	})
}
