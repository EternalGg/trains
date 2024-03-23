package crazy

import (
	"fmt"
	"testing"
	monitorcenter "train/monitor"
	"train/monitor/general"
	"train/monitor/monitorfile"
)

func TestCrazyMonitor(t *testing.T) {
	crazyHero := monitorcenter.Hero{
		Owner:       0,
		Id:          1,
		Health:      2,
		Name:        "狂战士",
		AttackPoint: 0,
	}

	hero1 := monitorcenter.Hero{
		Owner:       1,
		Id:          2,
		Health:      2,
		Name:        "木桩1",
		AttackPoint: 0,
	}

	hero2 := monitorcenter.Hero{
		Owner:       0,
		Id:          3,
		Health:      2,
		Name:        "木桩2",
		AttackPoint: 0,
	}
	mc := monitorcenter.MonitorCenterInit(0)

	general.GeneralHeroMonitor(&hero1, mc)
	general.GeneralHeroMonitor(&hero2, mc)
	CrazyMonitorInit(&crazyHero, mc)

	fmt.Println(mc.ListenAndFilter(2, monitorfile.MonitorIdMap("掉血")))
	fmt.Println(mc.ListenAndFilter(3, monitorfile.MonitorIdMap("掉血")))
	fmt.Println(mc.ListenAndFilter(1, monitorfile.MonitorIdMap("掉血")))

	fmt.Println(crazyHero)
}
