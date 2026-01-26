package model
//公牌
type Board struct {
	boardCards[]int
   Entry []*EntryInfor
}
type Process struct{
   id int
   allProcess[]int
}
func NewBoard(cards []int,person int)*Board{
   b:=&Board{boardCards:make([]int,0),
   Entry:make([]*EntryInfor,0)}
   b.boardCards=append(b.boardCards,cards...)
   return b
}
func (b*Board)ProcessCard(id int,hand []int)*Process{
   p:=&Process{id:id,allProcess:make([]int,0)}
   p.allProcess=append(p.allProcess,b.boardCards...)
   p.allProcess=append(p.allProcess,hand...)
   return p
}
func(b*Board) GetBoardCards()[]int{
	return b.boardCards
}
func(b*Board)Add(card[]int){
  b.boardCards=append(b.boardCards,card...)
}
func(b*Board)SetEntry(id int, cards[]int){
   b.Entry[id].card=cards
}


