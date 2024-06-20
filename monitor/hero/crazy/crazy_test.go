package crazy

import (
	"fmt"
	"testing"
	monitorcenter "train/monitor"
	"train/monitor/monitorfile"
	"train/monitor/routines"
	"train/monitor/routines/attack"
	"train/monitor/routines/damage/takedamage"
)

func TestCrazyMonitor(t *testing.T) {
	mc := monitorcenter.MonitorCenterInit(0)

	CrazyMonitorInit(mc)
	CrazyMonitorInit(mc)
	hero0 := mc.HeroMap[0]
	hero1 := mc.HeroMap[1]
	hero0.Name = "狂战士1"
	hero1.Name = "狂战士2"
	attack := attack.Attack{
		Name:     "攻击test",
		Attacker: hero0,
		Targets:  hero1,
		Damage:   0,
		Mc:       mc,
	}
	routines.Gates(attack, mc)

	hero0.Intro()
	hero1.Intro()
	routines.Gates(attack, mc)
	//PrintMonitorLogs(mc.MonitorLogs[len(mc.MonitorLogs)-1])
	hero0.Intro()
	hero1.Intro()
	routines.Gates(attack, mc)
	hero0.Intro()
	hero1.Intro()
	for _, log := range mc.MonitorLogs {
		PrintMonitorLogs(log)
	}
	routines.Gates(attack, mc)
	routines.Gates(attack, mc)
	routines.Gates(attack, mc)

	//fmt.Println(mc.MonitorLogs)
}

func PrintMonitorLogs(logs monitorcenter.Logs) {
	switch logs.MainEvent.(type) {
	case string:
		fmt.Println(logs)
	case monitorcenter.AttackCalculate:
		// 断言
		ac := logs.MainEvent.(monitorcenter.AttackCalculate)
		if len(ac.ErrorSession) != 0 {
			for _, u := range ac.ErrorSession {
				fmt.Println(monitorfile.ErrorSessionToStr(u), ",取消攻击！")
			}
		}
	}
	for _, m := range logs.SubEvent {

		switch m.(type) {
		case string:
			fmt.Println(m)
		case monitorcenter.AttackCalculate:
			// 断言
			ac := m.(monitorcenter.AttackCalculate)
			if len(ac.ErrorSession) != 0 {
				for _, u := range ac.ErrorSession {
					fmt.Println(monitorfile.ErrorSessionToStr(u), ",取消攻击！")
				}
			}
			fmt.Println("造成攻击力为", ac.FinalDamage, "点的伤害！")
		case monitorcenter.BeAttackCalculate:
			ac := m.(monitorcenter.BeAttackCalculate)
			if len(ac.ErrorSession) != 0 {
				for _, u := range ac.ErrorSession {
					fmt.Println(monitorfile.ErrorSessionToStr(u), ",取消攻击！")
				}
			}

			fmt.Println("被攻击，", ac.FinalDamage, "点伤害！")
		case takedamage.TakeDamage:
			ac := m.(takedamage.TakeDamage)

			fmt.Println("掉了，", ac.Damage, "点生命值！")
			fmt.Println("----------------")
		}

	}

}
