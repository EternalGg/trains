package solidattack

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/routines/attack/data"
)

type SolidAttack struct {
	Name     string
	Attacker *hero.Hero //攻击者
	Targets  *hero.Hero //攻击目标
	Damage   int
	Mc       *mc.MonitorCenter
}

func (a *SolidAttack) Checker() {

}

func (a *SolidAttack) Calculator() (result data.AttackCalculate) {

	attackerBefore := a.Mc.ListenAndFilter(
		a.Attacker.Tid,
		monitorfile.MonitorIdMap("攻击前"))
	a.Mc.MonitorsPublish(attackerBefore)
	attacker := a.Mc.HeroMonitorMap[a.Attacker]
	// publish 后
	for monitor, _ := range attacker {
		for key, value := range monitor.Bubble {
			switch key {
			case monitorfile.BubbleIdMap("固定攻击加成"):
				result.OtherDamage += value
			}
		}
	}
	a.Damage += result.OtherDamage
	result.FinalDamage = a.Damage
	return
}

func (a *SolidAttack) Processer() {

}

func (a *SolidAttack) Later() {
}
