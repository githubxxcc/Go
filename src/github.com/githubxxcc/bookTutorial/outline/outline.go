package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in parsing: %v", err)
		os.Exit(1)
	}

	forEachNode(doc, startElement, endElement)
}

func forEachNode(doc *html.Node, f1, f2 func(doc *html.Node)) {
	if f1 != nil {
		f1(doc)
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, f1, f2)
	}

	if f2 != nil {
		f2(doc)
	}
}

var depth int

func startElement(doc *html.Node) {
	if doc.Type == html.ElementNode {
		fmt.Printf("%*s<%s>", depth*2, " ", doc.Data)
		depth++

		fmt.Print("<")

		for _, v := range doc.Attr {
			fmt.Printf("%v", v.Val)
		}

		fmt.Println(">")
	}

}

func endElement(doc *html.Node) {
	if doc.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth*2, " ", doc.Data)
		depth--
	}
}

// func outline(stack []string, doc *html.Node) {
// 	if doc.Type == html.ElementNode {
// 		stack = append(stack, doc.Data)
// 		fmt.Println(stack)
// 	}

// 	for c := doc.FirstChild; c != nil; c = c.NextSibling {
// 		outline(stack, c)
// 	}
// }
