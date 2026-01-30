package model
//公牌
type Board struct {
	boardCards[]int
}

func NewBoard(cards []int)*Board{
   b:=&Board{}
   b.boardCards=append(b.boardCards,cards...)
   return b
}
func (b*Board)GetCards()[]int{
    return b.boardCards
}


