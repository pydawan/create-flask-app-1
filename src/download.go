package src

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

// Download create request to flask template
func Download(templateURL, currentPath, projectName string) {
	zipFileName := "template.zip"

	response, err := http.Get(templateURL)

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	defer response.Body.Close()

	// create a zip file
	createZipFile(zipFileName, response.Body)

	// extract a zip file
	extractRepo(zipFileName, currentPath, projectName)
}

func createZipFile(fileName string, body io.ReadCloser) {
	output, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	defer output.Close()

	_, err2 := io.Copy(output, body)

	if err2 != nil {
		log.Fatalf("error: %s", err2)
	}
}

func extractRepo(zipFileName string, currentPath, projectName string) {
	reader, err := zip.OpenReader(zipFileName)

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	defer reader.Close()

	for _, f := range reader.Reader.File {
		// search for dir and change dir name
		r := regexp.MustCompile(`.*-master\/`)
		projectNameCompiled := r.ReplaceAllLiteralString(f.Name, projectName+"/")

		zipped, err := f.Open()

		if err != nil {
			log.Fatalf("error: %s", err)
		}

		defer zipped.Close()

		path := filepath.Join(currentPath, projectNameCompiled)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, f.Mode())

			if err != nil {
				log.Fatalf("error: %s", err)
			}

			defer writer.Close()

			if _, err = io.Copy(writer, zipped); err != nil {
				log.Fatal(err)
			}
		}
	}
}
