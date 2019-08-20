package github

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestSearchIssues(t *testing.T) {
	result, err := SearchIssues([]string{"repo:golang/go"})

	if err != nil {
		log.Fatal(err)
	}
	output(10 * time.Hour, result)
}

func output(within time.Duration, result *IssuesSearchResult) {
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if time.Now().Sub(item.CreatedAt).Seconds() < within.Seconds() {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
