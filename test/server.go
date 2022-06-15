package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func testGet(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Working"))
}

func testString(w http.ResponseWriter, req *http.Request) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	fmt.Printf("--> String: [%s]\n", string(bodyBytes))
	w.Write([]byte("OK"))
}


func testJson(w http.ResponseWriter, req *http.Request) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	var data map[string]any
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	fmt.Println("--> Json:", data)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/testGet", testGet)

	http.HandleFunc("/testString", testString)
	http.HandleFunc("/testJson", testJson)

	http.ListenAndServe(":3000", nil)
}
