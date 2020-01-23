package main

import (
	"create-flask-app/src"
	"os"
	"path"
)

func main() {
	folderName := os.Args[1]
	currentDir := src.GetCurrentDir()
	projectPath := path.Join(currentDir, folderName)

	src.CloneRepo(projectPath)

}
