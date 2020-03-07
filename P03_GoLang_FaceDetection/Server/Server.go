package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

//--------------------------------------
//Function MAIN
//--------------------------------------
func main() {
	http.HandleFunc("/", Action)
	http.ListenAndServe(":8080", nil)
}

//--------------------------------------
//Function MAIN
//--------------------------------------
func Action(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 * 1024 * 1024)
	_, a, _ := r.FormFile("File")
	f, _ := a.Open()
	tempFile, _ := ioutil.TempFile("Uploads", "input-*.jpg")
	defer tempFile.Close()
	fileBytes, _ := ioutil.ReadAll(f)
	tempFile.Write(fileBytes)
	cmd := exec.Command("pigo", "-in", tempFile.Name(), "-json", "-out", "out.jpg", "-cf", "github.com/esimov/pigo/cascade/facefinder")
	_ = cmd.Run()
	fmt.Println(tempFile.Name())
	//-----------------------------
	var FileName string
	if r.FormValue("Type") == "1" {
		FileName = "out.jpg"
	} else if r.FormValue("Type") == "2" {
		FileName = "output.json"
	}
	fmt.Println(FileName)
	file, err := os.Open(FileName)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	io.Copy(w, file)

}
