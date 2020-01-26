package src

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"

	"github.com/MarcusMann/create-flask-app/notifications"
	"github.com/cheggaaa/pb"
)

// CreatePythonVenv create a venv for python using python -m venv
func CreatePythonVenv(currentPath, projectName string) {
	cmd := exec.Command("python3", "-m", "venv", "venv")

	if runtime.GOOS == "windows" {
		cmd = exec.Command("python", "-m", "venv", "venv")
	}

	cmd.Dir = currentPath

	if _, err := cmd.Output(); err != nil {
		log.Fatal(err)
	} else {
		installPythonDeps(currentPath, projectName)
	}
}

func installPythonDeps(currentPath, projectName string) {
	pipCmd := currentPath + "/venv/bin/pip"

	if runtime.GOOS == "windows" {
		pipCmd = currentPath + "/venv/Scripts/pip"
	}

	c := exec.Command(pipCmd, "install", "-r", "requirements.txt")
	c.Dir = currentPath

	if out, err := c.Output(); err != nil {
		log.Fatalln(err)
	} else {
		notifications.Welcome()

		fmt.Println("\n\nInstalling Dependencies")

		count := len(out)

		bar := pb.New64(int64(count)).SetUnits(pb.U_BYTES).Start()
		bar.ShowCounters = false
		bar.ShowTimeLeft = false

		for i := 0; i < count; i++ {
			bar.Increment()
			time.Sleep(time.Millisecond)
		}

		bar.Finish()
	}

	notifications.GetInstalledMessage(projectName)
}
