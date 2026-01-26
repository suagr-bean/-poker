package model 
type Deck struct {
	AllCards[]int
	Number[]int
	Color []int
	Knowns []int
	SendCard[]int
}
func (d*Deck)SetAllCards(allcards []int){
	d.AllCards=allcards
}
func(d*Deck)Get()[]int{
	return d.AllCards
}
