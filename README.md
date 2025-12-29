# Concurrent Web Scraper

A concurrent web scraper built in Go that extracts links and headings from multiple URLs using goroutines, worker pools, and channels.
Supports both CLI and REST API usage.

## Features:
- Concurrent scraping with worker pool
- Extracts links (href) and headings (h1–h6)
- CLI mode (terminal-based)
- API mode (Gin REST API)
- JSON input & output
- Clean, modular architecture
- Production-ready

## Tech Stack:
- Go
- Gin
- Goroutines & Channels
- net/http
- goquery

## Usage
- git clone https://github.com/hamza-s47/web-scrapper.git
- cd your-repo-name
- go mod tidy
- go build -o app (if needed otherwise it is already built)
- ./app --help (for helping instruction)

### For CLI
./app https://example1.com https://example2.com .....

### For API
./app --api (will run on port:8080)


## POST (/urls)

{
  "urls": ["https://example.com", "https://github.com"]
}

## Output
{
  "url": "https://example.com",
  "statusCode": 200,
  "linksCount": 25,
  "headingsCount": 6,
  "data": {
    "links": [...],
    "headings": [...]
  }
}

## Extensible:
Additional scraping features (metadata, text, images, etc.) can be added too as per need.
