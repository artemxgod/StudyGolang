package main

type Rooms []Room

func (r Rooms) GetByName(name string) Room{
	for _, v := range r{
		if v.name == name{
			return v
		}
	}
	panic("Комнаты нет")
}