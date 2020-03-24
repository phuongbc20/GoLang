package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type ToaDo struct {
	X int `json:"X"`
	Y int `json:"Y"`
}
type OutputJson struct {
	Min ToaDo `json:"Min"`
	Max ToaDo `json:"Max"`
}

//--------------------------------------
//Function MAIN
//--------------------------------------
func main() {
	http.HandleFunc("/facedetection", Action)
	http.ListenAndServe(":8080", nil)
}

//--------------------------------------
//Function MAIN
//--------------------------------------
func Action(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 * 1024 * 1024)
	_, a, _ := r.FormFile("File")
	f, _ := a.Open()
	var Find bool
	Find = dbFind(a.Filename, strconv.FormatInt(a.Size, 10))
	if Find == true {
		tempFile, _ := ioutil.TempFile("Uploads", "image-*.jpg")
		defer tempFile.Close()
		fileBytes, _ := ioutil.ReadAll(f)
		tempFile.Write(fileBytes)
		cmd := exec.Command("pigo", "-in", tempFile.Name(), "-json", "-out", tempFile.Name(), "-cf", "github.com/esimov/pigo/cascade/facefinder")
		_ = cmd.Run()
		dbInsert(a.Filename, strconv.FormatInt(a.Size, 10), tempFile.Name())
		var FileName string
		if r.FormValue("Type") == "1" {
			FileName = "PathOut"
			fmt.Println(FileName)
			file, err := os.Open(tempFile.Name())
			if err != nil {
				log.Fatalln(err)
			}
			defer file.Close()
			io.Copy(w, file)
		} else if r.FormValue("Type") == "2" {
			plan, _ := ioutil.ReadFile("output.json")
			var data []OutputJson
			err := json.Unmarshal(plan, &data)
			if err != nil {
				fmt.Print("Cannot unmarshal the json ", err)
			}
			json.NewEncoder(w).Encode(data)
		}

	} else {
		if r.FormValue("Type") == "1" {
			file, err := os.Open(dbOutIg(a.Filename, strconv.FormatInt(a.Size, 10)))
			if err != nil {
				log.Fatalln(err)
			}
			defer file.Close()
			io.Copy(w, file)
		} else if r.FormValue("Type") == "2" {
			plan := []byte(dbOutJson(a.Filename, strconv.FormatInt(a.Size, 10)))
			var data []OutputJson
			err := json.Unmarshal(plan, &data)
			if err != nil {
				fmt.Print("Cannot unmarshal the json ", err)
			}
			json.NewEncoder(w).Encode(data)
			//var out []byte
			//out, _ := json.Marshal(data)
			//fmt.Print(string(out))
		}
	}

	//-----------------------------

}

func dbFind(Name string, Size string) bool {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(192.168.1.8:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var name string
	err = db.QueryRow("SELECT NAME FROM result WHERE NAME=? AND SIZE=?", Name, Size).Scan(&name)
	if err != nil {
		return true
	} else {
		fmt.Println("Image da ton tai!")
		return false
	}

}

func dbInsert(Name string, Size string, Nameout string) {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(192.168.1.8:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	plan, _ := ioutil.ReadFile("output.json")
	//fmt.Print(string(plan))
	Create, err := db.Query("insert into result values (?,?,?,?)", Name, Size, string(plan), Nameout)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Print("Insert Done!")
	}
	defer Create.Close()
}
func dbOutIg(Name string, Size string) string {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(192.168.1.8:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var name string
	_ = db.QueryRow("SELECT NAMEOUT FROM result WHERE NAME=? AND SIZE=?", Name, Size).Scan(&name)
	return name
}
func dbOutJson(Name string, Size string) string {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(192.168.1.8:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var Json string
	_ = db.QueryRow("SELECT JSON FROM result WHERE NAME=? AND SIZE=?", Name, Size).Scan(&Json)
	return Json
}
