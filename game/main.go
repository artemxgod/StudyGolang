package main

import (
	"fmt"
)

var players Players 
var currentPlayer Player
var rooms Rooms

func InitGame(){
	rooms = Rooms{
		{id: 1, 
			name: "кухня", 
			associated: map[int]string{2: "коридор"}, 
			def: true,
		roomPlace: []RoomPlace{
			{"на столе", []Item{{"чай"}}},
		},
	},
	{id: 2, name: "коридор", associated: map[int]string{1: "кухня", 3: "комната"}},
	{id: 3, name: "комната", associated: map[int]string{2: "коридор"},
		roomPlace: []RoomPlace{
			{"на столе:", []Item{{"ключи"}, {"конспекты"}}},
			{"на стуле -", []Item{{"рюкзак"}}},
		},
	},
}
currentPlayer = players.NewPlay(rooms.GetByName("кухня"))
}


func main(){
	InitGame()

	fmt.Println("У меня есть рюкзак?", YesNo(currentPlayer.haveBackpack))
	fmt.Println("положить ключ?", currentPlayer.TakeItem("ключ"))
	fmt.Println("Что в рюкзаке?", currentPlayer.backpack.ItemsToString())
	fmt.Println("Я одел рюкзак?", currentPlayer.ClotheBackpack())
	fmt.Println("Что в рюкзаке?", currentPlayer.backpack.ItemsToString())
	fmt.Println("У меня есть рюкзак?", YesNo(currentPlayer.haveBackpack))
	fmt.Println(currentPlayer.room.Lookup())
	fmt.Println(currentPlayer.Go("комната"))
	fmt.Println(currentPlayer.room.Lookup())
	fmt.Println(currentPlayer.Go("коридор"))
	fmt.Println(currentPlayer.room.Lookup())
	fmt.Println(currentPlayer.Go("комната"))
	fmt.Println("--------------")
	fmt.Println(currentPlayer.room.Lookup())
	fmt.Println("положить ключи?", currentPlayer.TakeItem("ключи"))
	fmt.Println("Что в рюкзаке?", currentPlayer.backpack.ItemsToString())
	fmt.Println(currentPlayer.room.Lookup())
}

