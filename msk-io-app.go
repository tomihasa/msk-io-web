package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/pong", pongHandler)
	http.Handle("/", loggingHandler(handlers.CompressHandler(r)))

	assets := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	http.ListenAndServe(":3000", nil)
}

func loggingHandler(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
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

// func serveNotFound(w http.ResponseWriter, r *http.Request) {
// 	fp := path.Join("templates", "errors", "notfound.html")
// 	w.WriteHeader(http.StatusNotFound)
// 	serveTemplate(w, fp)
// }

// func templateHandler(w http.ResponseWriter, r *http.Request) {
//
// 	fp := path.Join("templates", r.URL.Path)
//
// 	// Return a 404 if the template doesn't exist
// 	info, err := os.Stat(fp)
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			serveNotFound(w, r)
// 			return
// 		}
// 	}
//
// 	// Return a 404 if the request is for a directory
// 	if info.IsDir() {
// 		serveNotFound(w, r)
// 		return
// 	}
//
// 	serveTemplate(w, fp)
//
// }

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
