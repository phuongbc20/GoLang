package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/qlysinhvien")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	Create, err := db.Query("CREATE TABLE RESULT( NAME VARCHAR(255), SIZE VARCHAR(255), JSON TEXT, NAMEOUT VARCHAR(255) )")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Print("Create Done!")
	}
	defer Create.Close()

	/*var name string
	var size string
	err = db.QueryRow("SELECT NAME,SIZE FROM RESULT WHERE NAME=? AND SIZE=?", "b", "3").Scan(&name, &size)
	if err != nil {
		panic(err.Error())
	}
	//	_ = Rs.Scan(&name, &size)
	fmt.Println(name)*/
	//_, _ = db.Query("DELETE FROM result where name='in.JPG'")

}
