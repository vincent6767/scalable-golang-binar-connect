package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("https://accelbyte.net")
	checkErr(err)
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(len(body))
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}