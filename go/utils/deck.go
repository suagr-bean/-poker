package utils

import "poker/model"
type TexasDealer struct{
	Board *model.Board
	Data *model.GameData
	Cards *model.Deck
	Id int
}
func NewTexasDealer()*TexasDealer{
	return &TexasDealer{}
}

func NewDeck()*model.Deck{
    d:=&model.Deck{SendCard:make([]int,0),Knowns:make([]int,0)}
	
	d.Color=[]int{0,1,2,3}
	d.Number=make([]int,15)
    for i:=2;i<15;i++{
	d.Number[i]=i
	}
	d.AllCards=make([]int,52)
	index:=0
   for i:=2;i<15;i++{
	for j:=0;j<4;j++{
       d.AllCards[index]=d.Number[i]<<2|d.Color[j]
	   index++
	}
   }
   return d
   
}