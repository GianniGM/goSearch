package main

type googleSearchResponse struct {
	Response googleResponseData `json: "responseData"`
	Details  string             `json: "responseDetails"`
	Status   int64              `json: "responseStatus"`
}

type googleResponseData struct {
	results []googleResults `json: "results"`
	cursor  googleCursor    `json: "cursor"`
}

type googleResults struct {
	resultClass string `json: "GsearchResultClass"`
	uURL        string `json: "unescapedUrl"`
	URL         string `json: "url"`
	vURL        string `json: "visibleUrl"`
	cURL        string `json: "cacheUrl"`
	title       string `json: "title"`
	titleNoForm string `json: "titleNoFormatting"`
	content     string `json: "json: content"`
}

type googleCursor struct {
	resultCount          int64         `json: "resultCount"`
	pages                []googlePages `json: "pages"`
	estimatedResultCount int64         `json: "estimatedResultCount"`
	currentPageIndex     int64         `json: "currentPageIndex"`
	moreResultsUrl       string        `json: "moreResultsUrl"`
	searchResultTime     float32       `json: "searchResultTime"`
}

type googlePages struct {
	start int `json: "start"`
	label int `json: "label"`
}
