package main

import "fmt"

func main(){
	// Task 1 "sign of the number"
	/*
	var a int

	fmt.Scan(&a)
	if a > 0 {
		fmt.Print("Число положительное")
	}else if a < 0 {
		fmt.Print("Число отрицательное")
	}else {
		fmt.Print("Ноль")
	}*/

//---------------------------------------\\

	// Task 2 "different digits of number"
	/*var num,one,ten,hundred int

	fmt.Scan(&num)
	one = num % 10
	ten = num / 10 % 10
	hundred = num / 100
	if one != ten && ten != hundred && one != hundred {
		fmt.Print("YES")
	}else {
		fmt.Print("NO")
	}*/

//---------------------------------------\\
	
	//Task 3 "first digit of the number"
	/*var num uint

	fmt.Scan(&num)
	switch {
	case num < 10:
		fmt.Print(num)
	case num < 100:
		fmt.Print(num / 10)
	case num < 1000:
		fmt.Print(num / 100)
	case num < 10000:
		fmt.Print(num / 1000)
	case num == 10000:
		fmt.Print(num /10000)
	}*/
	
//---------------------------------------\\

	//Task4 "Lucky ticket"
	/*var( num, one, one1, ten, ten1, hundred, hundred1 int
		sum int = 0
		sum1 int = 0
	)

	fmt.Scan(&num)
	one = num % 10
	ten = num / 10 % 10
	hundred = num / 100 % 10
	sum = one + ten + hundred
	one1 = num / 1000 % 10
	ten1 = num / 10000 % 10
	hundred1 = num / 100000
	sum1 = one1 + ten1 + hundred1
	if sum == sum1 {
		fmt.Print("YES")
	}else {
		fmt.Print("NO")
	}*/
	
//---------------------------------------\\

	//Task 4 "leap year"
	/*var year uint

	fmt.Scan(&year)
	if year % 400 == 0{
		fmt.Print("YES")
	}else if year % 4 == 0 && year % 100 != 0{
		fmt.Print("YES")
	}else {
		fmt.Print("NO")
	}*/
	fmt.Print("Done")
}