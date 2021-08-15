package main

import "fmt"

func main(){
	const (
		_ = iota   // pass 0 
		monday 
		tuesday
		wednesday
		thursday
		friday
		saturday 
		sunday 
		ten = iota + 2
	)
	fmt.Print("Monday - ", ten)
	// different 'const' units
	const y = iota // y == 0
	const x = iota + 2 // x == 2
}