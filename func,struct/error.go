package main 

import "fmt"

func divide(a, b int) int{
	if b == 0 {
		panic("division by zero")
	}else{
		fmt.Print(divide(a,b))
	}
	return (a / b)
}

func main(){
	var a,b int
	fmt.Scan(&a, &b)
	fmt.Print(divide(a,b))
	
}
//defer <funcname>() - let finish function even if panic was called