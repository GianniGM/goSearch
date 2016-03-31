package main

type googleSearchResponse struct {
	Response googleResponseData `json: "responseData"`
	Details  string             `json: "responseDetails"`
	Status   float64            `json: "responseStatus"`
}

type googleResponseData struct {
	Results []googleResults `json: "results"`
	Cursor  googleCursor    `json: "cursor"`
}

type googleResults struct {
	resultClass string `json: "GsearchResultClass"`
	uURL        string `json: "unescapedUrl"`
	URL         string `json: "url"`
	vURL        string `json: "visibleUrl"`
	cURL        string `json: "cacheUrl"`
	title       string `json: "title"`
	titleNoForm string `json: "titleNoFormatting"`
	content     string `json: "content"`
}

type googleCursor struct {
	resultCount          string        `json: "resultCount"`
	pages                []googlePages `json: "pages"`
	estimatedResultCount string        `json: "estimatedResultCount"`
	currentPageIndex     string        `json: "currentPageIndex"`
	moreResultsUrl       string        `json: "moreResultsUrl"`
	searchResultTime     string        `json: "searchResultTime"`
}

type googlePages struct {
	start string `json: "start"`
	label int    `json: "label"`
}
