package main

import (
    "fmt"
    "net/http"
    "time"
    "sync"
    "golang.org/x/net/html"
)

var visited = struct {
    urls map[string]bool
    mux  sync.Mutex
}{urls: make(map[string]bool)}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <url>")
        return
    }

    url := os.Args[1]
    crawl(url, 0)
}

func crawl(url string, depth int) {
    if depth > 3 {
        return
    }

    visited.mux.Lock()
    if visited.urls[url] {
        visited.mux.Unlock()
        return
    }
    visited.urls[url] = true
    visited.mux.Unlock()

    fmt.Println("Crawling:", url)

    client := http.Client{
        Timeout: 5 * time.Second,
    }

    resp, err := client.Get(url)
    if err != nil {
        fmt.Println("Failed:", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("Non-OK HTTP status:", resp.Status)
        return
    }

    tokenizer := html.NewTokenizer(resp.Body)
    for {
        tt := tokenizer.Next()
        if tt == html.ErrorToken {
            break
        }
        token := tokenizer.Token()
        if token.Type == html.StartTagToken && token.Data == "a" {
            for _, attr := range token.Attr {
                if attr.Key == "href" {
                    go crawl(attr.Val, depth+1)
                }
            }
        }
    }
}
