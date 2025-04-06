package blocks

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/Fepozopo/culinary-keepsakes/src/inline"
	"github.com/Fepozopo/culinary-keepsakes/src/nodes"
)

type BlockType int

const (
	Paragraph BlockType = iota
	Heading
	Code
	Quote
	UnorderedList
	OrderedList
)

// String method to return a string representation of the BlockType object
func (bt BlockType) String() string {
	typeStr := ""
	switch bt {
	case Paragraph:
		typeStr = "paragraph"
	case Heading:
		typeStr = "heading"
	case Code:
		typeStr = "code"
	case Quote:
		typeStr = "quote"
	case UnorderedList:
		typeStr = "unordered_list"
	case OrderedList:
		typeStr = "ordered_list"
	default:
		typeStr = "unknown"
	}
	return typeStr
}

// MarkdownToBlocks takes a raw markdown string and returns a list of "block" strings.
func MarkdownToBlocks(markdown string) []string {
	// Strip any leading or trailing whitespace from the entire document
	markdown = strings.TrimSpace(markdown)

	// Normalize the markdown by removing leading spaces from each line
	lines := strings.Split(markdown, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimLeft(line, " ")
	}

	// Merge consecutive lines starting with "> " into a single block
	mergedLines := []string{}
	currentBlock := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "> ") {
			// If the line starts with "> ", append it to the current block
			if currentBlock != "" {
				currentBlock += "\n" + line
			} else {
				currentBlock = line
			}
		} else {
			// If the line doesn't start with "> ", finalize the current block
			if currentBlock != "" {
				mergedLines = append(mergedLines, currentBlock)
				currentBlock = ""
			}
			mergedLines = append(mergedLines, line)
		}
	}
	// Add the last block if it exists
	if currentBlock != "" {
		mergedLines = append(mergedLines, currentBlock)
	}

	// Join the merged lines back into a single string
	markdown = strings.Join(mergedLines, "\n")

	// Remove any "empty" blocks due to excessive newlines
	re := regexp.MustCompile(`\n{3,}`)
	markdown = re.ReplaceAllString(markdown, "\n\n")

	// Split the markdown into blocks based on double newlines
	blocks := strings.Split(markdown, "\n\n")

	// Strip leading/trailing whitespace from each block individually
	for i, block := range blocks {
		blocks[i] = strings.TrimSpace(block)
	}

	return blocks
}

// BlockToBlockType takes a single block of markdown text as input and returns the BlockType it represents.
func BlockToBlockType(block string) BlockType {
	// Check for headings (1-6 # followed by a space)
	if matched, _ := regexp.MatchString(`^#{1,6} .+`, block); matched {
		return Heading
	}

	// Split the block into lines for further processing
	lines := strings.Split(block, "\n")

	// Check if the block is an unordered list. Every line in an unordered list block must start with *, -, or + followed by a space.
	unorderedList := true
	unorderedListRegex := regexp.MustCompile(`^(\*|\-|\+)\s`)
	for _, line := range lines {
		if !unorderedListRegex.MatchString(line) {
			unorderedList = false
			break
		}
	}
	if unorderedList {
		return UnorderedList
	}

	// Check if the block is an ordered list. Every line in an ordered list block must start with a number followed by a . and a space.
	orderedList := true
	for i, line := range lines {
		match := regexp.MustCompile(`^(\d+)\.\s`).FindStringSubmatch(line)
		if match == nil || match[1] != strconv.Itoa(i+1) {
			orderedList = false
			break
		}
	}
	if orderedList {
		return OrderedList
	}

	// Check if the block is a code block. Code blocks must start and end with 3 backticks.
	if strings.HasPrefix(block, "```") && strings.HasSuffix(block, "```") {
		return Code
	}

	// Check if the block is a quote. Every line in a quote block must start with > followed by a space.
	quote := true
	for _, line := range lines {
		if !strings.HasPrefix(line, "> ") {
			quote = false
			break
		}
	}
	if quote {
		return Quote
	}

	// If none of the above, the block is a paragraph
	return Paragraph
}

