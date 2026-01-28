package service

import (
	"poker/ranges"
	"testing"
)
func TestPost(t *testing.T){
	hand:=[]string{"CA","D5"}
	data:=ranges.CalHand(hand)
	url:="https://fluffy-fortnight-jj64qv579r7g2q4xr-8080.app.github.dev/cal"
	result,_:=ranges.PostStart(url,data)
	t.Log("结果",result)
}