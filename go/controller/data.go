package controller

import (
	"fmt"
	"poker/model"
	"poker/model/httpModel"
	"poker/service"
	"poker/utils"
)
func Begining(cal*httpModel.CalData)float32{
	var req float32
	 Hand:=cal.Cards.Hand
	 dealHand:=utils.DealMap(Hand)
     dealPublic:=utils.DealMap(cal.Cards.Hand)
	 fmt.Println(dealHand)
	win:=float32(0)
	beginner:=&model.Begin{}
	beginner.Hand=dealHand
	beginner.Id=cal.Cards.Position
	beginner.PublicCard=dealPublic
	beginner.Person=cal.Table.Person
    beginner.Frequency=50000
	result:=func(s float32){
		win+=s
	}
	for i:=0;i<beginner.Frequency;i++{
		service.GameMain(beginner,result)
	}
	req=(float32(win)/float32(beginner.Frequency))*100
	fmt.Println(req)
    return req
}