package main

import (
	"fmt"

	"github.com/Fepozopo/Heirloom-Kitchen/src"
)

func main() {
	node1 := src.TextNode{
		Type: src.LinkText,
		Text: "This is some anchor text",
		URL:  "https://www.boot.dev",
	}

	node2 := src.TextNode{
		Type: src.BoldText,
		Text: "This is some bold text",
		URL:  "",
	}

	if node1.Equals(node2) {
		fmt.Println("The nodes are equal")
	} else {
		fmt.Println("The nodes are not equal")
	}

	fmt.Println("Node 1:", node1.Repr())
	fmt.Println("Node 2:", node2.Repr())
}
