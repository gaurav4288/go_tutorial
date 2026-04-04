package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func getUsers(url string) ([]User, error) {
	// resp, err := http.Get(url)
	// if err != nil {
	// 	fmt.Println("Error in fetching user", err)
	// }

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	fmt.Println("Requesting users from API...", req)
	
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var user []User
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&user); err != nil {
	    fmt.Println("Error decoding response:", err)
	}
	return user, nil
}
