package main 

import "fmt"

type Tank struct{
	On bool
	Ammo, Power int
}
func (obj *Tank) Shoot() bool {
	if obj.On == false {
		return false
	}
	if obj.Ammo > 0{
		obj.Ammo--
		return true
	}else {
		return false
	}
}
func (obj *Tank) Ride() bool {
	if obj.On == false {
		return false
	}
	if obj.Power > 0{
		obj.Power--
		return true
	}else {
		return false
	}
}

func main(){
	tiger := Tank{true, 9,3}
	T34 := Tank {true, 1, 0}
	if tiger.Shoot() == true{
		fmt.Println("BOOM!")
	}
	if T34.Shoot() == true{
		fmt.Println("BAM!")
	}
	if tiger.Ride() == true{
		fmt.Println("VJJJ")
	}
	if T34.Ride() == true{
		fmt.Println("VJU!")
	}
}

