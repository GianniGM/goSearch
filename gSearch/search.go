package gSearch

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func SendGoogleSearchRequest(termToSearch string) []string {
	//TODO: must formatting termToSearch to url mode
	url := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&q=" + url.QueryEscape(termToSearch)
	fmt.Println(url)

	var response []string
	response = nil

	//	making httpRequest
	if resp, err := http.Get(url); err != nil {
		fmt.Printf("FATAL: Could not parse request: %v\n", err)
	} else {
		var m googleSearchResponse
		fmt.Println(resp.Status)

		defer resp.Body.Close()
		output, _ := ioutil.ReadAll(resp.Body)

		//fmt.Println(string(output))
		if err = json.Unmarshal(output, &m); err != nil {
			fmt.Println("error on unmarshalling", err)
		} else {
			if len(m.ResponseData.Results) > 0 {
				n := 20

				if len(m.ResponseData.Results) < n {
					n = len(m.ResponseData.Results)
				}

				response = make([]string, n)

				// var contents [n]string
				for i := 0; i < n; i++ {
					temp := strings.Replace(m.ResponseData.Results[i].Content, "<b>", "", -1)
					content := strings.Replace(temp, "</b>", "", -1)
					response[i] = html.UnescapeString(content + "\n" + m.ResponseData.Results[i].Url + "\n\n")
				}

				//response = html.UnescapeString((contents[0] + "\n" + m.ResponseData.Results[0].Url + "\n\n" + contents[1] + "\n" + m.ResponseData.Results[1].Url + "\n\n" + contents[2] + "\n" + m.ResponseData.Results[2].Url))
			}
		}
	}

	return response
}
