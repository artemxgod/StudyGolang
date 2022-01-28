package main

import "strings"

type RoomPlace struct{
	name string
	items []Item
}

func (rp RoomPlace) HaveItem(s string)(Item, bool){
	for _, v := range rp.items{
		if v.name == s{
			return v, true
		}
	}
	return Item{}, false
}

func (rp RoomPlace) ItemToString() string{
	ss:= make([]string, 0, len(rp.items))
		for _, v := range rp.items{
			ss = append(ss, v.name)
		}
		return strings.Join(ss, ", ")
	}