package main

import (
	"bufio"
	"io"
	"os"
	"time"
	"fmt"
	"strings"
)

//task 1 - print time and date in UnixDate format
func Print_time(){
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	str := scan.Text()
	data_time, err := time.Parse(time.RFC3339,str) //p1 - format of date and time, p2 - string to parse
	if err != nil{
		panic(err)
	}

	io.WriteString(os.Stdout, data_time.Format(time.UnixDate)) //p - format of date and time
}
//task 2 - special print time
func Spec_print(){
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	str := scan.Text()
	time, err := time.Parse("2006-01-02  15:04:05",str)
	if err != nil{	panic(err)	}

	if time.Hour() >= 13{
		time = time.AddDate(0,0,1) // day + 1
	}
	fmt.Print(time.Format("2006/01/02 15:04:05"))
}
//task 3 - duration between two dates
func time_Dur(){
	var a time.Duration
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	s := scan.Text()
	sArr := strings.Split(s, ",")

	first,err := time.Parse("02.01.2006  15:04:05", sArr[0])
	if err != nil{ panic(err) }
	second,err := time.Parse("02.01.2006  15:04:05", sArr[1])
	if err != nil{ panic(err) }

	if first.After(second){
		a = second.Sub(first) 
	}else{
		a = first.Sub(second)
	}
	fmt.Print(a)
}

//task 4 - now + duration
func New_time(){
	var hour,min,sec int
	cur := time.Unix(1589570165, 0)
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	duration := "0h"
	duration += scan.Text()
	duration = strings.ReplaceAll(duration, " мин. ", "m")
	duration = strings.ReplaceAll(duration, " сек.", "s")

	dur, err := time.ParseDuration(duration)
	if err != nil{
		panic(err)
	}
	hour = int(dur.Round(time.Hour).Hours())
	min = int(dur.Round(time.Minute).Minutes()) - hour*60
	sec = int(dur.Round(time.Second).Seconds()) - hour*3600 - min*60
	if sec > 30{
		min -= 1
	}
	cur = cur.UTC()
	cur = cur.Add(time.Hour * time.Duration(hour) + 
	time.Minute * time.Duration(min) + 
	time.Second *  time.Duration(sec))

	fmt.Print(cur.Format(time.UnixDate))
}

//task 4 - other solution
func New_Dur(){
	var m, s time.Duration
	fmt.Scanf("%d мин. %d сек.", &m, &s)
	t := time.Unix(1589570165, 0).UTC().Add(m * time.Minute).Add(s * time.Second)
	fmt.Println(t.Format(time.UnixDate))
}
func main(){
	New_Dur()
}