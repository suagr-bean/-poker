package service

import (
	"poker/ranges"
	"testing"
)
func TestCalHand(t*testing.T){
   hand:=[]string{"CA","HA"}
   data:=ranges.CalHand(hand)
   t.Log("数据",string(data))
}