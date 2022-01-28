package main

import "strings"

type Quests []Quest

type Quest struct{
	isDone bool
	name string
}

func (q *Quests) StartQuests(){
	*q = Quests{
		{name: "собрать рюкзак"},
		{name: "идти в универ"},
	}
}

func (q Quests) ToString() string{
	var res string 
	quests := make([]string, 0, 2)
	for _, v := range q{
		if v.isDone == false{
			quests = append(quests, v.name)
		}
	}
	res = strings.Join(quests, " и ")
	res = DefaultForEmplyString(res, "заданий нет")

	return res
}