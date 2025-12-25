package src

type Result struct {
	URL        string
	StatusCode int
	Err        error
}

type Urls struct {
	URLs []string `json:"urls"`
}
