package model 

type Dealer interface{
     Init(b*Board,data*GameData)
	}
type Decker interface{
	ShuffleCard()
    Dealing(need int)[]int
	Known(pool[]int)
	Process(g*GameData,d*Deck,b*Board)   
}
