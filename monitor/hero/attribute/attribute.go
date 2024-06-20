package attribute

import (
	"train/monitor/hero"
	"train/monitor/monitorfile"
)

type Attribute struct {
	Hero    *hero.Hero
	AttrMap map[int]int
}

func (a *Attribute) Checker() {
	//if a.Hero = nil {
	//
	//}
}

func (a *Attribute) Positive() {
	for key, value := range a.AttrMap {
		switch key {
		case monitorfile.BubbleIdMap("攻击力永久加成"):
			a.Hero.AttackPoint += value
		case monitorfile.BubbleIdMap("生命上限永久加成"):
			a.Hero.Health += value
			a.Hero.THealth += value
		case monitorfile.BubbleIdMap("速度"):
			a.Hero.Speed += value
		case monitorfile.BubbleIdMap("体力"):
			a.Hero.ActionPoint += value
		}
	}
}

func (a *Attribute) Negative() {
	for key, value := range a.AttrMap {
		switch key {
		case monitorfile.BubbleIdMap("攻击力永久加成"):
			a.Hero.AttackPoint -= value
		case monitorfile.BubbleIdMap("生命上限永久加成"):
			a.Hero.Health -= value
			a.Hero.THealth -= value
		case monitorfile.BubbleIdMap("速度"):
			a.Hero.Speed -= value
		case monitorfile.BubbleIdMap("体力"):
			a.Hero.ActionPoint -= value
		}
	}
}
