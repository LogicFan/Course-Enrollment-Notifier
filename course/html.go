package course

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Node same as html.Node
type Node struct {
	node *html.Node
}

// ParseHTML warpper of html.Parse
func ParseHTML(r io.Reader) (Node, error) {
	node, err := html.Parse(r)
	return Node{node: node}, err
}

// Children find the children of current node, an empty node is omitted
func (node Node) Children() Node {
	if node.node == nil {
		return Node{node: nil}
	} else if node.node.FirstChild == nil {
		return Node{node: nil}
	} else if strings.TrimSpace(node.node.FirstChild.Data) == "" {
		return Node{node: node.node.FirstChild}.Next()
	} else {
		return Node{node: node.node.FirstChild}
	}
}

// Next find the next sibling of current node, an empty node is omitted
func (node Node) Next() Node {
	if node.node == nil {
		return Node{node: nil}
	} else if node.node.NextSibling == nil {
		return Node{node: nil}
	} else if strings.TrimSpace(node.node.NextSibling.Data) == "" {
		return Node{node: node.node.NextSibling}.Next()
	} else {
		return Node{node: node.node.NextSibling}
	}
}

// Nil return true if node contains nil
func (node Node) Nil() bool {
	return node.node == nil
}

// Lexeme return the lexeme of current node
func (node Node) Lexeme() string {
	if node.node == nil {
		return ""
	}

	return strings.TrimSpace(node.node.Data)
}

// Type returns the type of node, same as html.Node
func (node Node) Type() int {
	if node.node == nil {
		return 0
	}

	return int(node.node.Type)
}

// Text return the text of current node and sub-nodes
func (node Node) Text() string {
	return strings.TrimSpace(node.textRec())
}

func (node Node) textRec() string {
	if node.Type() == 1 {
		return node.Lexeme() + " "
	}

	text := ""
	node = node.Children()
	for !node.Nil() {
		text = text + node.textRec()
		node = node.Next()
	}

	return text
}
