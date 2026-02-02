package  Utils

//根据范围截取手牌
func PlayerRange(chance int)AllHand{
   allhand:=RangeTable()
    
   index:=int((len(allhand.HandIndex)*chance)/100)
   result:=AllHand{
     HandIndex: allhand.HandIndex[:index],
   }

   return result 
}
