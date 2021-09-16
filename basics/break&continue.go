package main

import "fmt"

func main(){
	//Task 1 - multiple to c, not multiple to d
	var n,c,d int

	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)
	for i := 0; i<=n; i++{
		if(i % c == 0 && i % d != 0){
		fmt.Print(i)
		break
		}
	}
	
	//---------------------------------------\\

	// Task 2 - output number that is between 10 and 100
	var num int

	for fmt.Scan(&num); num <= 100; fmt.Scan(&num){
		if(num < 10){
			continue
		}else {
			fmt.Println(num)
		}
	}
//---------------------------------------\\

//	Task 3 - bank deposit
	var x,p,y,cnt int

	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)
	cnt = 0
	for x < y{
		x += x * p / 100
		cnt++
	}
	fmt.Println(cnt)
	//---------------------------------------\\

	// Task 4 - digits that are similar to both numbers
	
	var number1, number2, reversnum1, copynum2, tmp1, tmp2 int
	
	fmt.Scan(&number1, &number2)
	reversnum1 = 0
	for number1 > 0{
		reversnum1 = reversnum1*10 + number1 % 10
		number1 /= 10
	}
	for reversnum1 > 0{
		tmp1 = reversnum1 % 10
		copynum2 = number2
		for copynum2 > 0{
			tmp2 = copynum2 % 10
			if tmp1 == tmp2{
				fmt.Print(tmp1, " ")
			}
			copynum2 /= 10
		}
		reversnum1 /= 10
	}
}
