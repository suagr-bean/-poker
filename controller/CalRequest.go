package controller

import (
	"encoding/json"
	"net/http"
	"poker/model/httpModel"
) 

//处理德扑计算请求
func CalRequest(w http.ResponseWriter, req*http.Request){
	
	data:=httpModel.NewData()
	err:=json.NewDecoder(req.Body).Decode(data)
	if err!=nil{//解析JSON 失败 数据不合法
	 http.Error(w,"数据不合法",http.StatusBadRequest)
		return 
	}
	 
	 reqs:=PokerHandler(data)//计算结果
	 result,err:=json.Marshal(reqs)
	 if err!=nil{
		return 
	 }
	w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)

	w.Write([]byte(result))
 
}
