package gSearch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendGoogleSearchRequest(termToSearch string) ([]string, error) {

	SearchEngineID := "PASTE-YOUR-CUSTOM-SEARCH-ID"
	APIgSearchKey := "PASTE-YOUR-GOOGLECUSTOM-SEARCH-API"

	v := url.Values{}

	v.Set("q", termToSearch)
	v.Set("cx", SearchEngineID)
	v.Set("key", APIgSearchKey)

	u := url.URL{
		Scheme:   "https",
		Host:     "www.googleapis.com",
		Path:     "customsearch/v1",
		RawQuery: v.Encode(),
	}

	// u, _ := url.Parse("http://google.com")
	// // u := url.URL{}
	// u.Scheme = "https"
	// u.Host = "www.googleapis.com"

	// v := url.Values{}
	// v.Set("q", termToSearch)
	// v.Set("cx", SearchEngineID)
	// v.Set("key", APIgSearchKey)

	// u.RawQuery = v.Encode()

	fmt.Println(u)

	//	making httpRequest
	var resp *http.Response
	var err error

	if resp, err = http.Get(u.String()); err != nil {
		fmt.Printf("FATAL: Could not parse request: %v\n", err)
		return nil, errors.New("NOT FOUND")
	}

	fmt.Println(resp.Status)
	defer resp.Body.Close()

	var output []byte
	if output, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println("FATAL: Could not parse request: %v\n", err)
		return nil, errors.New("NOT FOUND")

	}

	//fmt.Println(string(output))
	var m googleSearchResponse
	if err = json.Unmarshal(output, &m); err != nil {
		fmt.Println("error on unmarshalling", err)
		return nil, errors.New("NOT FOUND")
	}

	if m.SearchInformation.TotalResults == "0" {
		return nil, errors.New("NOT FOUND")
	}

	n := len(m.Items)

	if n <= 0 {
		return nil, errors.New("NOT FOUND")
	} else {

		// if n > 20{
		// 	n = 20
		// }

		response := make([]string, n)

		// var contents [n]string
		for i := 0; i < n; i++ {
			content := m.Items[i].Snippet
			response[i] = content + "\n" + m.Items[i].FormattedUrl + "\n\n"
		}

		return response, nil

		//response = html.UnescapeString((contents[0] + "\n" + m.ResponseData.Results[0].Url + "\n\n" + contents[1] + "\n" + m.ResponseData.Results[1].Url + "\n\n" + contents[2] + "\n" + m.ResponseData.Results[2].Url))
	}
}
