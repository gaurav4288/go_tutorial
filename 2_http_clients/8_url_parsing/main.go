package main

import (
	"net/url"
	"fmt"
)

func getDomainNameFromURL(rawURL string) (string, error) {
	// ?
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	fmt.Println(parsedUrl)
	var ulr string = string(parsedUrl.Hostname())

	return ulr, nil
}
