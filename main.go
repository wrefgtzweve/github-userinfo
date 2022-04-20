package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v43/github"
)

var username string

func printVar(prefix string, value any) {
	if value == "" || value == nil || value == false {
		return
	}
	fmt.Printf("%v: %v\n", prefix, value)
}

func main() {
	ctx := context.Background()

	fmt.Println("Github username:")
	fmt.Scan(&username)

	client := github.NewClient(nil)
	userGet, _, err := client.Users.Get(ctx, username)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("\n")
	printVar("Username", userGet.GetLogin())
	printVar("Name", userGet.GetName())
	printVar("Site admin", userGet.GetSiteAdmin())
	printVar("Bio", userGet.GetBio())
	printVar("Location", userGet.GetLocation())
	printVar("Company", userGet.GetCompany())
	printVar("Blog", userGet.GetBlog())
	printVar("Twitter", userGet.GetTwitterUsername())
	printVar("Email", userGet.GetEmail())
	printVar("Public gists", userGet.GetPublicGists())
	printVar("Created at", userGet.GetCreatedAt())
	printVar("Public repositories", userGet.GetPublicRepos())
	printVar("Profile URL", userGet.GetHTMLURL())

	events, _, err := client.Activity.ListEventsPerformedByUser(ctx, username, true, &github.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nRecent commits (includes coauthored commits): ")

	commitArray := map[string]bool{}
	for _, event := range events {
		if event.GetType() == "PushEvent" {
			eventInterface, _ := event.ParsePayload()
			payload := eventInterface.(*github.PushEvent)
			tempString := payload.Commits[0].Author.GetName() + " | " + payload.Commits[0].Author.GetEmail()
			commitArray[tempString] = true
		}
	}
	for commit := range commitArray {
		fmt.Println(commit)
	}

	var empty string

	// Shitty hack to prevent the .exe from instantly closing
	fmt.Print("\n")
	fmt.Println("Press any key to exit")
	fmt.Scanln(&empty)
	fmt.Scanln(&empty)
}
