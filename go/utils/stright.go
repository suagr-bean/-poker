package utils

import "math/bits"
func DealStright(mask uint16)(bool,int){
	move:=mask&(mask>>1)&(mask>>2)&(mask>>3)&(mask>>4)
	if move!=0{
      return true,bits.Len16(move)+3
	}
	var small uint16=0x403c
   if (move&small)==small{
	return true,5
   }
   return false,0
}