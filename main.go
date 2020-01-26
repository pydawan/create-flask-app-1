package main

import (
	"flag"
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
	var name string
	var template bool

	flag.StringVar(&name, "name", "", "set the project name")
	flag.BoolVar(&template, "mvc", false, "use for download mvc template")

	flag.Parse()

	currentPath := path.Join(getCurrentDir(), name)

	src.Download(setRepo(template), getCurrentDir(), name)
	src.CreatePythonVenv(currentPath, name)
}

func setRepo(template bool) string {
	if template {
		return "https://storage.googleapis.com/spacedevs_storage_intervalo_1/repos/flask_template_mvc-master.zip"
	}
	return "https://storage.googleapis.com/spacedevs_storage_intervalo_1/repos/flask_minimal-master.zip"
}
