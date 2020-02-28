package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Res struct {
	Names string `json:"name"`
	//	Size  int64  `json:"size"`
	Path string `json:"path"`
}

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
	files := r.MultipartForm.File["File"]
	for _, a := range files {
		f, _ := a.Open()
		name := a.Filename[0 : len(a.Filename)-4]
		Type := a.Filename[len(a.Filename)-4 : len(a.Filename)]
		tempFile, _ := ioutil.TempFile("Uploads", name+"*"+Type)
		defer tempFile.Close()
		fileBytes, _ := ioutil.ReadAll(f)
		tempFile.Write(fileBytes)
		//-----------------------------
		resp := &Res{
			Names: a.Filename,
			//	Size:  a.Size,
			Path: tempFile.Name(),
		}
		json.NewEncoder(w).Encode(resp)
		//-----------------------------
	}
}
