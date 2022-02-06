package main

import "fmt"

func main(){
	var a,b int
	fmt.Scan(&a,&b)
	test(&a,&b)
	test1(&a,&b)
}
//task 1 - multiply
func test(x1 *int, x2 *int) {
	a := *x1 * *x2
	fmt.Println(a)
}
//task 2 - change
func test1(num1, num2 *int){
	*num1,*num2 = *num2,*num1
	fmt.Println(*num1, *num2)
}