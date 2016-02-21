package main

import (
	"fmt"
	"sync"
)

type FetchedUrls struct {
	fetched map[string]bool
	mux     sync.Mutex
}

var fetchedUrls FetchedUrls

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	fetchedUrls = FetchedUrls{
		fetched: make(map[string]bool),
	}
	ch := make(chan int)
	go crawl(url, depth, fetcher, ch)
	for range ch {}
}

func crawl(url string, depth int, fetcher Fetcher, ch chan int) {
	fetchedUrls.mux.Lock()
	if _, exists := fetchedUrls.fetched[url]; exists {
		close(ch)
		fetchedUrls.mux.Unlock()
		return
	}
	fetchedUrls.fetched[url] = true
	fetchedUrls.mux.Unlock()
	if depth <= 0 {
		close(ch)
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		close(ch)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	subchs := make([]chan int, len(urls))
	for i, u := range urls {
		subch := make(chan int)
		subchs[i] = subch
		go crawl(u, depth-1, fetcher, subch)
	}
	for _, subch := range subchs {
		for range subch {}
	}
	close(ch)
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
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

var fetcher = fakeFetcher{
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
			"http://golang.org/fmt/",
			"http://golang.org/os/",
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
		},
	},
}
