package filefunctions

import (
	"fmt"
	"net/http"
)

func DownLoadFile(w http.ResponseWriter, r *http.Request, content, filename string) {
	contentType := "text/plain"
	filename = `attachment; filename="` + filename + `"`

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))
	w.Header().Set("Content-Disposition", filename)

	_, err := fmt.Fprint(w, content)
	if err != nil {
		http.Error(w, "Error on printing text content", http.StatusInternalServerError)
		return
	}
}
