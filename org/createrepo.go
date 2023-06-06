package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v52/github"
)

// Repo - Represents a github repository.
type Repo struct {
	Name           string
	Description    string
	DefaultBranch  string
	AutoInitialize bool
	Private        bool
	GitIgnore      string
	CommitMessage  string
}

func main() {
	t := "Your user access token goes here"
	ctx := context.Background()
	c := github.NewTokenClient(ctx, t)

	r := Repo{
		Name:           "pantheon-repo" + fmt.Sprint(time.Now().Unix()),
		Description:    "Pantheon generated repo",
		DefaultBranch:  "main",
		AutoInitialize: true,
		Private:        false,
	}

	repo, _, err := c.Repositories.Create(ctx, "sachin-test-org-01", &github.Repository{
		Name:         &r.Name,
		Description:  &r.Description,
		MasterBranch: &r.DefaultBranch,
		AutoInit:     &r.AutoInitialize,
		Private:      &r.Private,
	})

	if err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)

	fmt.Printf("Repo ID: %v\n", *repo.ID)
	fmt.Printf("Repo NodeID: %s\n", *repo.NodeID)
	fmt.Printf("Repo Name: %s\n", *repo.Name)
	fmt.Printf("Repo Description: %s\n", *repo.Description)
	fmt.Printf("Repo URL: %s\n", *repo.HTMLURL)

}
