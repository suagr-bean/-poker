package model 
type Deck struct {
	Cards []int
	Color []int
	Number[]int
}
func NewDeck()*Deck{
	deck:=&Deck{}
	deck.Color=[]int{0,1,2,3}
	deck.Number=make([]int,13)
	for i:=2;i<15;i++{
	 deck.Number[i-2]=i
	}
	for _,v:=range deck.Number{
	 for _,j:=range deck.Color{
	   cards:=v<<2|j
	   deck.Cards=append(deck.Cards,cards) 
	 }
	}
	return deck
}
