package handlers

import (
	"fmt"
	"net/http"
	"github.com/Sirupsen/logrus"
)

// Hello echoes the Path component of the requested URL.
func Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	fmt.Fprintf(w, "Code is %s", code)
}
