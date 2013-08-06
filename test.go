package main

import (
  "log"
  r "github.com/christopherhesse/rethinkgo"
)

func simpleTest() {
  session, err := r.Connect("localhost:28015", "test")
  if err != nil {
    log.Fatal(err)
  }

  /*

  err = r.DbCreate("test").Run(session).Exec()
  if err != nil {
    log.Fatal(err)
  }
  */

  err = r.TableCreate("bookmarks").Run(session).Exec()
  if err != nil {
    log.Fatal(err)
  }
}
