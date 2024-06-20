package heros

import (
	"train/monitor/hero"
	"train/monitor/hero/crazy"
	"train/monitor/hero/crocodile"
	"train/monitor/hero/deer"
	"train/monitor/hero/ele"
	"train/monitor/hero/flamegolem"
	"train/monitor/hero/fly"
	"train/monitor/hero/hippo"
	"train/monitor/hero/jugg"
	"train/monitor/hero/leopard"
	"train/monitor/hero/lion"
	NatureGolem "train/monitor/hero/naturegolem"
	"train/monitor/hero/octopus"
	"train/monitor/hero/rhino"
	"train/monitor/hero/soldiers"
	"train/monitor/hero/zebra"
)

// 返回所有英雄信息 by list
func xSelectAllHeroByList() (result []*hero.Hero) {
	for i := 1; i < 100; i++ {
		if ChoseHeroById(i) == nil {
			return
		} else {
			result = append(result, ChoseHeroById(i))
		}
	}
	return
}

// 返回所有英雄信息 by map
func SelectAllHeroByMap() map[int]*hero.Hero {
	result := map[int]*hero.Hero{}
	for i := 1; i < 100; i++ {
		if ChoseHeroById(i) == nil {
			return result
		} else {
			result[i] = ChoseHeroById(i)
		}
	}
	return result
}

// 根据名称返回英雄
func ChoseHeroById(id int) *hero.Hero {
	switch id {
	case 1:
		return crazy.CrazyHeroInit()
	case 2:
		return ele.EleHeroInit()
	case 3:
		return hippo.HippoHeroInit()
	case 4:
		return rhino.RhinoHeroInit()
	case 5:
		return leopard.LeopardHeroInit()
	case 6:
		return crocodile.CrocodileHeroInit()
	case 7:
		return soldiers.SoldiersHeroInit()
	case 8:
		return zebra.ZebraHeroInit()
	case 9:
		return jugg.JuggHeroInit()
	case 10:
		return NatureGolem.NatureGolemHeroInit()
	case 11:
		return flamegolem.FlameGolemHeroInit()
	case 12:
		return octopus.OctopusHeroInit()
	case 13:
		return lion.LionHeroInit()
	case 14:
		return fly.FlyHeroInit()
	case 15:
		return deer.DeerHeroInit()
	default:
		return nil
	}
}
