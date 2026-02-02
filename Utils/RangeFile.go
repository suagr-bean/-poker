package Utils
  var Global AllHand
  //读起始表
func RangeFile (){
    path:="/workspaces/-poker/poker/go/RangeTable/sortvs1.jsonl"
	data,error:=ReadFile(path)
	if error!=nil{
		return 
	}
    Global=data
}
//全局手牌起始表
func RangeTable()AllHand{
   return Global
}