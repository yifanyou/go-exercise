package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", index)

	err := http.ListenAndServe(":21314", nil)
	if err != nil {
		log.Fatal("Start Server Err ", err)
	}
}

func index(res http.ResponseWriter, req *http.Request){
	fmt.Fprint(res, "WTF")
}
