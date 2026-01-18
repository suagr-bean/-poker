package controller

import (
	"poker/model"
	"poker/service"
	"fmt"
)
func Begining(){
	win:=float32(0)
	beginner:=&model.Begin{}
	beginner.Hand=[]int{58,59}
	beginner.Id=0
	beginner.PublicCard=[]int{}
	beginner.Person=6
    beginner.Frequency=100000
	result:=func(s float32){
		win+=s
	}
	for i:=0;i<beginner.Frequency;i++{
		service.GameMain(beginner,result)
	}
	fmt.Println((float32(win)/float32(beginner.Frequency))*100)
}