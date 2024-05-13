package solidattack

import (
	mc "train/monitor"
	"train/monitor/monitorfile"
)

type SolidAttack struct {
	Name     string
	Attacker *mc.Hero //攻击者
	Targets  *mc.Hero //攻击目标
	Damage   int
	Mc       *mc.MonitorCenter
}

func (a *SolidAttack) Checker() {

}

func (a *SolidAttack) Calculator() (result mc.AttackCalculate) {

	attackerBefore := a.Mc.ListenAndFilter(
		a.Attacker.Id,
		monitorfile.MonitorIdMap("攻击前"))
	a.Mc.Publish(attackerBefore)
	attacker := a.Mc.ListenAndFilter(a.Attacker.Id,
		monitorfile.MonitorIdMap("攻击前"))
	// publish 后
	for _, monitor := range attacker {
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
