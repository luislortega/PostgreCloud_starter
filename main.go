package main

import (
	"fmt"
	"database/sql"
	_"github.com/lib/pq"
)

const (
	host = "ec2-50-19-32-202.compute-1.amazonaws.com"
	port = 5432
	user = "pjlwbscxqfpprr"
	password = "6ed9f63fbe84327e5e6322c45112f97c477ffc03070f33f22b931d85e772000f"
	dbname = "d2e14vb5sir2of"
)

type Example struct {
	id int
	data string 
}

func main(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
							host, port, user, password, dbname)
	db,err := sql.Open("postgres", psqlInfo)
	
	if err != nil {
		panic(err)
	}
	
	_,err = db.Exec("CREATE TABLE IF NOT EXISTS example(id integer, data varchar(32))")
	
	if err != nil {
		panic(err)
	}

	insertData, err := db.Prepare("INSERT INTO example (id, data) values (1, 'data test')")
	
	if err != nil{
		panic(err)
	}

	insertData.Exec()

	var myExample Example
	query := "SELECT id, data FROM example WHERE id=1"
	err = db.QueryRow(query).Scan(&myExample.id, &myExample.data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data from Heroku:{ id:%d data:%s }", myExample.id, myExample.data)
	defer db.Close()
	err = db.Ping()
	if err != nil{
		panic(err)
	}
}