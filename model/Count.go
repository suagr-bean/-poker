package model
type Count struct {
	 
     Cards[]int
	 Number[]int
	 Color []int
     CountNumber[]int
	 CountColor[]int
     MaskNumber uint16
	 MaskColor [4]uint16
}
func NewCount(cards[]int)*Count{
	 count:=&Count{
       Cards:cards,
	   Number:make([]int,0),
       Color:make([]int,0),
	   CountNumber:make([]int,0),
	   CountColor:make([]int,0),
	 }
	 return count
}
func (c*Count)SetCount(countcolor[]int,countnumber[]int){
	c.CountColor=append(c.CountColor, countcolor...)
	c.CountNumber=append(c.CountNumber, countnumber...)
}