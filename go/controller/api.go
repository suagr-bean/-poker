package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"poker/model/httpModel"
) 
func Route() {
	mux:=http.NewServeMux()
	mux.HandleFunc("POST /cal",DealRequest)
	mux.HandleFunc("GET /go",Go)
	log.Fatal(http.ListenAndServe(":8080",mux))
}
func DealRequest(w http.ResponseWriter, req*http.Request){
	fmt.Println("收到请求")
	data:=httpModel.NewData()
	err:=json.NewDecoder(req.Body).Decode(data)
	if err!=nil{
		return 
	}
	fmt.Println(data)
	
	 reqs:=Begining(data)//计算结果
	 result,err:=json.Marshal(reqs)
	w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)

	w.Write([]byte(result))
    
}
func Go(w http.ResponseWriter,req*http.Request){
	fmt.Println("第二个请求收到")
	io.WriteString(w,"i love go ")
}