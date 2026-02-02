package Service

import "poker/model"
type TexasDealer struct{
	Deck *model.Deck
}

func NewTexasDealer()*TexasDealer{
	deck:=model.NewDeck()
	return &TexasDealer{Deck:deck}
}


