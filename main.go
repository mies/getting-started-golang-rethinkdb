package main

import (
	"log"
	"net/http"
	"encoding/json"
	r "github.com/christopherhesse/rethinkgo"
)

type Bookmark struct {
	Title string
	Url	  string
	Id	  string `json:"id,omitempty"`
}


var (
	cities  = []string{
		"Amsterdam", "San Francisco", "Paris", "New York", "Portland",
	}
)

func main() {
	flag.Parse()
	endpoint := fmt.Sprintf("%v:%v", *Address, *Port)

	http.HandleFunc("/", handleIndex)

	fmt.Printf("Hosting at %v\n", endpoint)
	if err := http.ListenAndServe(endpoint, nil); err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func handleIndex(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(rw)
	encoder.Encode(cities)
}
