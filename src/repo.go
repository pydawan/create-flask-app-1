package src

import (
	"log"
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

var repoURL = "https://github.com/spadevs/flask_minimal"

// GetCurrentDir return current directory
func GetCurrentDir() string {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return dir
}

// CloneRepo function clone repository from github
func CloneRepo(currentPath string) {
	_, err := git.PlainClone(currentPath, false, &git.CloneOptions{
		URL: repoURL,
	})

	if err != nil {
		log.Fatalln(err)
	}
}
