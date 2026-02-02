package model
type Result struct {
	Lose bool
	LoseScore []int
	Ev  float32 
	
}
type CalResult struct{
	Win float32
	LoseInfo CardType
}