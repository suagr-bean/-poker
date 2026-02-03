package service

import (
	"poker/Utils"
	"testing"
)
//测试抽出的手牌
func TestRandHand(t*testing.T){
	chance:=20
    cards,_:=Utils.Rand(chance,[]int{20,25})
	t.Log(cards)
}