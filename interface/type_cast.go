package main

import (
	"fmt"
	"strconv"
	"unicode"
	"strings"
	"bufio"
	"os"
)

func adding(s1, s2 string) int64{

	Rune1 := []rune(s1)
	Rune2 := []rune(s2)
	s1 = strings.Trim(s1, s1)
	s2 = strings.Trim(s2,s2)

	for _, sym := range Rune1{
		if unicode.IsDigit(sym){
			s1 += string(sym)
		}
	}
	for _,sym := range Rune2{
		if unicode.IsDigit(sym){
			s2 += string(sym)
		}
	}
	a, err := strconv.ParseInt(s1, 10, 10)
	b, err1 := strconv.ParseInt(s2, 10, 10)
	if err != nil || err1 != nil{
		panic("strings not cleared")
	}
	return  a + b
}
//alternative
func adding1(s1 string, s2 string) int64 { 

    parseNumber := func(s string) (r int64) {
        for _, c := range s {
            if unicode.IsDigit(c) {
                r = r*10 + int64(c) - '0'
            }
        }
        return
    }
    return parseNumber(s1) + parseNumber(s2)
}

func CSVdivision(s string) float64{
	var a,b float64
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, ",", ".", -1)
	fmt.Sscanf(s, "%f;%f", &a, &b )
	return a/b


}
func main(){
	//task1 - strings with trash into number's sum
	var s, s1 string
	fmt.Scan(&s, &s1)
	fmt.Print(adding(s,s1))

	//task2 - CSV division
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n")
	fmt.Printf("%.4f", CSVdivision(str))
}