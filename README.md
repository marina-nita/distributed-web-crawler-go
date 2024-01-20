# Distributed Web Crawler (Go)

🗓️ **January 2024**

A fault-tolerant web crawler built in Go, using:
- Goroutines and channels for concurrency
- A thread-safe URL map for deduplication
- Timeout and retry logic for robustness
- HTML parsing to extract metadata and recursively follow links

## Features
- Concurrent crawling with goroutines
- Link deduplication to avoid cycles
- Timeout and retry on HTTP requests
- HTML parsing with goquery

## Usage

```bash
go run main.go https://example.com
```

