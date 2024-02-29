package main

import (
	"fmt"
	"net/http"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/functions"
	"github.com/AlexDubtsov/SD_2024_public/m/v2/helpers"
)

func main() {
	// Create DB (in case of absence)
	if helpers.DB == nil {
		helpers.DBcreate()
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", functions.Homepage)
	fmt.Println("Server is running at: http://localhost:8080\n ")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
