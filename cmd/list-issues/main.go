package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/xanzy/go-gitlab"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gitlabToken := os.Getenv("GITLAB_TOKEN")
	git, err := gitlab.NewClient(gitlabToken)
	if err != nil {
		log.Fatal(err)
	}

	bla := true
	// Create new project
	p := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    2,
			PerPage: 200,
		},

		Membership: &bla,
	}

	projects, _, err := git.Projects.ListProjects(p)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("length project: %v", len(projects))

	for _, p := range projects {
		log.Printf("Found project: %s", p.Name)
	}
}
