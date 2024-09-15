package models

import "time"

type Member struct {
	GithubUsername string    `json:"githubUsername"`
	Name           string    `json:"name"`
	DateJoined     time.Time `json:"dateJoined"`
	Website        string    `json:"website"`
}
