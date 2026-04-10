package main

import (
	"io/fs"
	"net/http"

	"github.com/alwindoss/hon/ui"
)

func main() {

	// Strip the "dist" prefix from the embedded paths
	publicFS, _ := fs.Sub(ui.FS, "dist")

	// Serve the static files
	// http.Handle("/", http.FileServer(http.FS(serverRoot)))
	http.Handle("/", http.FileServer(http.FS(publicFS)))

	http.ListenAndServe(":8080", nil)
}
