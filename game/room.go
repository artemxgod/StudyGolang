package main

import "strings"

type Room struct{
	id int
	def bool
	name string
	roomPlace []RoomPlace
	associated map[int]string
}

func (r Room) Lookup()string{
	if r.def {
		return "ты находишься на кухне, " + r.ItemToString() + ", надо " +
		currentPlayer.quests.ToString() +
		 ". можно пройти - " + r.AssociatedToString()
	}
	if len(r.ItemToString()) > 0{
		return r.ItemToString() + ". можно пройти - " + r.AssociatedToString()
	}
	return "пустая" + r.name + ". можно пройти - " + r.AssociatedToString()
}

func (r Room) AssociatedToString() string{
	var res string
	roomsArray :=make([]string, 0, 2)
	for _, v := range r.associated{
		roomsArray = append(roomsArray, v)
	}
	res = strings.Join(roomsArray, ", ")
	return res
}

func (r Room) CanIGoToRoom(s string) bool{
	room := rooms.GetByName(s)
	_, ok := r.associated[room.id]
	return ok
}

func (r Room) ItemToString() string{
	var res string
	for _, v := range r.roomPlace{
		res += v.name + " " + v.ItemToString()
	}
	res = DefaultForEmplyString(res, "ничего")
	return res
}

func (r Room) GetItemByName(s string) (Item, bool){
	for _, v := range r.roomPlace {
		if item, ok := v.HaveItem(s); ok{
			return item, true
		}
	}
	return Item{}, false
}