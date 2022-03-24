package html

import (
	"strings"

	"golang.org/x/net/html"
)

type Node = html.Node

var Parse = html.Parse

func FindNodeByAttr(root *Node, key, value string) (results []*Node) {
	var f func(*Node)
	f = func(node *Node) {
		for _, a := range node.Attr {
			if a.Key == key && a.Val == value {
				results = append(results, node)
			}
		}
		for n := node.FirstChild; n != nil; n = n.NextSibling {
			f(n)
		}
	}
	f(root)
	return
}

func GetAttr(node *Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func NodeValue(cell *Node) byte {
	if cell.FirstChild == nil {
		return ' '
	}
	data := cell.FirstChild.Data
	data = strings.ToLower(data)
	data = strings.TrimSpace(data)
	b := data[0]
	return b
}
