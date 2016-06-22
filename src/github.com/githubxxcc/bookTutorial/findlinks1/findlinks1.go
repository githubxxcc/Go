package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	var list []string
	for _, url := range os.Args[1:] {
		a, err := Extract(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "error %s", err)
		}

		list = append(list, a...)
	}

	for _, link := range list {
		fmt.Println(link)
	}
}

// func visit(list []string, node *html.Node) []string {
// 	if node.Type == html.ElementNode && node.Data == "a" {
// 		for _, a := range node.Attr {
// 			if a.Key == "href" {
// 				list = append(list, a.Val)
// 			}
// 		}
// 	}

// 	for c := node.FirstChild; c != nil; c = c.NextSibling {
// 		list = visit(list, c)
// 	}

// 	return list
// }

func forEachNode(n *html.Node, f1, f2 func(n *html.Node)) {
	if f1 != nil {
		f1(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, f1, f2)
	}

	if f2 != nil {
		f2(n)
	}

}

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("error %s", url)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("err in parsing %s", url)
	}

	var links []string

	visitAll := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Val != "href" {
					continue
				}

				link, err := resp.Request.URL.Parse(a.Val)

				if err != nil {
					continue
				}

				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visitAll, nil)
	return links, nil
}
