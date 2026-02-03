package Utils
  var Global AllHand
  //读起始表
func RangeFile ()error{
    path:="/workspaces/-poker/RangeTable/sortvs1.jsonl"
	data,error:=ReadFile(path)
	if error!=nil{
		return  error
	}
    Global=data
    return nil
}
//全局手牌起始表
func RangeTable()AllHand{
   return Global
}