package gSearch

type googleSearchResponse struct {
	Kind              string           `json: "kind"`
	Url               googleUrl        `json: "url"`
	Queries           googleQueries    `json: "queries"`
	Context           googleContext    `json: "context"`
	SearchInformation googleSearchInfo `json: "searchinformation"`
	Items             []googleItems    `json: "items"`
}

type googleUrl struct {
	Type     string `json: "type"`
	Template string `json: "template"`
}

type googleQueries struct {
	Request  []googleRequest `json: "request"`
	NextPage []googleRequest `json: "nextpage"`
}

type googleContext struct {
	Title string `json: "title"`
}

type googleSearchInfo struct {
	SearchTime            float64 `json: "searchTime"`
	FormattedSearchTime   string  `json: "formattedSearchTime"`
	TotalResults          string  `json: "totalResults"`
	FormattedTotalResults string  `json: "formattedTotalResults"`
}

type googleItems struct {
	Kind             string          `json: "kind"`
	Title            string          `json: "title"`
	HtmlTitle        string          `json: "htmlTitle"`
	Link             string          `json: "link"`
	DisplayLink      string          `json: "displayLink"`
	Snippet          string          `json: "snippet"`
	HtmlSnippet      string          `json: "htmlSnippet"`
	CacheId          string          `json: "cacheId"`
	FormattedUrl     string          `json: "formattedUrl"`
	HtmlFormattedUrl string          `json: "htmlFormattedUrl"`
	pagemap          []googlePageMap `json: "pagemap"`
}

type googleRequest struct {
	Title          string  `json: "title"`
	TotalResults   string  `json: "totalResults"`
	SearchTerms    string  `json: "searchTerms"`
	Count          float64 `json: "count"`
	StartIndex     float64 `json: "startIndex"`
	InputEncoding  string  `json: "inputEncoding"`
	OutputEncoding string  `json: "outputEncoding"`
	Safe           string  `json: "safe"`
	Cx             string  `json: "cx"`
}

type googlePageMap struct {
	Cse_thumbnail []thumb    `json: "cse_thumbnail"`
	MetaTags      []metaTags `json: "metatags"`
	Cse_image     []image    `json: "cse_image"`
}

type thumb struct {
	Width  string `json: "width"`
	Height string `json: "height"`
	Src    string `json: "src"`
}

type metaTags struct {
	Referrer string `json: "referrer"`
}

type image struct {
	Src string `json: "src"`
}
