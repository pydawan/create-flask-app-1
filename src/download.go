package src

import (
	"log"

	git "gopkg.in/src-d/go-git.v4"
)

var repoURL = "https://github.com/spadevs/flask_minimal"

// CloneRepo function clone repository from github
func CloneRepo(currentPath string) {
	_, err := git.PlainClone(currentPath, false, &git.CloneOptions{
		URL: repoURL,
	})

	if err != nil {
		log.Fatalln(err)
	}
}
