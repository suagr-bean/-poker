package model

type Score interface{
	Node(d*Count)(bool,int)
}
type Chain struct{
   Handler[]Score
   
}

func (c*Chain)Add(h Score)*Chain{
  c.Handler=append(c.Handler,h)
   return c
}
func (c*Chain)Start(d*Count)(bool,int){
	for _,h:=range c.Handler{
		deal,point:=h.Node(d)
		if deal{
		 return deal,point
		}
	}
	return  false,0
}