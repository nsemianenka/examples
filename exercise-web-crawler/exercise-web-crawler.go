package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type UrlsBodyMapping map[string]string
type Crawler struct {
	m UrlsBodyMapping
	mux sync.Mutex
}

func (c * Crawler) crawl (url string, depth int, fetcher Fetcher, ch chan string) {
	if depth < 0 {
		close(ch)
		return
	}
	c.mux.Lock()
	if _, ok := c.m[url]; ok {
		c.mux.Unlock()
		close(ch)
		return
	} else {
		body, urls, err := fetcher.Fetch(url)

		if err != nil {
			c.mux.Unlock()
			c.m[url] = err.Error()
			ch <- err.Error()
			//fmt.Println(err)
			close(ch)
			return
		}
		c.m[url] = body
		c.mux.Unlock()
		ch <- fmt.Sprintf("found: %s %q\n", url, body)
		chans := make([]chan string, len(urls))
		for i := range chans {
			chans[i] = make(chan string)
		}

		for i, u := range urls {
			go c.crawl(u, depth-1, fetcher, chans[i])
		}

		for _, chr := range chans {
			for c := range chr {
				ch <- c
			}
		}
		close(ch)
	}

}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.

	c := Crawler{m: make(map[string]string)}
	ch := make(chan string)
	go c.crawl(url, depth, fetcher, ch)
	for res := range ch {
		fmt.Print(res)
	}
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s\n", url)
}

// fetcher is a populated fakeFetcher.
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
		},
	},
}