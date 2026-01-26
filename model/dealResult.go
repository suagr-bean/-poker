package model
type EntryInfor struct {
	card [] int
   	
}
func NewEntryInfor(cards[]int)*EntryInfor{
    e:=&EntryInfor{card:cards}
	return e
}
func(e*EntryInfor)SetEntry(cards[]int){
   e.card=append(e.card,cards...)
  
}
func (e*EntryInfor)GetEntryCards()[]int{
	return e.card
}