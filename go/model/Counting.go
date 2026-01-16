package model
type CountId struct{
	Id []*Detail
}
type Detail struct {
     Cards[]int
	 Number[]int
	 Color []int
}
func NewCount()*CountId{
	c:=&CountId{Id:make([]*Detail,0)}
	 detail:=&Detail{
       Cards:make([]int,0),
	   Number:make([]int,0),
       Color:make([]int,0),
	 }
	 c.Id=append(c.Id,detail)
	return c
}
func (Count*CountId)Set(d*Detail){
    Count.Id=append(Count.Id,d)
}
func (Count*CountId)Get()[]*Detail{
	return Count.Id
}
func (Count*CountId)SetAll(color[]int,number[]int,cards[]int){
  d:=&Detail{}
  d.Cards=cards
  d.Color=color
  d.Number=number
  Count.Id=append(Count.Id,d)
}
func (Count*CountId)GetAll()[]*Detail{
	return Count.Id
}
