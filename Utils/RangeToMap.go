package Utils
//处理好手牌 过map 同时把已知手牌过滤掉
func RangeToMap(chance int,cards[]int)AllHand{
   
	hand:=PlayerRange(chance)
	OuterLoop:
	for i:=0;i<len(hand.HandIndex);i++{
		pool:=hand.HandIndex[i].Hand
		result:=DealMap(pool)
		
		 for _,j:=range cards{
			if result[0]==j||result[1]==j{
			  continue  OuterLoop
			}
		 } 
		
		hand.HandIndex[i].MapHand=result
	}
    j:=0
   for i:=0;i<len(hand.HandIndex);i++{ 
	if hand.HandIndex[i].MapHand != nil{
       hand.HandIndex[j]=hand.HandIndex[i]
	   j++
	}
   }
   hand.HandIndex=hand.HandIndex[:j]
   return hand
}