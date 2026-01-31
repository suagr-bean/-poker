package model



type Pool struct {
   DeletePool [] int
}
func  NewPool (p*PlayerInfo,b*Board)*Pool{
	pool:=&Pool{}
	if b.GetCards()!=nil{
    pool.DeletePool=append(pool.DeletePool, b.GetCards()...)//删除公牌
	}
	a:
	for i:=0;i<p.Person;i++{
	   hand:=p.Players[i].Hand
	   if hand==nil{
		continue a
	   }
	  pool.DeletePool=append(pool.DeletePool, hand...)//删除手牌
	}
	 return pool
}
//添加要删除的牌
func (p*Pool)Add(pool[]int){
	if pool==nil{
		return  
	}
     p.DeletePool=append(p.DeletePool, pool...)
}
//获取已经发出的牌
func (p*Pool)GetPool()[]int{
	return p.DeletePool
}


