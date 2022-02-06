package main

import . "fmt" // now we can use functions from "fmt" without "fmt."
import mat "math" // instead of math. now we will use mat. 
// we can import more than one package in one line using () Expl: import ("a"; "b")  
func main(){
	Println("Hi!") //works 
	Print(mat.Abs(-3)) //works
	
}