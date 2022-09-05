package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


// PrintByID print student by id
func PrintByID(id int64) {
	var fio string
	var info sql.NullString
	// var info string
	var score int
	row := db.QueryRow("SELECT fio, info, score FROM students WHERE id = ?", id)
	// fmt.Println(row)
	err := row.Scan(&fio, &info, &score)
	PanicOnErr(err)
	fmt.Println("PrintByID:", id, "fio:", fio, "info:", info, "score:", score)
}

var db *sql.DB

func main() {
	var err error

	// создаём структуру базы
	// но соединение происходит только при мервом запросе
	db, err = sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/mysql_th1")
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10) // max connections at the time

	// проверяем что подключение реально произошло ( делаем запрос )
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Connected to db succesfully")

	// READ DATA

	// итерируемся по многим записям
	// Exec исполняет запрос и возвращает записи
	// Query() let us get all the results
	rows, err := db.Query("SELECT fio, score FROM students")
	PanicOnErr(err)
	for rows.Next() {
		var fio, score string
		err = rows.Scan(&fio, &score)
		PanicOnErr(err)
		fmt.Println("row.Next fio: ", fio, score)
	}
	//always close connection
	rows.Close()

	// QueryRow() let us get only one result
	var fio string
	row := db.QueryRow("SELECT fio FROM students WHERE id = 1")
	err = row.Scan(&fio)
	PanicOnErr(err)
	fmt.Println("db.QueryRow fio:", fio)

	// WRITE down new data
	// Exec исполняет запрос и возвращает сколько строк было затронуто и последнйи ИД вставленной записи
	// символ ? является placeholder-ом. все последующие значения авто-экранируются и подставляются с правильным кавычками
	result, err := db.Exec(
		"INSERT INTO students (`fio`, `score`) VALUES (?, ?)",
		"Artur", 10,
	)
	PanicOnErr(err)

	//result returns this:
	affected, err := result.RowsAffected()
	PanicOnErr(err)
	lastID, err := result.LastInsertId()
	PanicOnErr(err)

	fmt.Println("Insert - Affected:", affected, "LastId:", lastID)

	PrintByID(lastID)

	// UPDATE DATA
	// использование prepared statements
	// Prepare подготовливает запись ( шлёт запрос на сервере, там он парсится и готов принимать данные)
	stmt, err := db.Prepare("UPDATE students SET info = ?, score = ? WHERE id = ?")
	PanicOnErr(err)

	// Exec для prepares statement отправляет даныне на сервер - там запрос уже распашрен, только исполняется с новыми данными
	_, err = stmt.Exec("prepared stmt update", 5, lastID)
	PanicOnErr(err)
	result, err = stmt.Exec("8 update", lastID, 5)
	PanicOnErr(err)

	affected, err = result.RowsAffected()
	PanicOnErr(err)
	fmt.Println("Update - RowsAffected:", affected)
	PrintByID(lastID)
	// // set data
	// insert, err := db.Query("INSERT INTO `students`  (`fio`, `info`, `score`) VALUES ('Vasily Romanov', 'company: Mail.ru Group', '10')")
	// if err != nil {
	// 	panic(err)
	// }

	// defer insert.Close()	

}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
