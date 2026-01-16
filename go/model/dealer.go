package model 

type Dealer interface{
     Init(b*Board,data*GameData)[][]int
	}
type Decker interface{
	ShuffleCard()
    Dealing(need int)[]int
	Known()
	Process()   
}
