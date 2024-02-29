package helpers

import (
	"fmt"
	"net/http"
)

func Handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 not found")
}
