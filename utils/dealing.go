package utils

func(T*TexasDealer) Dealing(need int)[]int{
	cards:=T.Cards.Get()
    card:=cards[:need]
	T.Cards.AllCards=T.Cards.AllCards[need:]
	
	T.Cards.SendCard=append(T.Cards.SendCard,card...)
	return card	
}