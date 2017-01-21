package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

//  {"response":[{"id":728530,"first_name":"Elena","last_name":"Grahovac","hidden":1}]}

type UserResponse struct {
	Response []User `json:"response"`
	Error *ResponseError `json:"error,omitempty"`
}

type ResponseError struct {
	Code int `json:"error_code"`
	Message string `json:"error_msg"`
}

type User struct {
	ID int32 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}


func main() {
	// ac278e4a7b0c56939e
	//url := "https://api.vk.com/method/account.getInfo?v=5.62"
	url := "https://api.vk.com/method/users.get?user_ids=grahovac&v=5.62"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Couldn't get: %v\n", err)
	}

	/*b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf("Couldn't receive body: %v\n", err)
	}

	fmt.Printf("The body is: %s\n", b)*/

	users := new(UserResponse)
	err = json.NewDecoder(resp.Body).Decode(users)

	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	if users.Error != nil {
		// handle error

		users.Error.Code = 5
	}

	fmt.Printf("%v", users.Response)
}
