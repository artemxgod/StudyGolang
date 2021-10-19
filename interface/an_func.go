package main

import (
	"fmt"
	"strconv"
	"errors"
)

//Task 1 - delete all odd digits and zero from number using fn, declared in main()
func main(){
	fn := func(num uint) uint{
		var ten, newnum uint64
		Rune := []rune(strconv.FormatUint(uint64(num), 10))
		ten,newnum = 1,0
		for i:= (len(Rune) - 1); i >= 0; i--{
			a, _ := strconv.ParseUint(string(Rune[i]), 10, 64)
			if a % 2 == 0 && a != 0{
				newnum += ten*a
				ten*= 10
			}
		}
		if newnum == 0{
			newnum = 100
		}
		return uint(newnum)
	}
	a := fn(110)
	fmt.Print(a)
}
//solution without strconv
func fn(v uint) uint{
    var r, d, i uint
    for i=1 ; v>0; v/=10{
        if d= v%10; d&1==0 {
            r+= d*i
            i*=10
        }
    }
    if r==0 { r=100 }
    return r
}

//defer am func
//defer выполняется только после выполнения внешней функции, порядок выполнения - стек
func someFuncWithPanic() (err error) {
	defer func() {
		// отложенный вызов анонимной функции, проверяющей, что работа функции завершена
		// без ошибок. Если функция recover() возвращает что угодно кроме nil, значит в ходе
		// выполнения функции возникла паника.
		if e := recover(); e != nil {
			// Здесь происходит приведение интерфейса (об этом мы расскажем буквально в
			// следующем уроке. Результат приведения присваивается переменной err типа error
			// которая уже объявлена при самом вызове функции someFuncWithPanic.
			err = e.(error)

			// после этого анонимная функция завершает свою работу, паника обработана,
			// переменная err, в которой содержится информации о возникшей панике,
			// возвращается как результат выполнения функции.
		}
	}()

	panic(errors.New("fatal error"))
}

func ExamplePanicRecover() {
	if err := someFuncWithPanic(); err != nil {
		fmt.Println(err)
	}

	// Output:
	// fatal error
}