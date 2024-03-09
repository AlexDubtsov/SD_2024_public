package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/webhandler"
)

func main() {

	fileserver := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fileserver)
	http.HandleFunc("/basicList", webhandler.BasicListEventsHandler)
	http.HandleFunc("/basicEventEdit", webhandler.BasicEditEventHandler)
	http.HandleFunc("/basicMembersEdit", webhandler.BasicEditMembersHandler)
	http.HandleFunc("/basicEventCreate", webhandler.BasicCreateEventHandler)
	fmt.Println("Server is running at: http://localhost:8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		log.Fatal(err)
		fmt.Printf("Error starting server")
		os.Exit(1)
	}
}
