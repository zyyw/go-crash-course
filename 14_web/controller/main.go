package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "####")
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	url := "http://localhost:5001/api/v0/add"
	proxyReq, err := http.NewRequest(req.Method, url, bytes.NewReader(body))
	proxyReq.Header = make(http.Header)
	for h, val := range req.Header {
		proxyReq.Header[h] = val
	}

	httpClient := http.Client{}
	resp, err := httpClient.Do(proxyReq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

var redirectUrl = "http://localhost:3001"

var btfsUrl = "http://localhost:5001/api/v0/add"

//var btfsUrl = "http://localhost:5001"

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

func redirect2(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "http://www.google.com", 301)
}

func redirectToBTFS(w http.ResponseWriter, r *http.Request) {
	//r.Header.Add("Content-Type", "multipart/form-data")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	http.Redirect(w, r, btfsUrl, http.StatusPermanentRedirect)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("process file upload...")

	file, handle, err := r.FormFile("file[]")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	fmt.Printf("%s", data)
	fmt.Println("filename: ", handle.Filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	//err = ioutil.WriteFile("./files/"+handle.Filename, data, 0666)
	err = ioutil.WriteFile("./files/"+handle.Filename, data, 0666)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Printf("File uploaded successfully!.")
}

// UploadFile uploads a file to the server
func UploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	file, handle, err := r.FormFile("file")
	fmt.Println("filename", handle.Filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer file.Close()

	http.Redirect(w, r, "127.0.0.1:5001", http.StatusPermanentRedirect)

}

func main() {
	fmt.Println("controller server listen on port 3002...")
	//http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "controller server")
	//})
	http.HandleFunc("/", redirectToBTFS)
	//http.HandleFunc("/", handler)
	//http.HandleFunc("/r1", redirect)
	//http.HandleFunc("/r2", redirect2)
	//
	//http.HandleFunc("/upload", handleUpload)

	//http.HandleFunc("/", handler)
	http.ListenAndServe(":3002", nil)

}
