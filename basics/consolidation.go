package main

import "fmt"

func main() {

	//Task 1 - sum of digits of the number
	var num, sum int

	fmt.Scan(&num)

	if num < 99 || num > 999{
		fmt.Print("wrong number")
	}else{
		sum += (num % 10) + (num / 10 % 10) + (num / 100)
		fmt.Print(sum)
	}

	//Task 2 - reverse number output
	var num, digit int

	fmt.Scan(&num)
	digit = num % 10
	if num < 99 || num > 999 && digit == 0{
			fmt.Print("wrong number")
		}else{
			num = num / 100 + (num / 10 % 10) * 10 + digit * 100
			fmt.Print(num)
		}

	//Task 3 - second to hours and minutes
	var k, hours, minutes int

	fmt.Scan(&k)

	if k > 0 && k < 86399{
		hours = k / 3600
		minutes = (k - hours*3600) / 60
		fmt.Print("It is ", hours, " hours ", minutes," minutes.")
	}

	//Task 4 - is rectangular
	var a, b, c int

	fmt.Scan(&a, &b, &c)
	if (a*a + b*b) == c*c{
		fmt.Print("Прямоугольный")
	}else{
		fmt.Print("Непрямоугольный")
	}

	Task 5 - is triangle
	var a,b,c int

	fmt.Scan(&a, &b, &c)

	if a+b > c && a+c > b && b+c > a{
		fmt.Print("Существует")
	}else {
		fmt.Print("Не существует")
	}

	//Task 6 - av ar
	var a, b float32

	fmt.Scan(&a, &b)

	b = (a+b)/2
	fmt.Print(b)

	Task 7 - amount of zero
	var n, num, cnt int

	cnt = 0
	fmt.Scan(&n)
	for i := 0; i < n; i++{
		fmt.Scan(&num)
		if num == 0{
			cnt++
		}
	}
	fmt.Print(cnt)

	//Task 8 - amount of min
		var n, num, cnt, min int

		min = 1000
		cnt = 0
		fmt.Scan(&n)
		for i := 0; i < n; i++{
			fmt.Scan(&num)
			if num < min {
				min = num
				cnt = 1
			}else if min == num{
				cnt++
			}
		}
		fmt.Print(cnt)

	//Task 9 - num root. you can easily solve this task using formule "root = 1 + (num-1) mod 9"
	var (
		workArray [20]int
		num, tmp  int
	)

	fmt.Scan(&num)
	tmp = num
	if num >= 10 {
		for i := 0; tmp > 0; i++ {
			workArray[i] = tmp % 10
			tmp /= 10
		}
		for workArray[1] != 0 {
			tmp = 0
			for i := 0; i < 20; i++ {
				tmp += workArray[i]
			}
			for i := 0; i < 20; i++ {
				workArray[i] = 0
			}
			for i := 0; tmp > 0; i++ {
				workArray[i] = tmp % 10
				tmp /= 10
			}
		}
	} else {
		fmt.Print(num)
	}
	fmt.Print(workArray[0])

	//Task 10 - max num multiple to 7
	var a,b,max int
	
	fmt.Scan(&a, &b)
	max = -1000
	for i := a; i <= b; i++{
		if i % 7 == 0 && i > max{
			max = i
		}
	}
	if max != -1000{
		fmt.Print(max)
	}else {
		fmt.Print("NO")
	}

	//Task 11 - word ending

	var num int 
	
	fmt.Scan(&num)
	switch {
	case num == 1: fmt.Print("1 korova")
	case num > 1 && num < 5: fmt.Print(num, " korovy")
	case num >= 5 && num < 20: fmt.Print(num, " korov")
	case num > 20 && num < 100 && num % 10 == 1: fmt.Print(num, " korova")
	case num > 20 && num < 100 && num % 10 < 5: fmt.Print(num, " korovy")
	case num > 20 && num < 100 : fmt.Print(num, " korov")
	default: fmt.Print("wrong number")
	}

	//Task 12 - 2^ < num

	var num, bin int
	
	fmt.Scan(&num)
	bin = 1
	for bin <= num{
		fmt.Print(bin, " ")
		bin *= 2
	}

	//Task 13 - Fib num
	var (tmp, fnum1, fnum2, num, cnt int
	 check bool)

	fmt.Scan(&num)
	check = false
	fnum1 = 1
	fnum2 = 1
	cnt = 2
	tmp = 1 
	if num == 1 {
		fmt.Print(1)
	}else{
		for	tmp <= num {
			tmp = fnum1 + fnum2
			fnum1 = fnum2
			fnum2 = tmp
			cnt++
			if tmp == num{
				check = true
				fmt.Print(cnt)
				break
			}
		}
	}
	if check == false{
		fmt.Print(-1)
	}
	
	//Task 11 - delete digit from number

	var (
		digit, num int
		workArray [20]int 
	)
	fmt.Scan(&num)
	fmt.Scan(&digit)

	for i:= 0; i < 20; i++{
		workArray[i] = -1
	}
	for i:= 0; num > 0; i++{
		workArray[i] = num % 10
		num /= 10
	}
	for i:= 0; i < 20; i++{
		if workArray[i] == digit{
			workArray[i] = -1
		}
	}
	for i := 19; i >= 0; i--{
		if workArray[i] != -1{
		fmt.Print(workArray[i])
		}
	}

}
