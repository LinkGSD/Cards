package main

import (
	"fmt"
	"math/rand"
)

var Site [2][3]*Card
var OSite [2][3]*Card

type ET int

const (
	SameColumn ET = iota
	SameRow
	Random
	Self
	All
)

type EM int

const (
	EMRuChang    EM = iota //入场
	EMXianShou             //先手
	EMYiYan                //遗言
	EMChengZhang           //成长
	EMAttach               //附着
	EMAttack               //进攻
)

type Effect struct {
	Harm       bool
	Friendly   bool
	EffType    ET
	EffNum     int
	EffMethod  EM
	HarmNum    int
	AttackIncr int
	HPIncr     int
}

func (e Effect) Trigger(c *Card) {
	cards := e.getEffectCards(c)
	for _, card := range cards {
		if e.Harm {
			card.HP -= e.HarmNum
		} else {
			card.Attack += e.AttackIncr
			card.HP += e.HPIncr
		}
	}
}

func (e Effect) getEffectCards(c *Card) []*Card {
	switch e.EffType {
	case SameColumn:
		return []*Card{Site[1-c.Position.X][c.Position.Y]}
	case SameRow:
		return Site[c.Position.X][:]
	case Random:
		cards := make([]*Card, 0)
		for i := 0; i < e.EffNum; i++ {
			if e.Friendly {
				cards = append(cards, Site[rand.Intn(2)][rand.Intn(3)])
			} else {
				cards = append(cards, OSite[rand.Intn(2)][rand.Intn(3)])
			}
		}
		return cards
	case Self:
		return []*Card{Site[c.Position.X][c.Position.Y]}
	case All:
		if e.Friendly {
			return append(Site[0][:], Site[1][:]...)
		} else {
			return append(OSite[0][:], OSite[1][:]...)
		}
	}
	return []*Card{}
}

type Race int

const (
	Beast Race = iota
	Mechanical
	Ghost
)

type Card struct {
	Name       string
	Attack     int
	HP         int
	Race       Race
	Attached   map[EM][]Effect
	Damage     int
	Position   Position
	Effect     *Effect
	Attach     bool
	FullAttach bool
}

func (c *Card) Trigger(em EM) {
	if c.Effect != nil && c.Effect.EffMethod == em {
		c.Effect.Trigger(c)
	}
	if c.Attached != nil {
		for _, e := range c.Attached[em] {
			e.Trigger(c)
		}
	}
}

func (c *Card) Attacks() {
	if c.Effect != nil && c.Effect.EffMethod == EMAttack {
		c.Effect.Trigger(c)
	}
}

type Position struct {
	X int
	Y int
}

func main() {
	c := &Card{Name: "卡1", Attack: 10, HP: 10, Race: Beast, Position: Position{X: 0, Y: 0}, Effect: &Effect{Friendly: true, EffType: Random, EffNum: 2, AttackIncr: 3, HPIncr: 1, EffMethod: EMChengZhang}}
	Site[0][0] = c
	Site[0][1] = &Card{Name: "卡2", Attack: 11, HP: 10, Race: Beast, Position: Position{X: 0, Y: 1}, Effect: &Effect{Friendly: true, EffType: SameColumn, AttackIncr: 10, HPIncr: 10, EffMethod: EMXianShou}}
	c2 := &Card{Name: "卡3", Attack: 12, HP: 10, Race: Beast, Position: Position{X: 0, Y: 2}, Effect: &Effect{Friendly: true, EffType: Self, AttackIncr: 7, HPIncr: 7, EffMethod: EMAttack}}
	Site[0][2] = c2
	Site[1][0] = &Card{Name: "卡4", Attack: 13, HP: 10, Race: Beast, Position: Position{X: 1, Y: 0}}
	Site[1][1] = &Card{Name: "卡5", Attack: 14, HP: 10, Race: Beast, Position: Position{X: 1, Y: 1}}
	Site[1][2] = &Card{Name: "卡6", Attack: 15, HP: 10, Race: Beast, Position: Position{X: 1, Y: 2}}
	//准备阶段
	fmt.Println("准备阶段结束")
	// 初始化
	for i := 0; i < len(Site); i++ {
		for j := 0; j < len(Site[0]); j++ {
			Site[i][j].Trigger(EMChengZhang)
		}
	} //
	for i := 0; i < len(Site); i++ {
		for j := 0; j < len(Site[0]); j++ {
			fmt.Println(Site[i][j].Name, Site[i][j].Attack, Site[i][j].HP)
		}
	} //
	fmt.Println("先手阶段")
	for i := 0; i < len(Site); i++ {
		for j := 0; j < len(Site[0]); j++ {
			Site[i][j].Trigger(EMXianShou)
		}
	} //
	c2.Attacks()
	for i := 0; i < len(Site); i++ {
		for j := 0; j < len(Site[0]); j++ {
			fmt.Println(Site[i][j].Name, Site[i][j].Attack, Site[i][j].HP)
		}
	} //

}
