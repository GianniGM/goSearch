package gSearch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendGoogleSearchRequest(termToSearch string) *[]string {

	SearchEngineID := "PASTE-YOUR-CUSTOM-SEARCH-ID"
	APIgSearchKey := "PASTE-YOUR-GOOGLECUSTOM-SEARCH-API"

	url := "https://www.googleapis.com/customsearch/v1?q=" + url.QueryEscape(termToSearch) + "&cx=" + SearchEngineID + "&key=" + APIgSearchKey

	//TODO: must formatting termToSearch to url mode
	// url := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&q=" + url.QueryEscape(termToSearch)

	fmt.Println(url)

	//	making httpRequest
	var resp *http.Response
	var err error

	if resp, err = http.Get(url); err != nil {
		fmt.Printf("FATAL: Could not parse request: %v\n", err)
		return nil
	}

	fmt.Println(resp.Status)
	defer resp.Body.Close()

	var output []byte
	if output, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println("FATAL: Could not parse request: %v\n", err)
		return nil
	}

	//fmt.Println(string(output))
	var m googleSearchResponse
	if err = json.Unmarshal(output, &m); err != nil {
		fmt.Println("error on unmarshalling", err)
		return nil
	}

	n := len(m.Items)

	if n <= 0 {
		return nil
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

		return &response

		//response = html.UnescapeString((contents[0] + "\n" + m.ResponseData.Results[0].Url + "\n\n" + contents[1] + "\n" + m.ResponseData.Results[1].Url + "\n\n" + contents[2] + "\n" + m.ResponseData.Results[2].Url))
	}
}
