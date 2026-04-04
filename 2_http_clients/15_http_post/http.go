package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func createUser(url, apiKey string, data User) (User, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return User{}, err
	}

	defer res.Body.Close()

	var user User
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&user); err != nil {
		return User{}, err
	}
	fmt.Println("Status Code: ",res.Status)
	return user, nil
}
