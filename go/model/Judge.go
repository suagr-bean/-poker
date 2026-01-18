package model
type Judge interface{
	DealCount()
	InitCard(id int,card[][]int)
	CallBack(func(score float32))
}
type Score interface{
	Processing(d*Detail)(bool,int)
}
type Chain struct{
   Handler[]Score
}

func (c*Chain)Add(h Score)*Chain{
  c.Handler=append(c.Handler,h)
   return c
}
func (c*Chain)Process(d*Detail)(bool,int){
	for _,h:=range c.Handler{
		deal,point:=h.Processing(d)
		if deal{
		 return deal,point
		}
	}
	return  false,-100
}