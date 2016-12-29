package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

func init() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/pong", pongHandler)}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "home.html")
	serveTemplate(w, fp)
}

func pongHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, getIPAddress(r))
}

func getIPAddress(r *http.Request) string {
	if strings.HasPrefix(r.RemoteAddr, "[") {
		return strings.Split(strings.Replace(r.RemoteAddr, "[", "", 1), "]:")[0]
	}
	return strings.Split(r.RemoteAddr, ":")[0]
}

func serveTemplate(w http.ResponseWriter, fp string) {
	lp := path.Join("templates", "layout.html")

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Fatal(err.Error(), err)
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Fatal(err.Error(), err)
		http.Error(w, http.StatusText(500), 500)
	}
}
