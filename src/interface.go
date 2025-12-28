package src

type ScrapedData struct {
	Links    []string `json:"links"`
	Headings []string `json:"headings"`
}
type Result struct {
	URL        string      `json:"url"`
	StatusCode int         `json:"statusCode"`
	Data       ScrapedData `json:"data"`
	Err        error       `json:"err"`
}
type Urls struct {
	URLs []string `json:"urls"`
}
