package functions

import (
	"net/http"
)

func EventsPage(w http.ResponseWriter, r *http.Request) {
	// Create template for homepage, which displays all posts and user status.

	// CombinedTemplate := helpers.EventTemplate{
	// 	AllEvents: AllUserPosts,
	// }

	// events_tmpl, err := template.ParseFiles("static/events.html")
	// if err != nil {
	// 	http.Error(w, "Error parsing template", http.StatusInternalServerError)
	// 	fmt.Printf("Error parsing template: %v\n", err)
	// 	return
	// }

	// err = events_tmpl.Execute(w, CombinedTemplate)
	// if err != nil {
	// 	http.Error(w, "Error executing template", http.StatusInternalServerError)
	// 	fmt.Println("Error executing template:", err)
	// }
}

func ParticipantsPage(w http.ResponseWriter, r *http.Request) {
	// Create template for homepage, which displays all posts and user status.

	// CombinedTemplate := helpers.Template{
	// 	UserStatus:   UserStatus,
	// 	AllUserPosts: AllUserPosts,
	// }

	// partic_tmpl, err := template.ParseFiles("static/participants.html")
	// if err != nil {
	// 	http.Error(w, "Error parsing template", http.StatusInternalServerError)
	// 	fmt.Printf("Error parsing template: %v\n", err)
	// 	return
	// }

	// err = partic_tmpl.Execute(w, CombinedTemplate)
	// if err != nil {
	// 	http.Error(w, "Error executing template", http.StatusInternalServerError)
	// 	fmt.Println("Error executing template:", err)
	// }
}
