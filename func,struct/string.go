package main

import (
	"bufio"
	"fmt"
   	"os"
 	str "strings"
   "unicode"
   "unicode/utf8"
)

func IsPolindrom(s string) bool{
	Rune := []rune(s)
	amount := utf8.RuneCountInString(s)
	j := amount - 1
	for i := 0; i <= amount / 2 ; i++{
		fmt.Println(Rune[i],  Rune[j])
		if Rune[i] != Rune[j] {
			return false
		}
		j--
	}
	return true
}
func Odd_sym(s string) string{
	amount := utf8.RuneCountInString(s)
	Rune1 := []rune(s)
	Rune2 := make([]rune, amount/2)
	j := 0
	for i:=1; i<amount; i++ {
		if i % 2 == 1{
			Rune2[j] = Rune1[i]
			j++
		}
	}
	return string(Rune2)
}
func Del_rep(s string) string{
	amount := utf8.RuneCountInString(s)
	for i:= 0; i < amount; i++{
		cnt := str.Count(s, string(s[i]))
		if(cnt > 1){
			s = str.ReplaceAll(s,string(s[i]), "")
			amount -= cnt
		}
	}
	return s
}
func Password(s string) bool{
	Rune := []rune(s)
	amount := utf8.RuneCountInString(s)
	if amount < 5{
		return false
	}
	for i:= 0; i < amount; i++{
		if (unicode.IsDigit(Rune[i]) == false) && (unicode.Is(unicode.Latin, Rune[i]) == false){
			return false
		}
	}
	return true
}
func main(){
	//Task 1 - first Upper 
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	Rune := []rune(text)
	amount := utf8.RuneCountInString(text)
	if unicode.IsUpper(Rune[0]) && Rune[amount - 3] == '.'{
		fmt.Print("Right")
	}else{
		fmt.Print("Wrong")
	}

	//Task 2 - polindrom
	var text string
	fmt.Scan(&text)
	if IsPolindrom(text) == true {
		fmt.Print("Палиндром")
	}else{
		fmt.Print("Нет")
	}
	
	//Task 3 - Contains
	var str1,str2 string
	fmt.Scan(&str1, &str2)
	fmt.Print(str.Index(str1, str2))

	//Task 4 - odd only
	var s string 
	fmt.Scan(&s)
	fmt.Print(Odd_sym(s))

	//Task 5 - delete repetitive symbols
	var str string
	fmt.Scan(&str)
	fmt.Print(Del_rep(str))

	//Task 6 - is good password
	var s string

	fmt.Scan(&s)
	if Password(s) == true{
		fmt.Print("Ok")
	}else {
		fmt.Print("Wrong password")
	}

}

