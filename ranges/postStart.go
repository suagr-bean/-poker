package ranges

import (
	"bytes"
	"io"
	"net/http"
	"fmt"
)
func PostStart(url string ,data []byte)(string,error){
	posturl:=url
    resp,err:=http.Post(posturl,"application/json",bytes.NewBuffer(data))
	if err!=nil{
		fmt.Println("发送失败")
		return "",err
	}
	defer resp.Body.Close()
	body,_:=io.ReadAll(resp.Body)
	result:=string(body)

	return result,nil
}