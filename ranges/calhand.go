package ranges

import ("encoding/json"

)

func CalHand(hand []string)[]byte{
	data:=&Data{}
	data.Cards.Hand=hand
	data.Tables.Person=6
    data.Tables.Action=[]int{1}
    data.Cards.Public=nil
    data.Cards.Position=0
    jsonData,_:=json.Marshal(data)
	return jsonData
}