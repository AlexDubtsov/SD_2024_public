package console

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func ConsoleSave() {

	// Relative path to another folder
	relativePath := "SD_2024_private"

	// Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Get the parent directory
	parentDir := filepath.Dir(currentDir)

	// Create the full path to the other folder
	newPath := filepath.Join(parentDir, relativePath)

	// Create and configure the first command
	cmd1 := exec.Command("git", "add", ".")
	cmd1.Dir = newPath

	// Execute the first command
	output1, err := cmd1.Output()
	if err != nil {
		fmt.Println("Error executing the GIT ADD .", err)
		return
	}

	// Output the result of the first command
	fmt.Println(string(output1))

	// Create and configure the second command
	cmd2 := exec.Command("git", "commit", "-m", "\"Save with button\"")
	cmd2.Dir = newPath

	// Execute the second command
	output2, err := cmd2.Output()
	if err != nil {
		fmt.Println("Executing the GIT COMMIT", err)
		fmt.Println("No chandes in Data base")
		return
	}

	// Output the result of the second command
	fmt.Println(string(output2))

	// Create and configure the third command
	cmd3 := exec.Command("git", "push")
	cmd3.Dir = newPath

	// Execute the third command
	output3, err := cmd3.Output()
	if err != nil {
		fmt.Println("Error executing the GIT PUSH", err)
		return
	}

	// Output the result of the third command
	fmt.Println(string(output3))

	// Get the current time
	currentTime := time.Now()

	// Format the current time in an easy-to-read format
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	fmt.Println("Saved on", formattedTime)
}
