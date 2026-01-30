package service

import (
	"poker/model"
	"testing"
)
//测试牌堆是否正常
func TestDeck(t *testing.T){
   deck:=model.NewDeck()
   t.Log(len(deck.Cards))
}