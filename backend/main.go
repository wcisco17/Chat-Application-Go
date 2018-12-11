package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("../static", t.filename)))
	})
	t.templ.Execute(w, nil)
}

// Calls the chat app html
func main() {
	r := newRoom()

	// Main Template
	http.Handle("/", &templateHandler{filename: "chat.html"})
	// Chat Room
	http.Handle("/room", r)

	// Start Chat Room
	go r.run()

	// Start the Web Server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
