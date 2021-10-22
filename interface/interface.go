package main

import (
	"fmt"   
	"strings"                   
)

func Opresult(num1, num2 float64, op string ){
	switch op{
	case "+": fmt.Printf("%.4f", num1 + num2)
	case "-": fmt.Printf("%.4f", num1 - num2)
	case "*": fmt.Printf("%.4f", num1 * num2)
	case "/": fmt.Printf("%.4f", num1 / num2)
	}

}
func arif(num1, num2, operation interface{}){
	switch v := num1.(type){
	case float64:
		switch v1 := num2.(type){
		case float64: 
			switch v2:= operation.(type){	
			case string: 
				if strings.Contains("+-*/", v2){
				Opresult(v,v1,v2)
				} else{
					fmt.Print("Неизвестная операция")
				}
			default: 
				fmt.Print("Неизвестная операция")
			}
		default : 
			fmt.Print("value = ", v1)
			fmt.Printf(":%T", v1)
		}
	default :
	fmt.Print("value = ", v)
	fmt.Printf(":%T", v)	
}
}

type battery struct {
	str string
}

func (b battery) String() string{
	var tmp string
	cnt := 0
	for _, elem := range b.str{
		if string(elem) == "0"{
			tmp += " "
		}else{
			cnt++
		}
	}
	for cnt > 0{
		tmp += "X"
		cnt--
	}
	return fmt.Sprintf("[%v]", tmp)
}



func main(){
	//task 1 - check type
	c := "+"
	a,b := 85.5, 75.3
	arif(a,b,c)

	//task 2 - battery
	var batteryForTest = battery{}
	fmt.Scan(&batteryForTest.str)
	fmt.Print(batteryForTest.String())


}