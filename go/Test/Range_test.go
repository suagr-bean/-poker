package service

import (
	"poker/Utils"
	"testing"
)
//测试给范围拿到牌堆
func TestRange(t *testing.T){
	chance:=10
   rangehand:=Utils.PlayerRange(chance)
   for i:=0;i<len(rangehand.HandIndex);i++{
	t.Log("\n","数据",rangehand.HandIndex[i].Hand)
   }
   t.Log("多少个id",len(rangehand.HandIndex))
}