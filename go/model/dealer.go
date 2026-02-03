package model 

type Dealer interface{
	  DealHand(p*PlayerInfo,b*Board)*Pool
}
type Decker interface{
	ShuffleCard()
    Dealing(need int)[]int
	Known()
	Process()   
}
