package main

import (
	"github.com/wwg-ru/go-school/ws-04/client"
	"log"
)

func main() {
	parkings, err := velobike.GetParkings()

	if err != nil {
		log.Fatalf("Error is %s", err)
	}

	//parkings.Items[0].Address
}
