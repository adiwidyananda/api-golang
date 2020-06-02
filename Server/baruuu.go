package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

var mySigningKey = []byte("mysupersecret")

type student struct {
	Nama     string `json:"Nama"`
	Nim      string `json:"Nim"`
	Jurusan  string `json:"Jurusan"`
	Angkatan string `json:"Angkatan"`
	username string `json:"username"`
	password string `json:"password"`
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/datamahasiswa")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getData() []student {
	var students []student
	db := dbConn()

	rows, err := db.Query("SELECT * FROM mahasiswa")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {

		var Nama string
		var Nim string
		var Jurusan string
		var Angkatan string
		var username string
		var password string

		err = rows.Scan(&Nama, &Nim, &Jurusan, &Angkatan, &username, &password)
		students = append(students, student{Nama: Nama, Nim: Nim, Jurusan: Jurusan, Angkatan: Angkatan, username: username, password: password})

	}
	defer db.Close()
	return students

}
func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Header["Token"] != nil {
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return mySigningKey, nil
		})

		if err != nil {
			fmt.Fprint(w, err.Error())
		}
		if token.Valid {
			var data = getData()
			json.NewEncoder(w).Encode(data)
		}
	} else {
		fmt.Println(w, "Not Authorized")
	}

}

func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Header["Token"] != nil {
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return mySigningKey, nil
		})

		if err != nil {
			fmt.Fprint(w, err.Error())
		}
		if token.Valid {
			var data = getData()
			parameter := mux.Vars(r)
			for _, item := range data {
				if item.Nim == parameter["nim"] {
					json.NewEncoder(w).Encode(item)
					return
				}
			}
			json.NewEncoder(w).Encode(&student{})
		}
	} else {
		fmt.Println(w, "Not Authorized")
	}
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Header["Token"] != nil {
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return mySigningKey, nil
		})

		if err != nil {
			fmt.Fprint(w, err.Error())
		}
		if token.Valid {
			r.ParseForm()
			db := dbConn()
			ins, err := db.Prepare("INSERT INTO mahasiswa(Nama, Nim, Jurusan, Angkatan, username, password) VALUES(?,?,?,?,?,?)")
			if err != nil {
				fmt.Fprintf(w, "Create data failed")
			}
			ins.Exec(r.FormValue("Nama"), r.FormValue("Nim"), r.FormValue("Jurusan"), r.FormValue("Angkatan"), r.FormValue("username"), r.FormValue("password"))
			defer db.Close()
			fmt.Fprintln(w, "Data Created")
		}
	} else {
		fmt.Println(w, "Not Authorized")
	}
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Header["Token"] != nil {
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return mySigningKey, nil
		})

		if err != nil {
			fmt.Fprint(w, err.Error())
		}
		if token.Valid {
			parameter := mux.Vars(r)
			r.ParseForm()
			db := dbConn()
			ins, err := db.Prepare("UPDATE mahasiswa SET Nama=?, Jurusan=?, Angkatan=?, username=?, password=? WHERE Nim=?")
			if err != nil {
				fmt.Fprintf(w, "Create data failed")
			}
			ins.Exec(r.FormValue("Nama"), r.FormValue("Jurusan"), r.FormValue("Angkatan"), r.FormValue("username"), r.FormValue("password"), parameter["nim"])
			defer db.Close()
			fmt.Fprintln(w, "data Update")
		}
	} else {
		fmt.Println(w, "Not Authorized")
	}
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Header["Token"] != nil {
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return mySigningKey, nil
		})

		if err != nil {
			fmt.Fprint(w, err.Error())
		}
		if token.Valid {
			parameter := mux.Vars(r)
			db := dbConn()
			ins, err := db.Prepare("DELETE FROM mahasiswa WHERE Nim=?")
			if err != nil {
				fmt.Fprintf(w, "Create data failed")
			}
			ins.Exec(parameter["nim"])
			defer db.Close()
			fmt.Fprintln(w, "Data Delete")
		}
	} else {
		fmt.Println(w, "Not Authorized")
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/students", getStudents).Methods("GET")
	r.HandleFunc("/api/students/{nim}", getStudent).Methods("GET")
	r.HandleFunc("/api/students", createStudent).Methods("POST")
	r.HandleFunc("/api/students/{nim}", updateStudent).Methods("PUT")
	r.HandleFunc("/api/students/{nim}", deleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
