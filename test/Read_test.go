package service

import (
	
	"poker/ranges"
	"testing"
)

func TestReadFile(t *testing.T){
   path:="../../Range/Hands.jsonl"
	data,err:=ranges.ReadFile(path)
  if err!=nil{
    t.Log("错误",err)
  }
 for _,v:=range data.HandIndex{
  t.Log("datawin",v)
 }
  
}