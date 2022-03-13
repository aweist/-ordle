package html

import (
	"strings"

	"golang.org/x/net/html"
)

func FindNodeByAttr(root *html.Node, key, value string) (results []*html.Node) {
	var f func(*html.Node)
	f = func(node *html.Node) {
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

func GetAttr(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func nodeValue(cell *html.Node) byte {
	if cell.FirstChild == nil {
		return ' '
	}
	b := strings.ToLower(cell.FirstChild.Data)[0]
	return b
}
