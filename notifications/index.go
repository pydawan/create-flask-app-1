package notifications

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

// GetInstalledMessage Show message when an app is installed
func GetInstalledMessage(name string) {
	fmt.Println("\n\n\n👉 Successfully created project")
	fmt.Println("👉 Get started with the following commands")
	fmt.Println()
	fmt.Printf("\033[1;36m%s\033[0m \n", "$ cd "+name)
	fmt.Printf("\033[1;36m%s\033[0m \n", "$ source venv/bin/activate")
	fmt.Printf("\033[1;36m%s\033[0m \n\n", "$ flask run")
}

// Welcome show the app name
func Welcome() {
	myFigure := figure.NewFigure("Create Flask App", "", true)
	myFigure.Print()
}
