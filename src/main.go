package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

func main() {
	resp, err := http.Get("http://api:8000")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}