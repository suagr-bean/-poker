package utils

import (
	"math/rand/v2"
)
func (T*TexasDealer)ShuffleCard(){
	card:=T.Cards.AllCards
	rand.Shuffle(len(card),func(i,j int){
		card[i],card[j]=card[j],card[i]
	})
}