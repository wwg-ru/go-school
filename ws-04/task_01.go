package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://apivelobike.velobike.ru/ride/parkings"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Couldn't get: %v\n", err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf("Couldn't receive body: %v\n", err)
	}

	fmt.Printf("The body is: %s\n", b)
}
