package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func SendGoogleSearchRequest(termToSearch string) string {
	//TODO: must formatting termToSearch to url mode
	response := "sorry! nothing found"

	url := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&q=" + url.QueryEscape(termToSearch)
	fmt.Println(url)

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
			var contents [3]string
			for i := 0; i < 3; i++ {
				temp := strings.Replace(m.ResponseData.Results[i].Content, "<b>", "", -1)
				contents[i] = strings.Replace(temp, "</b>", "", -1)
			}

			response = (contents[0] + "\n" + m.ResponseData.Results[0].Url + "\n\n" + contents[1] + "\n" + m.ResponseData.Results[1].Url + "\n\n" + contents[2] + "\n" + m.ResponseData.Results[2].Url)
		}
	}
	return response
}
