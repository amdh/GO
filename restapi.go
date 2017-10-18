package main

import (
    "encoding/json"
    "io"
    "net/http"
    "fmt"
 
    "github.com/gorilla/mux"
)

type Employee struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
}

var emp []Employee

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func GetEmployee(w http.ResponseWriter, req *http.Request){
	 json.NewEncoder(w).Encode(emp)
}

func main() {

	router := mux.NewRouter()
	emp = append(emp, Employee{ID: "1", Firstname: "Rob", Lastname: "Roy"})

	router.HandleFunc("/emp", GetEmployee).Methods("GET")

	http.HandleFunc("/", hello)
	fmt.Println("go server started at 8000")
	http.ListenAndServe(":8000", router)
}
