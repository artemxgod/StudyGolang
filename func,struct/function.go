package main

import "fmt"

func main(){
	var n,a,b int

	 fmt.Scan(&n)
	 n = fibonacci(n)
	 fmt.Print(n)

	 a,b = sumInt(1,2,3)
	 fmt.Print(a,b)



}
//Task 1 
func fibonacci(n int) int {
	var first,second,fibnum int
	first = 1
	second = 1
	if n == 1{
		return 1
	}else if n == 2{
		return 1
	}else{
		for i := 3; i <= n; i++{
			fibnum = first + second
			first = second
			second = fibnum
		}
		return fibnum
	}
}
//Task2
//получение аргументов типа int любого кол-ва
	func sumInt(num ...int) (int, int){
	var sum,cnt int

	sum = 0
	cnt = 0
	for i:= 0; i<len(num); i++{
		sum += num[i]
		cnt++
	}
	return cnt, sum
}

//we are able to name returnable variable
func Foo1(a int) (res int){
	res+= a
	return res
}




