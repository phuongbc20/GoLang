package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//--------------------------------------
//Function MAIN
//--------------------------------------
func main() {
	//--------------------------------------
	var Path, Type, ITime, FTime string
	input(&Path, &Type, &ITime, &FTime)
	//--------------------------------------
	t, _ := strconv.Atoi(ITime)
	FT, _ := strconv.Atoi(FTime)
	TimeStop := time.Now()
	for time.Since(TimeStop).Seconds() < float64(FT) {
		Action(Path, Type)
		time.Sleep(time.Duration(t) * time.Second)

	}
}

//--------------------------------------
//Function input
//--------------------------------------
func input(Path *string, Type *string, ITime *string, FTime *string) {
	cin := bufio.NewScanner(os.Stdin)
	fmt.Print("Path(vd: C:/Users/ASUS/go/src/Final/image ) :")
	cin.Scan()
	*Path = cin.Text()
	fmt.Print("Type(1:sequentially || 2:concurrency ): ")
	cin.Scan()
	*Type = cin.Text()
	fmt.Print("Interval Time(sec): ")
	cin.Scan()
	*ITime = cin.Text()
	fmt.Print("Total Processing time(sec): ")
	cin.Scan()
	*FTime = cin.Text()
}

//--------------------------------------
//Function tao request
//--------------------------------------
func Action(Path string, Type string) {
	f, err := os.Open(Path)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	//--------------------------------------
	if Type == "1" {
		for _, file := range files {
			MakeRequest(Path + "/" + file.Name())
		}
	} else if Type == "2" {
		ch := make(chan bool)
		for _, file := range files {
			go MakeRequestConcurrency(ch, Path+"/"+file.Name())
		}
		for _, a := range files {
			if a == nil {
				fmt.Println(nil)
			}
			<-ch
		}
	}
}
func MakeRequest(PathFile string) {
	file, err := os.Open(PathFile)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()
	requestBody := &bytes.Buffer{}

	multiPartWriter := multipart.NewWriter(requestBody)

	fileWriter, err := multiPartWriter.CreateFormFile("File", filepath.Base(PathFile))
	if err != nil {
		log.Fatalln(err)
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		log.Fatalln(err)
	}

	multiPartWriter.Close()

	req, err := http.NewRequest("POST", "http://192.168.1.13:8080", requestBody)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	fmt.Println(result)
}

//--------------------------------------
//Function to concurrency request
//--------------------------------------
func MakeRequestConcurrency(ch chan bool, PathFile string) {
	MakeRequest(PathFile)
	ch <- true
}
