package attack

import (
	mc "train/monitor"
	"train/monitor/monitorfile"
)

// Attack 总攻击类
type (
	Attack struct {
		Name     string
		Attacker *mc.Hero //攻击者
		Targets  *mc.Hero //攻击目标
		Damage   int      //固定技能
		Mc       mc.MonitorCenter
	}
)

// 攻击行为
//  1. 计算返回
//     pre attack
//     攻击后可能的状态
//     before attack
//     这三个攻击session
//  2. 如果确认攻击 则计算attack
//     返回攻击结算session
//  3. 攻击，攻击后的monitor
func (a *Attack) Calculate() mc.PreAttackCalculate {
	attackerBefore := a.Mc.ListenAndFilter(
		a.Attacker.Id,
		monitorfile.MonitorIdMap("攻击前"))
	a.Mc.Publish(attackerBefore)
	attacker := a.Mc.ListenAndFilter(a.Attacker.Id,
		monitorfile.MonitorIdMap("攻击前"))
	// publish 后
	attackData := mc.PreAttackCalculate{
		BaseDamage:         int(a.Attacker.AttackPoint),
		DamageAddition:     0,
		CriticalHitRate:    0,
		CriticalStrikeRate: 0,
		OtherDamage:        0,
	}
	for _, monitor := range attacker {
		for key, value := range monitor.Bubble {
			switch key {
			case monitorfile.BubbleIdMap("暴击率"):
				attackData.CriticalHitRate += value
			case monitorfile.BubbleIdMap("暴击倍率"):
				attackData.CriticalStrikeRate += value
			case monitorfile.BubbleIdMap("攻击加成"):
				attackData.DamageAddition += value
			case monitorfile.BubbleIdMap("固定攻击加成"):
				attackData.OtherDamage += value
			}
		}
	}
	// 如果暴击率为0或者没有没有暴击
	// 加成完毕 ，计算伤害
	//2 是暴击-暴击后 （伤害+伤害加成）*暴击倍率 + 固定加成
	//3 否暴击 （伤害+伤害加成）+ 固定加成
	if attackData.CriticalHitRate <= 0 || !mc.RandomByTime(attackData.CriticalHitRate) {
		attackData.IsCritical = false
		attackData.FinalDamage = attackData.BaseDamage + attackData.OtherDamage + attackData.DamageAddition
	} else {
		attackData.IsCritical = true
		attackData.FinalDamage = (attackData.BaseDamage + attackData.DamageAddition) +
			((attackData.BaseDamage+attackData.DamageAddition)*attackData.CriticalStrikeRate)/100 +
			attackData.OtherDamage
	}
	return attackData
}
func (a *Attack) Session() (result mc.SessionsAttack) {
	result.Name = a.Name
	result.PreAttackMonitor.Changes = []mc.MonitorSummary{}
	result.BeforeAttackMonitor.Changes = []mc.MonitorSummary{}
	attackerBefore := a.Mc.ListenAndFilter(
		a.Attacker.Id, monitorfile.MonitorIdMap("攻击前"))
	for _, monitor := range attackerBefore {
		c := mc.MonitorSummary{
			Name:    monitor.MID,
			Summary: monitor.Bubble,
		}
		result.PreAttackMonitor.Changes = append(result.PreAttackMonitor.Changes, c)
	}
	enemyBeAttack := a.Mc.ListenAndFilter(
		a.Targets.Id, monitorfile.MonitorIdMap("被攻击"))
	for _, monitor := range enemyBeAttack {
		c := mc.MonitorSummary{
			Name:    monitor.MID,
			Summary: monitor.Bubble,
		}
		result.BeforeAttackMonitor.Changes = append(result.BeforeAttackMonitor.Changes, c)
	}
	return
}
