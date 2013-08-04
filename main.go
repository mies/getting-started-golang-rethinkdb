package main

import (
	"log"
	"net/http"
	"encoding/json"
	"flag"
	r "github.com/christopherhesse/rethinkgo"
)

type Bookmark struct {
	Title string
	Url	  string
	Id	  string `json:"id,omitempty"`
}

func main() {
}



func main() {

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/new", createBookmark)

	if err := http.ListenAndServe(, nil); err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func handleIndex(rw http.ResponseWriter, req *http.Request) {
  session, err := r.Connect("localhost:28015", "bookmarks")
  if err != nil {
    log.Fatal("no connection")
	rw.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(rw)
	encoder.Encode(cities)
}
