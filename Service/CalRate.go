package Service

import (
	"poker/model"
	"poker/model/httpModel"
	
	"poker/Utils"
)

//业务入口 蒙特卡洛计算德州
func CalRate(cal*httpModel.CalData)model.Result{
   final:=model.Result{}
   data:= ToDealData(cal)//处理好的对象
   win:=float32(0.0)
   for i:=0;i<data.Frequency;i++{
    result:=Start(data)//每一次的
     if result.Ev>0{
		win++
	 }
   }
    final.Ev=win/float32(data.Frequency)*100.0
	
   return final
}
//把数据简单处理了 该map的map
func ToDealData(cal*httpModel.CalData)*model.Begin{
	hand:=cal.Cards.Hand
	public:=cal.Cards.PublicCards
	Handmap:=Utils.DealMap(hand)
	PublicMap:=Utils.DealMap(public)
	begin:=&model.Begin{
	  Person: cal.Table.Person,
	  Id: cal.Cards.Position,
      Hand: Handmap,
      PublicCard: PublicMap,
	  Frequency: 10000,
	  Action:cal.Table.PlayerAction,
	}
	return begin
}