package Utils

import "math/rand"
//根据范围随机抽一组牌 
func RandHands(chance int,cards[]int)[]int{
	result:=RangeToMap(chance,cards)
   length:=len(result.HandIndex)
   index:=rand.Intn(length)
   card:=result.HandIndex[index].MapHand
   
   return card
}