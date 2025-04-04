package main

import (
	"fmt"

	"github.com/Fepozopo/heirloom-kitchen/src/nodes"
)

func main() {
	node1 := nodes.TextNode{
		Type: nodes.LinkText,
		Text: "This is some anchor text",
		URL:  "https://www.boot.dev",
	}

	node2 := nodes.TextNode{
		Type: nodes.BoldText,
		Text: "This is some bold text",
		URL:  "",
	}

	if node1.Equals(node2) {
		fmt.Println("The nodes are equal")
	} else {
		fmt.Println("The nodes are not equal")
	}

	fmt.Println("Node 1:", node1.String())
	fmt.Println("Node 2:", node2.String())
}
