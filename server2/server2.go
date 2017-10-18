package main

import (
	"io"
	"net/http"
	"fmt"
)
import "io/ioutil"
import "strconv"

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!   @ 9000")
}

func moveXY(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("Successfully moved.")
	fmt.Fprintf(w, string("Successfully moved."))

}

func add(w http.ResponseWriter, req *http.Request){
	fmt.Println("adding two paramters from post request..")
	

	val1 := req.FormValue("val1")
	val2 := req.FormValue("val2")
	fmt.Println("value 1 : ", val1)

	
	v1, err := strconv.Atoi(val1)//because the put request requires an integer
    if err != nil {
        panic(err)
	}
	
	v2, err := strconv.Atoi(val2)//because the put request requires an integer
    if err != nil {
        panic(err)
	}
	
	ans := v1 + v2
	fmt.Println("answer is ", ans)
	fmt.Fprintf(w, "answer is ", ans)
	
}

func requestIdleTool(w http.ResponseWriter, r *http.Request){
	resp, err := http.Get("http://10.104.100.168:8000/idleTool")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("returned data from idletool:\n", body)
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/move", moveXY)
	http.HandleFunc("/add", add)
	http.HandleFunc("/start", requestIdleTool)


	fmt.Println("go server started at 9000")
	http.ListenAndServe(":9000", nil)


	
}
