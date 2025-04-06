package website

import "testing"

func TestExtractTitleWithH1Header(t *testing.T) {
	markdown := `
    # This is the title
    `

	result, err := ExtractTitle(markdown)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "This is the title"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestExtractTitleWithNoH1Header(t *testing.T) {
	markdown := `
    This is the title
    `

	_, err := ExtractTitle(markdown)
	if err == nil {
		t.Fatal("expected an error but got none")
	}

	expectedError := "no h1 header found"
	if err.Error() != expectedError {
		t.Errorf("expected error %q, got %q", expectedError, err.Error())
	}
}

func TestExtractTitleWithWhitespace(t *testing.T) {
	markdown := `
        # This is a heading
     
        This is a paragraph of text. It has some **bold** and *italic* words inside of it.
        
        * This is the first list item in a list block
        * This is a list item
        * This is another list item
    `

	result, err := ExtractTitle(markdown)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "This is a heading"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
