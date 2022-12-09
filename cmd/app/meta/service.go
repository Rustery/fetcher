package meta

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type Meta struct {
	LinksCount    int
	ImagesCount   int
	LastFetchedAt time.Time
}

func FetchInformation(content []byte) error {
	metaInformation := Meta{}
	doc, err := html.Parse(strings.NewReader(string(content)))
	if err != nil {
		return err
	}

	metaInformation.LinksCount = FindTagQty(doc, "a")
	metaInformation.ImagesCount = FindTagQty(doc, "img")
	metaInformation.LastFetchedAt = time.Now()

	fmt.Printf("Links count:\t\t%d\n", metaInformation.LinksCount)
	fmt.Printf("Images count:\t\t%d\n", metaInformation.ImagesCount)
	fmt.Printf("Last fetched at:\t%s\n", metaInformation.LastFetchedAt.Format(time.RFC1123))

	return nil
}

func FindTagQty(doc *html.Node, tag string) int {
	var qty int
	Finder(doc, func(t *html.Node) {
		if t.Data == tag {
			qty++
		}
	})
	return qty
}

func Finder(n *html.Node, process func(*html.Node)) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			process(c)
		}
		res := Finder(c, process)
		if res != nil {
			return res
		}
	}
	return nil
}
