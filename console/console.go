package console

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var NextDBload time.Time
var NextDBsave time.Time

func dbFolderPath() (string, error) {
	// Name of folder with DB
	var dbFolderName = "SD_2024_private"

	// Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory")
		log.Fatal(err)
		return "", err
	}

	// Get the parent directory
	parentDir := filepath.Dir(currentDir)

	// Create the full path to the other folder
	return filepath.Join(parentDir, dbFolderName), nil

}

func ConsoleSave() {

	dbPath, err := dbFolderPath()

	if err != nil {
		fmt.Println("Error on getting DB path")
		log.Fatal(err)
		return
	}

	// Create and configure the first command
	cmd1 := exec.Command("git", "add", ".")
	cmd1.Dir = dbPath

	// Execute the first command
	output1, err := cmd1.Output()
	if err != nil {
		fmt.Println("Error executing the GIT ADD .")
		log.Fatal(err)
		return
	}

	// Output the result of the first command
	fmt.Println(string(output1))

	// Create and configure the second command
	cmd2 := exec.Command("git", "commit", "-m", "\"Save DB\"")
	cmd2.Dir = dbPath

	// Execute the second command
	output2, err := cmd2.Output()
	if err != nil {
		fmt.Println("Executing the GIT COMMIT")
		fmt.Println("No changes in Data base")
		return
	}

	// Output the result of the second command
	fmt.Println(string(output2))

	// Create and configure the third command
	cmd3 := exec.Command("git", "push")
	cmd3.Dir = dbPath

	// Execute the third command
	output3, err := cmd3.Output()
	if err != nil {
		fmt.Println("Error executing the GIT PUSH")
		log.Fatal(err)
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

func ConsoleLoad() {

	dbPath, err := dbFolderPath()

	if err != nil {
		fmt.Println("Error on getting DB path")
		log.Fatal(err)
		return
	}

	// Create and configure the first command
	cmd1 := exec.Command("git", "pull")
	cmd1.Dir = dbPath

	// Execute the first command
	output1, err := cmd1.Output()
	if err != nil {
		fmt.Println("Error executing the GIT PULL")
		log.Fatal(err)
		return
	}

	// Output the result of the first command
	fmt.Println(string(output1))

	// Get the current time
	currentTime := time.Now()

	// Format the current time in an easy-to-read format
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	fmt.Println("DB loaded on", formattedTime)

}
