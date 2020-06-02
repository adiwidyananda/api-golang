package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type students struct {
	Nama     string `json:"Nama"`
	Nim      string `json:"Nim"`
	Jurusan  string `json:"Jurusan"`
	Angkatan string `json:"Angkatan"`
	username string `json:"username"`
	password string `json:"password"`
}

type DetailInfo struct {
	Nama     string
	Jurusan  string
	Angkatan string
}

type info struct {
	Judul  string
	Detail map[string]DetailInfo
}

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTEwNjU3ODAsInVzZXIiOiJFbGxpb3QgRm9yYmVzIn0.z2pXx-O9bBI-W0NFuQmpKu9tMIqPiGZhUi-DPBFR73g"

func index_handler(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8000/api/students"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Token", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var s []students
	datastudent := make(map[string]DetailInfo)
	// fmt.Println(body)

	if err := json.Unmarshal(body, &s); err != nil {
		log.Fatal(err)
	}

	for idx := range s {
		datastudent[s[idx].Nim] = DetailInfo{s[idx].Nama, s[idx].Jurusan, s[idx].Angkatan}
	}

	// bytes := nimData
	// data_mahasiswa := make(map[string]DetailInfo)
	// var s Info
	// if err := xml.Unmarshal(bytes, &s); err != nil {
	// 	log.Fatal(err)
	// }
	// for idx, _ := range s.Nama{
	// 	data_mahasiswa[s.Nama[idx]] = DetailInfo{s.Jurusan[idx], s.Kelas[idx], s.Angkatan[idx]}
	// }
	p := info{Judul: "Data Mahasiswa", Detail: datastudent}
	t, _ := template.ParseFiles("baru.html")
	fmt.Println(t.Execute(w, p))
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8001", nil)

	// url := "http://localhost:8000/api/students"
	// method := "GET"

	// client := &http.Client{}
	// req, err := http.NewRequest(method, url, nil)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// req.Header.Add("Token", token)

	// res, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer res.Body.Close()
	// body, err := ioutil.ReadAll(res.Body)

	// var s []students
	// datastudent := make(map[string]DetailInfo)
	// // fmt.Println(body)

	// if err := json.Unmarshal(body, &s); err != nil {
	// 	log.Fatal(err)
	// }

	// for idx := range s {
	// 	datastudent[s[idx].Nim] = DetailInfo{s[idx].Nama, s[idx].Jurusan, s[idx].Angkatan}
	// }

	// fmt.Println(datastudent)
}
