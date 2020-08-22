package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type Crawler struct {
	safeCrawledMap map[string]bool
	mux            sync.Mutex
}

func (c *Crawler) addUrlToCrawled(url string) {
	c.mux.Lock()
	c.safeCrawledMap[url] = true
	c.mux.Unlock()
}

func (c *Crawler) hasUrlBeenVisited(url string) bool {
	c.mux.Lock()
	_, ok := c.safeCrawledMap[url]
	c.mux.Unlock()
	return ok
}

func (c *Crawler) Crawl(url string, depth int, fetcher Fetcher) {
	var wg sync.WaitGroup
	if url == "" {
		return
	}
	if visitedAlready := c.hasUrlBeenVisited(url); visitedAlready || depth < 1 {
		return
	}
	fmt.Printf("scraping: %s, at depth: %d\n", url, depth)

	body, urls, err := fetcher.Fetch(url)
	c.addUrlToCrawled(url)

	if err != nil {
		fmt.Println(err)
		fmt.Println()
		return
	}

	fmt.Printf("URL: %s, Body: %s\n", url, body)
	fmt.Printf("scraped URLS: %v\n", urls)
	fmt.Println()

	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			c.Crawl(url, depth-1, fetcher)
			wg.Done()
		}(u)
	}
	wg.Wait()
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher fakeFetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
			"http://golang.org/os/car",
		},
	},
	"http://golang.org/os/car": &fakeResult{
		"oscar is king",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/car/hmmmm",
		},
	},
	"http://golang.org/pkg/os/car/hmmmm": &fakeResult{
		"oscar is king",
		[]string{},
	},
}

func main() {
	var mux sync.Mutex
	var c *Crawler = &Crawler{
		safeCrawledMap: make(map[string]bool),
		mux:            mux,
	}
	c.Crawl("http://golang.org/", 5, fetcher)
}
