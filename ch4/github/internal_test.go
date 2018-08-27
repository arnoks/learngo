package github

import (
	"fmt"
	"log"
	"testing"
)

func TestSeach(t *testing.T) {
	itms := []string{"repo:arnoks/learngo", "is:open"}
	result, err := SearchIssues(itms)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number,
			item.User.Login,
			item.Title)
	}
}
