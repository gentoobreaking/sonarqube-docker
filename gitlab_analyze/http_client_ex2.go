package main

import (
	//      "bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url" //httpPost_values
	"strings" //httpPost
)

func httpGet() {
	resp, err := http.Get("https://tw.yahoo.com/")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("https://tw.yahoo.com/",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=test"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPost_values() {
	resp, err := http.PostForm("https://tw.yahoo.com/",
		url.Values{"key": {"Value"}, "id": {"test"}})
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPost_complex() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://tw.yahoo.com/", strings.NewReader("name=test"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=test")

	resp, err := client.Do(req)

	defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                // handle error
        }

	//fmt.Println(client.Do(req))
	fmt.Println(string(body))
}

func main() {
	//	httpGet()
	//httpPost()
	//httpPost_values()
	httpPost_complex()
}
