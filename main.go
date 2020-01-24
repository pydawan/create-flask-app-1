package main

import (
	"log"
	"os"
	"path"

	"github.com/create-flask-app/src"
)

func getCurrentDir() string {
	currentDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return currentDir
}

func main() {
	projectName := os.Args[1]

	projectPath := path.Join(getCurrentDir(), projectName)

	src.CloneRepo(projectPath)
	src.CreatePythonVenv(projectPath, projectName)
}
