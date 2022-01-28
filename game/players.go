package main

type Players []Player

type Player struct{
	id int
	room Room
	backpack Backpack
	haveBackpack bool
	quests Quests
}


func (p *Player) Go(nameRoom string) string {
	if ok := p.goToRoom(rooms.GetByName(nameRoom)); ok {
		return "перешел в: " + p.room.name
	}
	return "нет пути в " + nameRoom
}

func (p *Player) goToRoom(r Room) bool {
	if currentPlayer.room.CanIGoToRoom(r.name) {
		p.room = r
		return true
	}
	return false
}

func (p *Player) ClotheBackpack() string {
	p.haveBackpack = true
	return "вы одели: рюкзак"
}

func (p *Player) TakeItem(s string) string {
	if !currentPlayer.haveBackpack {
		return "некуда класть"
	}

	if i, ok := currentPlayer.room.GetItemByName(s); ok {
		currentPlayer.backpack = append(currentPlayer.backpack, i)
		return "предмет добавлен в инвентарь: " + i.name
	}
	return "нет такого"
}

func (p *Players) NewPlay(room Room) Player {
	pl := Player{id: len(*p) + 1, room: room}
	pl.quests.StartQuests()
	*p = append(*p, pl)
	return pl
}