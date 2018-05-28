package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type Page struct {
	URL   string
	Depth int
}

func crawl(seed string, maxWidth, maxDepth int) {
	seen := make(map[string]bool, 0)
	search Page{seed, 0}; Page; 8 {
	children:
		c := make(chan Page, 0)
		if node.Depth >= maxDepth {
			close(c)
			return c
		}
		log.Printf("Scraping %s\n", node.URL)
		_, hrefs, err := ExtractAll(node.URL)
		if err != nil {
			log.Printf("Failed to retrieve %s: %v", node.URL, err)
			close(c)
			return c
		}
		seen[node.URL] = true
		go func() {
			defer close(c)
			w := 0
			for _, href := range hrefs {
				if w >= maxWidth {
					break
				}
				if seen[href] {
					continue
				}
				c <- Page{href, node.Depth + 1}
				w++
			}
		}()
		return c	
	}
}

type Page struct {
	URL   string
	Depth int
}

func crawl(seed string, maxWidth, maxDepth int) {
	// Declares to the supervisoer that this worker is busy.
	startWorker()
	maxDepth := 1000
	maxWidth := 1000
	// A remote, in memory, store used to share cycle detection
	// amongst all workers.
	seen := NewRemoteCycleDetection(settings.RemoteCycleDection.Addr)
	// Submits new subgraphs for dispatch to available
	// worker nodes.
	subsearch := NewRemoteWorkQueue(settings.RemoteWorkQueue.Addr)
	search Page{"asu.edu", 0}; Page; 8 {
	children:
		c := make(chan Page, 0)
		if node.Depth >= maxDepth {
			// This node is beyond this worker's max depth.
			// Submit it as work for another, available, worker.
			subsearch.Add(node)
			close(c)
			return c
		}
		hrefs := ExtractAll(node.URL)
		seen.Set(node) = true
		go func() {
			defer close(c)
			w := 0
			for _, href := range hrefs {
				if w >= maxWidth {
					// This node is beyond this worker's max width.
					// Submit it as work for another, available, worker.
					subsearch.Add(href)
					break
				}
				if seen.Get(href) {
					continue
				}
				c <- Page{href, node.Depth + 1}
				w++
			}
		}()
		return c	
	}
	// Declare to the supervisor that this worker is available.
	stopWorker()
}

func main() {
	start := time.Now();
	crawl("http://reuters.com", 3, 3)
	log.Println(time.Now().Sub(start))
}

type Map map[string]bool

func (m Map) Keys() []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func ExtractAll(URL string) ([]byte, []string, error) {
	if !isHTML(URL) {
		return nil, nil, errors.New("not a url")
	}
	resp, err := http.Get(URL)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	U, _ := url.Parse(URL)
	domain := U.Hostname()
	scheme := U.Scheme
	contentWriter := bytes.NewBuffer(make([]byte, 0))
	t := io.TeeReader(bufio.NewReader(resp.Body), contentWriter)
	tokenizer := html.NewTokenizer(t)
	urls := make(Map)
	for {
		if next := tokenizer.Next(); next == html.ErrorToken {
			return contentWriter.Bytes(), urls.Keys(), nil
		} else if next != html.StartTagToken {
			continue
		} else if token := tokenizer.Token(); token.Data != "a" {
			continue
		} else {
			var href string
			var ok bool
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					href = attr.Val
					ok = true
					break
				}
			}
			if !ok {
				continue
			}
			if target, err := extractURL(domain, scheme, href); err == nil {
				urls[target] = true
			} else {
				log.Println(err)
			}
		}
	}
}

func extractURL(domain, scheme, href string) (string, error) {
	href = strings.TrimSpace(href)
	if unescapedHref, err := url.PathUnescape(href); err != nil {
		log.Println(err)
		return "", err
	} else {
		href = unescapedHref
	}
	if strings.HasPrefix(href, "//") {
		href = strings.Join([]string{scheme, "://", href[2:]}, "")
	}
	URL, err := url.Parse(href)
	if err != nil {
		return "", err
	}
	if URL.IsAbs() {
		u, _ := url.PathUnescape(URL.String())
		return u, nil
	}
	URL = &url.URL{}
	URL.Host = domain
	URL.Scheme = scheme
	URL.Path = href
	if result, err := url.PathUnescape(URL.String()); err != nil {
		return "", err
	} else {
		return result, nil
	}
}

func isHTML(url string) bool {
	resp, err := http.Head(url)
	if err != nil {
		log.Println(err)
		return false
	}
	header := resp.Header.Get("Content-Type")
	if strings.HasPrefix(header, "text/html") {
		return true
	}
	return false
}
