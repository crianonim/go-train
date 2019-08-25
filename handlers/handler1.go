package handlers

import (
	"fmt"
	"net/http"
)

// HandlerOne handles it.
func HandlerOne(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h3>Bye</h3>")
}
