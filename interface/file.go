package main

import (
	"archive/zip"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

//task 1 find sum using bufio
func Sum() {
	var sum int = 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		strnum := scanner.Text()
		num, err := strconv.Atoi(strnum)
		if err != nil{
			panic("found \n in num")
		}
		sum += num
	}
	io.WriteString(os.Stdout, strconv.Itoa(sum))
}

//task 2 - find right file
func FindF(){
	z, _ := zip.OpenReader("task.zip")
	defer z.Close()
	for _, f := range z.File {
		r, _ := f.Open()
		if rows, _ := csv.NewReader(r).ReadAll(); len(rows) == 10 && len(rows[4]) == 10 {
			fmt.Println(f.Name, rows[4][2])
		}
		r.Close()
	}
}
//task 2 - other solution(using path/filepath)
func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
	   return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}
 
	if info.IsDir() {
	   return nil // Проигнорируем директории
	}
 
	file, errOpen := os.Open(path)
	if errOpen != nil {
	   panic(err)
	}
 
	rCSV := csv.NewReader(file)
 
	data, _ := rCSV.ReadAll()
 
	if len(data) == 10 {
	   fmt.Println("data[4][2]", data[4][2])
	}
	return nil
 }
 
 //task 3 - find zero num index
func Readf(){
	var index uint64 = 0

	file, err := os.Open("interface/text.data")
	if err != nil{
		panic("error")
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 379400)
	buf := make([]byte, 379400)
	n, err := reader.Read(buf)
	if n != 379400{
		panic("byte err")
	}
	str := strings.Split(string(buf), ";")
	for _, elem := range str{
		index++
		num, _ := strconv.Atoi(elem)
		if(num == 0){
			fmt.Print("index - ", index)
		}
	}
	
	
}

func NewReadf(){
	file, err := os.Open("interface/task.data")
   if err != nil {
      return
   }
   defer file.Close()
   r := bufio.NewReader(file)
   var count int
   for {
      str, err := r.ReadString(';')
      if err != nil {
         break
      }
      count++
      if str == "0;" {
         fmt.Print(count)
         break
      }
   }
}
func main(){
	if err := filepath.Walk(".", walkFunc); err != nil{
		fmt.Printf("ошибка : %v/n", err)
	}
	 Readf()
	NewReadf()

}