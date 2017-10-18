package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
	"net/url"
	
)


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/status", getStatus).Methods("GET")
	router.HandleFunc("/server2status", getStatusFromServer2).Methods("GET")
	router.HandleFunc("/move", move).Methods("GET")
	router.HandleFunc("/move", move).Methods("PATCH")
	router.HandleFunc("/report/{zipcode}", report).Methods("GET")
	router.HandleFunc("/forcast", forcast).Methods("POST")
	http.ListenAndServe(":8080", router)
}

func getStatus(res http.ResponseWriter, req *http.Request) {
	//res.Header().Set("Content-Type", "application/json")

	 st := map[string]string{"status" : "idle"}

	outgoingJSON, error := json.Marshal(st)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, string(outgoingJSON))
}


func getStatusFromServer2(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	resp, err := http.Get("http://10.104.98.230:9090/status")
	if err != nil {
		panic(err)
	}

	fmt.Println("returned data from move:\n", resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	responseString := string(body)
	fmt.Println("returned data from move:\n", responseString)

	st := map[string]string{"server2 status": responseString }

	outgoingJSON, error := json.Marshal(st)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, string(outgoingJSON))
}

func move(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	resp, err := http.Get("http://localhost:9090/move")
	if err != nil {
		panic(err)
	}

	fmt.Println("returned data from move:\n", resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	responseString := string(body)
	fmt.Println("returned data from move:\n", responseString)

	st := map[string]string{"'move status'" : responseString}

	outgoingJSON, error := json.Marshal(st)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, string(outgoingJSON))
}


func report(res http.ResponseWriter, req *http.Request) {
	//res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	zipcodestr := vars["zipcode"]
	
	resp, err := http.Get("http://10.104.98.230:9090/report/"+zipcodestr)

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		responseString := string(body)
		fmt.Println("get:\n", responseString)
		fmt.Fprint(res, responseString)
}

func forcast(res http.ResponseWriter, req *http.Request) {
	//res.Header().Set("Content-Type", "application/json")


	cityStr := req.FormValue("City")
	daysStr := req.FormValue("days")
	
	fmt.Println(cityStr)
	fmt.Println(daysStr)


	resp, err := http.PostForm("http://10.104.98.230:9090/forcast",
		url.Values{"City": {cityStr}, "days": {daysStr}})

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		responseString := string(body)
		fmt.Println("post:\n", responseString)
		fmt.Fprint(res, responseString)
	}
	
