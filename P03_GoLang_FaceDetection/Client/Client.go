package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

//--------------------------------------
//Function MAIN
//--------------------------------------
func main() {
	//--------------------------------------
	var Path, Type string
	input(&Path, &Type)
	//--------------------------------------
	MakeRequest(Path, Type)
}

//--------------------------------------
//Function input
//--------------------------------------
func input(Path *string, Type *string) {
	cin := bufio.NewScanner(os.Stdin)
	fmt.Print("Path(vd: C:/Users/ASUS/go/src/Final/image/in.jpg ) :")
	cin.Scan()
	*Path = cin.Text()
	fmt.Print("Type(1:Image || 2:Json ): ")
	cin.Scan()
	*Type = cin.Text()
}

//--------------------------------------
//Function tao request
//--------------------------------------

func MakeRequest(PathFile string, Type string) {
	url := "http://192.168.1.13:8080/facedetection"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(PathFile)
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("File", "input.png")
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
	}
	_ = writer.WriteField("Type", Type)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if Type == "1" {
		f, _ := os.Create("Image/output.jpg")
		_, _ = io.Copy(f, res.Body)
	} else if Type == "2" {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
}
