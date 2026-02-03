package httpModel
type CalData struct{
	Game string
	 Cards CardInfo `json:"cards"`
	 Table TableInfo `json:"table"`
}
type CardInfo struct{
   Hand []string `json:"hand"`
   Position int `json:"position"`  
   PublicCards[] string `json:"publicCards"`
}
type TableInfo struct{
	Person int `json:"person"`
	DeadMoney float32`json:"deadMoney"`
	CallMoney float32 `json:"callMoney"`
	PlayerAction []int `json:"action"`
}
func NewData()*CalData{
	data:=&CalData{}
	return data
}