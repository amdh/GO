package main

import (
	"io"
	"net/http"
	"fmt"
	"net/url"
)
import "io/ioutil"

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!  @8000")
}

func idleTool(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" tool is idle requesting move..")

	requestMove();
	requestAdd();
}

func requestAdd(){

	resp, err := http.PostForm("http://10.104.102.122:9000/add",
		url.Values{"val1": {"9"}, "val2": {"123"}})

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		responseString := string(body)
		fmt.Println("post:\n", responseString)
}

func requestMove(){
	resp, err := http.Get("http://10.104.102.122:9000/move")
	if err != nil {
		panic(err)
	}

	fmt.Println("returned data from move:\n", resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	responseString := string(body)
	fmt.Println("returned data from move:\n", responseString)
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/idleTool", idleTool)



	fmt.Println("go server started at 8000")
	http.ListenAndServe(":8000", nil)
}
