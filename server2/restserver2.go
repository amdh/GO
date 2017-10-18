package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	 "strconv"
	"github.com/gorilla/mux"
	"io/ioutil"
)
type Weather  struct {
    zipcode        int   `json:"zipcode,omitempty"`
    City		 string   `json:"city,omitempty"`
	day  string   `json:"day,omitempty"`
	temperature  int   `json:"celius,omitempty"`
	description  string   `json:"lastname,omitempty"`
}

var wlist []Weather


func main() {
	wlist = append(wlist, Weather{zipcode: 95112 , City:"San Jose", day:"1", temperature:83, description:"sunny"})
	wlist = append(wlist, Weather{zipcode: 95118 , City:"San Jose", day:"2", temperature:61, description:"cloudy"})
	wlist = append(wlist, Weather{zipcode: 94032 , City:"Santa Clara", day:"1", temperature:76, description:"sunny"})
	wlist = append(wlist, Weather{zipcode: 94085 , City:"Sunnyvale", day:"1", temperature:74, description:"Cloudy"})


	router := mux.NewRouter()
	router.HandleFunc("/status", getStatus).Methods("GET")
	router.HandleFunc("/server1status", getStatusFromServer1).Methods("GET")
	router.HandleFunc("/move", move).Methods("GET")
	router.HandleFunc("/report/{zipcode}", report).Methods("GET")
	router.HandleFunc("/forcast", forcast).Methods("POST")

	fmt.Println("server started at 9090.....")
	http.ListenAndServe(":9090", router)
}

func getStatus(res http.ResponseWriter, req *http.Request) {
	//res.Header().Set("Content-Type", "application/json")

	 st := map[string]string{"status" : "running with no tasks"}

	outgoingJSON, error := json.Marshal(st)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, string(outgoingJSON))
}

func getStatusFromServer1(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	resp, err := http.Get("http://10.35.240.121:8080/status")
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
	//res.Header().Set("Content-Type", "application/json")

	 st := map[string]string{"status" : "moved item x to y"}

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
	zipcode, err := strconv.Atoi(zipcodestr)
	 if err != nil {
        // handle error
        fmt.Println(err)
        return
    }
	fmt.Println(zipcode)
	st :=  make(map[string]string)

	for i, wObj := range wlist {
        if wObj.zipcode == zipcode {
			fmt.Println("index:", i)
			fmt.Println(wObj)
			
			wObjStr := "city: " + wObj.City + " day: " +wObj.day + " description: " +wObj.description
			fmt.Println(wObjStr)
			z := strconv.Itoa(zipcode)
			st[z] = wObjStr
			break
        }
	}
	
	fmt.Println("return object: ", st)


	outgoingJSON, error := json.Marshal(st)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, string(outgoingJSON))
}

func forcast(res http.ResponseWriter, req *http.Request) {
	//res.Header().Set("Content-Type", "application/json")


	cityStr := req.FormValue("City")
	daysStr := req.FormValue("days")
	
	fmt.Println(cityStr)
	fmt.Println(daysStr)
	days, err := strconv.Atoi(daysStr)
	if err != nil {
	   // handle error
	   fmt.Println(err)
	   return
   }
	fmt.Println(days)
	st :=  make(map[string]string)
	temp := 0
		for i, wObj := range wlist {
			if wObj.City == cityStr {
				fmt.Println("index:", i)
				fmt.Println(wObj)
				temp += wObj.temperature
				fmt.Println(temp)
			}
		}
		forcastTemp := temp/days
		fmt.Println(forcastTemp)
		desc := "unknown"
		if forcastTemp >= 70 {
			desc = "sunny"
		} else if forcastTemp < 60 {
			desc = "cloudy"
		} else {
			desc = "windy"
		}


	
		wObjStr := " Weather forecast  with temperature "+ 	strconv.Itoa(forcastTemp) +". It would be mostly " + desc
		fmt.Println(wObjStr)
	
		st[cityStr] = wObjStr
		
		fmt.Println("return object: ", st)

	outgoingJSON, error := json.Marshal(st)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, string(outgoingJSON))
}

