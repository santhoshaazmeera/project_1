package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connectiondetails := "user=postgres dbname=mydb_2 sslmode=disable password = test@123"
	db, err := sql.Open("postgres", connectiondetails)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to the database !")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("succesfully connected")

	rows:= "insert into employeetable (name,phonenumber,email) values ("san",)"

}


