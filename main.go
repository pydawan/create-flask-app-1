package main

import (
	"log"
	"os"
	"path"
)

func getCurrentDir() string {
	currentDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return currentDir
}
func main() {
	folderName := os.Args[1]
	currentDir := src.GetCurrentDir()
	projectPath := path.Join(currentDir, folderName)

	src.CloneRepo(projectPath)

}
