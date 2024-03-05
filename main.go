package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/functions"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/helpers"
)

func main() {
	// Create DB (in case of absence)
	if helpers.DB == nil {
		helpers.DBcreate()
	}

	fileserver := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fileserver)
	http.HandleFunc("/basic", functions.BasicProcessListPage)
	http.HandleFunc("/basicEvent", functions.BasicProcessEventPage)
	fmt.Println("Server is running at: http://localhost:8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
