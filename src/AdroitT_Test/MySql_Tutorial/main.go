package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json: "name"`
}

func main() {
	fmt.Println("Go SQL Tutorial")

	db, err := sql.Open("mysql", "root:Austin@36@tcp(127.0.0.1:3306)/Work")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("SELECT name FROM Employee")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var employee User

		err = results.Scan(&employee.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(employee.Name)
	}

	// inserting Query
	//insert, err := db.Query("INSERT INTO Employee VALUES('Austine')")
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer insert.Close()

}
