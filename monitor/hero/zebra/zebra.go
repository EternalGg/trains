package zebra

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
	"train/monitor/skills"
)

func ZebraHeroInit() *hero.Hero {
	Zebra := hero.Hero{
		Id:          monitorfile.HeroNameToint("斑马"),
		Health:      5,
		THealth:     5,
		Name:        "斑马",
		AttackPoint: 1,
		Price:       250,
		ActionPoint: 1,

		Speed:          6,
		PositiveSkills: []int{1, 2, 3, 4},
	}
	return &Zebra
}

func Skill1Zebra(hero *hero.Hero) (result *monitors.Monitor) {
	result = skills.Spirit1()
	// 体力+1
	result.Owner = hero
	return result
}

func Skill2Zebra(hero *hero.Hero) (result *monitors.Monitor) {
	result = &monitors.Monitor{
		Name:          "速度光环",
		Owner:         hero,
		MID:           monitorfile.MonitorIdMap("增加体力"),
		Tid:           0,
		Froze:         false,
		Logs:          []string{},
		IsForever:     true,
		LifeTimeState: 0,
		LifeTime:      0,
		Ring: monitors.Ring{
			Effect:   map[int]int{},
			Distance: 1,
		},
	}
	// 距离为1的速度Buff
	result.Ring.Effect[monitorfile.BubbleIdMap("速度")] = 2
	result.Ring.Distance = 1
	return result
}

func ZebraMonitorInit(mcc *mc.MonitorCenter, h *hero.Hero) {
	//zb := ZebraHeroInit()
	//mcc.AddHeroInHeroMap(zb)
	// 被动技能 skill1 体力+1
	// 被动技能 skill2 移动速度光环 附近1距离友方单位移动速度+2 （Speed - 2）
	skill1 := Skill1Zebra(h)
	skill2 := Skill2Zebra(h)
	mcc.MonitorsActive([]*monitors.Monitor{skill1, skill2})
}
