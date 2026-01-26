package service
import "poker/model"
func SetPlayer(player*model.GameData,action[]int){
   for _,v:=range action{
	 player.SetAction(v,1)
   }
   
}