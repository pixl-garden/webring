package utils

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pixl-garden/webring/internal/models"
)

func UpdateReadme(members []models.Member) error {
	readmeContent, err := ioutil.ReadFile("README.md")
	if err != nil {
		return err
	}

	content := string(readmeContent)
	startMarker := "<!-- MEMBERS_START -->"
	endMarker := "<!-- MEMBERS_END -->"

	startIndex := strings.Index(content, startMarker)
	endIndex := strings.Index(content, endMarker)

	if startIndex == -1 || endIndex == -1 {
		return fmt.Errorf("couldn't find markers in README")
	}

	var membersList strings.Builder
	membersList.WriteString(startMarker + "\n")
	for _, member := range members {
		membersList.WriteString(fmt.Sprintf("- [%s](%s) - [@%s](https://github.com/%s)\n", 
			member.Name, member.Website, member.GithubUsername, member.GithubUsername))
	}
	membersList.WriteString(endMarker)

	newContent := content[:startIndex] + membersList.String() + content[endIndex+len(endMarker):]

	return ioutil.WriteFile("README.md", []byte(newContent), 0644)
}