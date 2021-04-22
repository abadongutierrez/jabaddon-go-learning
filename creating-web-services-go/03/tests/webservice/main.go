package main

import (
	"net/http"
)

// like defining a type with properties/attributes
type fooHandler struct {
	Message string
}

// wtf is happening here?
// this is like attaching method ServeHTTP to type fooHandler
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar called"))
}

func main() {
	http.Handle("/foo", &fooHandler{Message: "Hello World!"})
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":5000", nil)
}
