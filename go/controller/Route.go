package controller

import (
	"fmt"
	"net/http"
)
func Route(){
	mux:=http.NewServeMux()
	mux.HandleFunc("/cal",CalRequest)
	fmt.Println("api is start")
	
	http.ListenAndServe(":8080",mux)
}