package blocks

import "testing"

// Unit tests for MarkdownToBlocks
func TestMarkdownToBlocks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:  "Basic markdown",
			input: "# This is a heading\n\nThis is a paragraph of text. It has some **bold** and *italic* words inside of it.\n\n* This is the first list item in a list block\n* This is a list item\n* This is another list item",
			expected: []string{
				"# This is a heading",
				"This is a paragraph of text. It has some **bold** and *italic* words inside of it.",
				"* This is the first list item in a list block\n* This is a list item\n* This is another list item",
			},
		},
		{
			name:  "Markdown with excessive empty lines",
			input: "# This is a heading\n\n\n\nThis is a paragraph of text. It has some **bold** and *italic* words inside of it.\n\n\n* This is the first list item in a list block\n* This is a list item\n* This is another list item",
			expected: []string{
				"# This is a heading",
				"This is a paragraph of text. It has some **bold** and *italic* words inside of it.",
				"* This is the first list item in a list block\n* This is a list item\n* This is another list item",
			},
		},
		{
			name: "Markdown with leading/trailing whitespace",
			input: `
            # This is a heading
         
            This is a paragraph of text. It has some **bold** and *italic* words inside of it.
            
            * This is the first list item in a list block
            * This is a list item
            * This is another list item
        `,
			expected: []string{
				"# This is a heading",
				"This is a paragraph of text. It has some **bold** and *italic* words inside of it.",
				"* This is the first list item in a list block\n* This is a list item\n* This is another list item",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MarkdownToBlocks(test.input)
			if len(result) != len(test.expected) {
				t.Errorf("Expected %d blocks, got %d", len(test.expected), len(result))
			}
			for i := range result {
				if result[i] != test.expected[i] {
					t.Errorf("Block %d: expected %q, got %q", i, test.expected[i], result[i])
				}
			}
		})
	}
}
