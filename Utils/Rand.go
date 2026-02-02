package Utils

import (
	"fmt"
	"math/rand/v2"
)
func Rand(chance int,cards[]int)([]int,error){
	if len(cards)==0{
		error:=fmt.Errorf("cards不合法,要删除cards不能为0")
		return nil,error 
	}
	allhand:=RangeTable()//拿到所有牌的可能性
	RangeIndex:=int(len(allhand.HandIndex)*chance/100) //除100原因是前端给的是Int 20就是20%
     if RangeIndex<=0{
		error:=fmt.Errorf("RangeIndex 不合法没有范围")
		return nil,error 
	 }
	 //拒绝采样函数 尝试50次
	hand:=Simpling(RangeIndex,cards,50)
	if hand==nil{
		//50次没成功 玩家范围很窄
		//回退 过滤
	 hand:=Backup(RangeIndex,cards)
	  if hand==nil{
        error:=fmt.Errorf("牌堆抽空")
		return nil,error
	  }
	  return hand,nil
	}
	

    return hand,nil
}
//拒绝采样
func Simpling(index int,cards[]int,maxtry int)[]int{
	allhand:=RangeTable()//所有牌
	
	if maxtry<=0{
		return  nil
	}
	for i:=0;i<maxtry;i++{
		randIndex:=rand.IntN(index)
		hand:=allhand.HandIndex[randIndex].MapHand//这里不能用hand hand是原始的string maphand才是能用的
	check:=false
	//和牌堆检查发现牌堆有重新生成
	for _,v:=range cards{
     if hand[0]==v||hand[1]==v{
		check=true
		break
	 }
	}
	 if !check{//不冲突return
		return hand
	 }
	}
	return nil //尝试最大次数还是没成 说明范围太窄
}
func Backup(index int,cards []int)[]int{
   filter:=[][]int{}
   allhand:=RangeTable()

   for i:=0;i<index;i++{
	hand:=allhand.HandIndex[i].MapHand
	check :=false
	for _,v:=range cards{
		if hand[0]==v||hand[1]==v{
			check=true
			break
		}
	}
	if !check{//不冲突的可能
     filter=append(filter,hand)
	}
   }
   if len(filter)==0{//牌堆被抽空了
	 return nil
   }
   indexRand:=rand.IntN(len(filter))
   return filter[indexRand]
}