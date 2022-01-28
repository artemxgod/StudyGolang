package main

type Backpack []Item

func (b Backpack) ItemsToString() string{
	var res string
	for _, v := range b{
		res += v.name
	}
	res = DefaultForEmplyString(res, "Ничего нет")
	return res
}
