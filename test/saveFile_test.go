package service

import (
	
	"poker/ranges"
	"testing"
)
func TestSave(t*testing.T){
	dir:="../../Range"
	fileName:="range.jsonl"
    data:=[]string{"hello,i love go"}
	err:=ranges.SaveFile(dir,fileName,data,"hello-go","10")
	if err!=nil{
		t.Log("错误",err)
	}
}