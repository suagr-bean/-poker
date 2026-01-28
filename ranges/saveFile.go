package ranges

import (
	"encoding/json"
	"os"
	"path/filepath"
)
func SaveFile(dir string,filename string,data []string,win string ,jsondata string,)error{
	err:=os.MkdirAll(dir,0755)
    if err!=nil{
	 return err
	}
   path:=filepath.Join(dir,filename)
   open,err:=os.OpenFile(path,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
   if err!=nil{
	return err
   }
   defer open.Close()
   encoder:=json.NewEncoder(open)
   datajson:=map[string]any{

	 "jsondata":data,
    "win":win,
   }
   encoder.Encode(datajson)
   return nil
}