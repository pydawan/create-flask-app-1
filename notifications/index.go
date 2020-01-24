package notifications

import "fmt"

// GetInstalledMessage Show message when an app is installed
func GetInstalledMessage(name string) {
	fmt.Print(` ____ ____  _____    _  _____ _____   _____ _        _    ____  _  __     _    ____  ____  
 / ___|  _ \| ____|  / \|_   _| ____| |  ___| |      / \  / ___|| |/ /    / \  |  _ \|  _ \ 
| |   | |_) |  _|   / _ \ | | |  _|   | |_  | |     / _ \ \___ \| ' /    / _ \ | |_) | |_) |
| |___|  _ <| |___ / ___ \| | | |___  |  _| | |___ / ___ \ ___) | . \   / ___ \|  __/|  __/ 
 \____|_| \_\_____/_/   \_\_| |_____| |_|   |_____/_/   \_\____/|_|\_\ /_/   \_\_|   |_|`)

	fmt.Println("\n\n\n👉 Successfully created project")
	fmt.Println("👉 Get started with the following commands")
	fmt.Println()
	fmt.Printf("\033[1;36m%s\033[0m \n", "$ cd "+name)
	fmt.Printf("\033[1;36m%s\033[0m \n", "$ source venv/bin/activate")
	fmt.Printf("\033[1;36m%s\033[0m \n\n", "$ flask run")
}
