package assets

import (
	"fetcher/cmd/app/files"
	"fetcher/cmd/app/meta"
	"fmt"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func FetchAssets(base *url.URL, content []byte) error {
	doc, err := html.Parse(strings.NewReader(string(content)))
	if err != nil {
		return err
	}
	rawUrls := FindTagsUrls(doc)
	urls := ResolveUrlsToBase(base, rawUrls)
	var wg sync.WaitGroup
	for _, v := range urls {
		wg.Add(1)
		go SaveAsset(&wg, v)
	}
	wg.Wait()
	return nil
}

func SaveAsset(wg *sync.WaitGroup, url *url.URL) {
	defer wg.Done()

	content, err := files.LoadUrlContent(url.String())
	if err != nil {
		fmt.Printf("%s asset loading error: %v\n", url.String(), err)
		return
	}
	err = files.SaveToFile(url, content)
	if err != nil {
		fmt.Printf("%s asset loading error: %v\n", url.String(), err)
		return
	}
}

func ResolveUrlsToBase(base *url.URL, rawUrls []string) []*url.URL {
	urls := []*url.URL{}
	for _, v := range rawUrls {
		u, err := base.Parse(v)
		if err != nil {
			continue
		}
		if strings.HasPrefix(u.Hostname(), base.Hostname()) {
			urls = append(urls, u)
		}
	}
	return urls
}

func FindTagsUrls(doc *html.Node) []string {
	urls := []string{}
	meta.Finder(doc, func(t *html.Node) {
		expectedTags := map[string]string{
			"link":   "href",
			"img":    "src",
			"script": "src",
		}
		for tag, attr := range expectedTags {
			if t.Data == tag {
				for _, a := range t.Attr {
					if a.Key == attr {
						urls = append(urls, a.Val)
						break
					}
				}
				break
			}
		}
	})
	return urls
}
