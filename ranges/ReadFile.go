package ranges

import (
	"bufio"
	"encoding/json"
	"os"
)
type AllHand struct {
	HandIndex []Hands
}
type Hands struct {
	Hand []string `json:"hand"`
	Win float32  `json:"win,string"`
}
func ReadFile(path string)(AllHand,error){
	allhand:=AllHand{}
    file,err:=os.Open(path)
	if err!=nil{
	return allhand,err
	}
    defer file.Close()
	scanner:=bufio.NewScanner(file)
	   
    for scanner.Scan(){
		line:=scanner.Bytes()
		h:=Hands{}
      json.Unmarshal(line,&h)
	  allhand.HandIndex = append(allhand.HandIndex,h)
	}
	
	return allhand,nil
}