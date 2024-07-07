package attribute

import (
	"train/monitor/hero"
	"train/monitor/monitorfile"
)

type Attribute struct {
	Hero    *hero.Hero
	AttrMap map[int]int
}

func (a *Attribute) Publish() {
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

func (a *Attribute) ToNegative() {
	for key, value := range a.AttrMap {
		switch key {
		case monitorfile.BubbleIdMap("攻击力永久加成"):
			a.AttrMap[monitorfile.BubbleIdMap("攻击力永久加成")] = -value
		case monitorfile.BubbleIdMap("生命上限永久加成"):
			a.AttrMap[monitorfile.BubbleIdMap("生命上限永久加成")] = -value
		case monitorfile.BubbleIdMap("速度"):
			a.AttrMap[monitorfile.BubbleIdMap("速度")] = -value
			//fmt.Println("negativeL", a.AttrMap[monitorfile.BubbleIdMap("速度")])
		case monitorfile.BubbleIdMap("体力"):
			a.AttrMap[monitorfile.BubbleIdMap("体力")] = -value
		}
	}
}
