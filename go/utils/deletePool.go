package utils
import "fmt"
func(T*TexasDealer) DeletePool()[]int{
    pool:=make([]int,0)
	cards:=T.Board.GetBoardCards()
	player:=T.Data.GetAllPlayer()
	pool=append(pool,cards...)
	for _,v:=range player{
		pool=append(pool,v.Hand...)
	}
	fmt.Println(pool)
	return pool
}