// TextToChildren converts a string of text into a list of HTMLNode objects.
func TextToChildren(text string) []nodes.HTMLNode {
	textNodes := inline.TextToTextNodes(text)
	htmlNodes := []nodes.HTMLNode{}
	for _, textNode := range textNodes {
		htmlNode, _ := nodes.TextNodeToHTMLNode(textNode)
		htmlNodes = append(htmlNodes, htmlNode)
	}
	return htmlNodes
}

// BlockToHTMLNode converts a block of markdown into an appropriate HTMLNode.
func BlockToHTMLNode(block string) nodes.HTMLNode {
	blockType := BlockToBlockType(block)

	switch blockType {
	case Heading:
		level := len(strings.SplitN(block, " ", 2)[0]) // Count the number of # symbols
		textContent := strings.TrimSpace(block[level+1:])
		return &nodes.ParentNode{
			Tag:      "h" + strconv.Itoa(level),
			Children: TextToChildren(textContent),
		}

	case UnorderedList:
		items := []nodes.HTMLNode{}
		for _, line := range strings.Split(block, "\n") {
			items = append(items, &nodes.ParentNode{
				Tag:      "li",
				Children: TextToChildren(strings.TrimSpace(line[2:])),
			})
		}
		return &nodes.ParentNode{
			Tag:      "ul",
			Children: items,
		}

	case OrderedList:
		items := []nodes.HTMLNode{}
		for _, line := range strings.Split(block, "\n") {
			items = append(items, &nodes.ParentNode{
				Tag:      "li",
				Children: TextToChildren(strings.TrimSpace(line[3:])),
			})
		}
		return &nodes.ParentNode{
			Tag:      "ol",
			Children: items,
		}

	case Code:
		codeContent := strings.TrimSpace(block[3 : len(block)-3])
		return &nodes.ParentNode{
			Tag: "pre",
			Children: []nodes.HTMLNode{
				&nodes.LeafNode{
					Tag:   "code",
					Value: codeContent,
					Props: nil,
				},
			},
		}

	case Quote:
		// Preserve line breaks within the blockquote
		lines := strings.Split(block, "\n")
		quoteContent := []nodes.HTMLNode{}
		currentParagraph := ""

		for _, line := range lines {
			trimmedLine := strings.TrimSpace(line[2:]) // Remove "> " prefix
			if trimmedLine == "" {
				// If the line is blank, finalize the current paragraph
				if currentParagraph != "" {
					quoteContent = append(quoteContent, &nodes.LeafNode{
						Tag:   "span",
						Value: currentParagraph,
						Props: nil,
					})
					quoteContent = append(quoteContent, &nodes.LeafNode{
						Tag:   "br",
						Value: "",
						Props: nil,
					})
					currentParagraph = ""
				}
			} else {
				// Append the line to the current paragraph
				if currentParagraph != "" {
					currentParagraph += " " + trimmedLine
				} else {
					currentParagraph = trimmedLine
				}
			}
		}

		// Add the last paragraph if it exists
		if currentParagraph != "" {
			quoteContent = append(quoteContent, &nodes.LeafNode{
				Tag:   "span",
				Value: currentParagraph,
				Props: nil,
			})
		}

		return &nodes.ParentNode{
			Tag:      "blockquote",
			Children: quoteContent,
		}

	default:
		return &nodes.ParentNode{
			Tag:      "p",
			Children: TextToChildren(block),
		}
	}
}

// MarkdownToHTMLNode converts a full markdown document into a single HTMLNode.
func MarkdownToHTMLNode(markdown string) nodes.HTMLNode {
	blocks := MarkdownToBlocks(markdown)
	parentNode := &nodes.ParentNode{
		Tag:      "div",
		Children: []nodes.HTMLNode{},
	}

	for _, block := range blocks {
		blockHTMLNode := BlockToHTMLNode(block)
		parentNode.Children = append(parentNode.Children, blockHTMLNode)
	}

	return parentNode
}
