package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//resp, err := http.Get("http://localhost:1122")
	resp, err := http.Get("http://docker-network-server:1122")

	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(body))
		}
	}

	fmt.Println("==1==")
	c := make(chan string)
	<-c
}
