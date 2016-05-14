package gSearch

type googleSearchResponse struct {
	ResponseData    googleResponseData `json: "responseData"`
	ResponseDetails string             `json: "responseDetails"`
	ResponseStatus  float64            `json: "responseStatus"`
}

type googleResponseData struct {
	Results []googleResults `json: "results"`
	Cursor  googleCursor    `json: "cursor"`
}

type googleResults struct {
	GsearchResultClass string `json: "GsearchResultClass"`
	UnescapedUrl       string `json: "unescapedUrl"`
	Url                string `json: "url"`
	VisibleUrl         string `json: "visibleUrl"`
	CacheUrl           string `json: "cacheUrl"`
	Title              string `json: "title"`
	TitleNoForm        string `json: "titleNoFormatting"`
	Content            string `json: "content"`
}

type googleCursor struct {
	ResultCount          string        `json: "resultCount"`
	Pages                []googlePages `json: "pages"`
	EstimatedResultCount string        `json: "estimatedResultCount"`
	CurrentPageIndex     float64       `json: "currentPageIndex"`
	MoreResultsUrl       string        `json: "moreResultsUrl"`
	SearchResultTime     string        `json: "searchResultTime"`
}

type googlePages struct {
	Start string  `json: "start"`
	Label float64 `json: "label"`
}
