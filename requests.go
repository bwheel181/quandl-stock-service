package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func MakeLatestDataRequest(ticker string, ch chan<- string) {
	reqStr := fmt.Sprintf("%s%s.json?rows=1&api_key=%s", BASE_URL, ticker, API_KEY)
	res, err := http.Get(reqStr)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer res.Body.Close()
		contents, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		ch <- string(contents)
	}
}
