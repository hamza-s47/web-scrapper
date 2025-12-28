package src

type ScrapedData struct {
	Links    []string `json:"links"`
	Headings []string `json:"headings"`
}
type Result struct {
	URL           string      `json:"url"`
	StatusCode    int         `json:"statusCode"`
	LinksCount    int         `json:"linksCount"`
	HeadingsCount int         `json:"headingsCount"`
	Data          ScrapedData `json:"data"`
	Err           error       `json:"err"`
}
type Urls struct {
	URLs []string `json:"urls"`
}
