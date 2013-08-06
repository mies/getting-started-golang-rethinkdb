package main

import (
	"log"
	"os"
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

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/new", createBookmark)

	session, err := r.Connect(os.Getenv("WERCKER_RETHINKDB_URL"), "gettingstarted")
	if err != nil {
		log.Fatal(err)
		return
	}

	err = r.DbCreate("gettingstarted").Run(session).Exec()
	if err != nil {
	  log.Println(err)
    }

	err = r.TableCreate("bookmarks").Run(session).Exec()
    if err != nil {
	  log.Println(err)
    }

	err = http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("Error: %v", err)
	}
}

func createBookmark(res http.ResponseWriter, req *http.Request) {
	session, err := r.Connect(os.Getenv("WERCKER_RETHINKDB_URL"), "gettingstarted")
	if err != nil {
		log.Fatal(err)
		return
	}

	b := new(Bookmark)
	json.NewDecoder(req.Body).Decode(b)

	var response r.WriteResponse

	err = r.Table("bookmarks").Insert(b).Run(session).One(&response)
	if err != nil {
		log.Fatal(err)
		return
	}
	data, _ := json.Marshal("{'bookmark':'saved'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	session, err := r.Connect(os.Getenv("WERCKER_RETHINKDB_URL"), "gettingstarted")
	if err != nil {
	  log.Fatal(err)
	  return
    }

	var response []Bookmark

	err = r.Table("bookmarks").Run(session).All(&response)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := json.Marshal(response)

	res.Header().Set("Content-Type", "application/json")
	res.Write(data)
}
