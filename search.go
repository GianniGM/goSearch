package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	SendRequest("ciao mondo")
}

func SendRequest(termToSearch string) {
	//TODO: must formatting termToSearch to url mode

	url := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&q=" + url.QueryEscape(termToSearch)
	fmt.Println(url)

	//	making httpRequest
	if resp, err := http.Get(url); err != nil {
		fmt.Printf("FATAL: Could not parse request: %v\n", err)
		os.Exit(1)
	} else {
		var m googleSearchResponse
		fmt.Println(resp.Status)
		defer resp.Body.Close()
		output, _ := ioutil.ReadAll(resp.Body)

		if err = json.Unmarshal(output, &m); err != nil {
			fmt.Println(m.Details)
		} else {
			fmt.Println("error on unmarshalling")
		}
	}
}
