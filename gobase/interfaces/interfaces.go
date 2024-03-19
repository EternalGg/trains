package interfaces

import "fmt"

// 实现暴乱狂战士 monitor
type Hero struct {
	Health      uint
	Name        string
	AttackPoint uint
	below       uint
}

func (h *Hero) BeDamage() {
	h.Health--

}

type Crazy struct {
	listener  *Hero
	listeners []*Hero
}

func (c *Crazy) Subscribe(h *Hero) {
	c.listeners = append(c.listeners, h)
}

func (e *Crazy) CrazySkill() {
	for i := 0; i < len(e.listeners); i++ {
		e.listener.AttackPoint++
	}
}

var CrazyList = []*Crazy{}

func Mains() {
	Hero1 := Hero{Health: 1, Name: "hero1"}
	Hero2 := Hero{Health: 1, Name: "hero2"}
	carzy := Hero{Health: 1, Name: "Carzy", AttackPoint: 0}
	c := Crazy{}
	Crazylist := CrazyList
	Crazylist = append(Crazylist, &c)
	fmt.Println(Crazylist)
	c.listener = &carzy
	c.Subscribe(&Hero1)
	c.Subscribe(&Hero2)

	fmt.Println(Hero2.Name, Hero2.Health, carzy.Name, carzy.AttackPoint)
	Hero2.BeDamage()
	if len(Crazylist) > 0 {
		for _, C := range Crazylist {
			C.CrazySkill()
		}
	}
	fmt.Println(Hero2.Name, Hero2.Health, carzy.Name, carzy.AttackPoint)
	fmt.Println(Hero1.Name, Hero1.Health, carzy.Name, carzy.AttackPoint)
	Hero1.BeDamage()
	fmt.Println(Hero1.Name, Hero1.Health, carzy.Name, carzy.AttackPoint)
	if len(Crazylist) > 0 {
		for _, C := range Crazylist {
			C.CrazySkill()
		}
	}
}
