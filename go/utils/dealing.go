package utils
import "poker/model"
func(T*TexasDealer) Dealing(need int)[]int{
	cards:=T.Cards.Get()
    card:=cards[:need]
	Allcards:=d.AllCards[need:]
	d.SetAllCards(Allcards)
	d.SendCard=append(d.SendCard,card...)
	return card
	
}