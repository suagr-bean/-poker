package utils

func (T*TexasDealer)Known(){
	pool:=T.DeletePool()
    d:=T.Cards
	if len(pool)!=0{
     for _,v:=range pool{
		for j:=len(d.AllCards)-1;j>=0;j--{
			if v==d.AllCards[j]{
				d.Knowns=append(d.Knowns,d.AllCards[j])
               last:=len(d.AllCards)-1
				d.AllCards[j]=d.AllCards[last]
				d.AllCards=d.AllCards[:last]
				break
			}
		}
	 }
	}
  
}
