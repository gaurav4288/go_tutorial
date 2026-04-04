package main

func fetchTasks(baseURL, availability string) []Issue {
	// ?
	switch availability {
	case "Low":
		baseURL += "?limit=1"
	case "Medium":
		baseURL += "?limit=3"
	case "High":
		baseURL += "?limit=5"
	}

	baseURL += "&sort=estimate"

	fullURL := baseURL
	return getIssues(fullURL)
}
