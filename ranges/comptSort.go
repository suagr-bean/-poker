package ranges
func  Compt (){
	path:="../../Range/handvs1.jsonl"
	dir:="../../Range"
	file:="sortVs1.jsonl"
	data,_:=ReadFile(path)
	  sortdata:=SortFile(data)
	for _,v:=range sortdata.HandIndex{
       hand:=v.Hand
	   win:=v.Win
	   SaveFile(dir,file,hand,win)
	}
   
}