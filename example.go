package main

import (
	"fmt"
	r "github.com/christopherhesse/rethinkgo"
)

type Employee struct {
	FirstName string
	LastName  string
	Job       string
	Id        string `json:"id,omitempty"` // (will appear in json as "id", and not be sent if empty)
}

func addEmployee() {
	// To access a RethinkDB database, you connect to it with the Connect function
	session, err := r.Connect("localhost:28015", "company_info")
	if err != nil {
		fmt.Println("error connecting:", err)
		return
	}

	var response []Employee
	// Using .All(), we can read the entire response into a slice, without iteration
	err = r.Table("employees").Run(session).All(&response)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("response:", response)
	}

	// If we want to iterate over each result individually, we can use the rows
	// object as an iterator
	rows := r.Table("employees").Run(session)
	for rows.Next() {
		var row Employee
		if err = rows.Scan(&row); err != nil {
			fmt.Println("err:", err)
			break
		}
		fmt.Println("row:", row)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("err:", err)
	}
}
