package service
import "poker/model"
func GameFactory (name string)model.Game{
  if name=="texas"{
	return &TexasGame{}
  }
  return nil
}