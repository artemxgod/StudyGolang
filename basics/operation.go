package main

import "fmt"

func main(){
	var hours,minutes,a uint
	fmt.Scan(&a)
	hours = a / 30
	minutes = a % 30
	fmt.Print("it is ", hours, " hours ", minutes, " minutes ")
}