package fetch

import (
	"fetcher/cmd/app/assets"
	"fetcher/cmd/app/files"
	"fetcher/cmd/app/meta"
	"fmt"
	"net/url"
	"sync"
)

func ProcessUrls(rawUrls []string, metadata bool, loadAssets bool) {
	var wg sync.WaitGroup
	for _, url := range rawUrls {
		wg.Add(1)
		go ProcessUrl(&wg, url, metadata, loadAssets)
	}
	wg.Wait()
}

func ProcessUrl(wg *sync.WaitGroup, rawUrl string, metadata bool, loadAssets bool) {
	defer wg.Done()

	url, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Printf("%s loading error: %v\n", rawUrl, err)
		return
	}

	content, err := files.LoadUrlContent(rawUrl)
	if err != nil {
		fmt.Printf("%s loading error: %v\n", rawUrl, err)
		return
	}

	err = files.SaveToFile(url, content)
	if err != nil {
		fmt.Printf("%s loading error: %v\n", rawUrl, err)
		return
	}

	fmt.Printf("%s loaded successfully\n", rawUrl)

	if metadata {
		err = meta.FetchInformation(content)
		if err != nil {
			fmt.Printf("%s loading error: %v\n", rawUrl, err)
			return
		}
		fmt.Println()
	}

	if loadAssets {
		assets.FetchAssets(url, content)
	}
}
