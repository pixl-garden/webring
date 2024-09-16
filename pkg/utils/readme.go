package utils

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/pixl-garden/webring/pkg/models"
)

func UpdateReadme(members []models.Member) error {
	sort.Slice(members, func(i, j int) bool {
		return members[i].DateJoined.Before(members[j].DateJoined)
	})

	var content strings.Builder
	content.WriteString("# Pixl Garden Webring\n\n")
	content.WriteString("Welcome to the Pixl Garden Webring! Here are our current members:\n\n")

	for _, member := range members {
		content.WriteString(fmt.Sprintf("- [%s](%s) (Joined: %s)\n", 
			member.Name, 
			member.Website, 
			member.DateJoined.Format("2006-01-02")))
	}

	content.WriteString("\n## How to Join\n\n")
	content.WriteString("To join the Pixl Garden Webring, please submit a pull request with your information.\n")

	err := os.WriteFile("README.md", []byte(content.String()), 0644)
	if err != nil {
		return fmt.Errorf("failed to write README.md: %w", err)
	}

	return nil
}