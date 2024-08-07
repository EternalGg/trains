package normalattack

import (
	"fmt"
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
	"train/monitor/routines/attack/data"
)

type (
	SingleAttack struct {
		Attacker *hero.Hero //攻击者
		Targets  *hero.Hero //攻击目标
		Mc       *mc.MonitorCenter
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
func (a *SingleAttack) Calculator() data.AttackCalculate {
	fmt.Println(a.Targets.Name)
	attackerBefore := a.Mc.ListenAndFilter(
		a.Attacker.Tid,
		monitorfile.MonitorIdMap("攻击前"))
	a.Mc.MonitorsPublish(attackerBefore)
	attacker := a.Mc.HeroMonitorMap[a.Attacker]
	// publish 后
	attackData := data.AttackCalculate{
		BaseDamage:         int(a.Attacker.AttackPoint),
		DamageAddition:     0,
		CriticalHitRate:    0,
		CriticalStrikeRate: 0,
		OtherDamage:        0,
	}

	for monitor, _ := range attacker {
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
	if !mc.RandomByTime(attackData.CriticalHitRate) {
		attackData.IsCritical = false
		attackData.FinalDamage = attackData.BaseDamage + attackData.OtherDamage + attackData.DamageAddition
		fmt.Println(attackData.FinalDamage, attackData.DamageAddition)
	} else {
		attackData.IsCritical = true
		attackData.FinalDamage = (attackData.BaseDamage + attackData.DamageAddition) +
			((attackData.BaseDamage+attackData.DamageAddition)*attackData.CriticalStrikeRate)/100 +
			attackData.OtherDamage
		fmt.Println(attackData.FinalDamage, attackData.DamageAddition)
	}
	return attackData
}

func (a *SingleAttack) Checker() (result mc.SessionsAttack) {
	result.PreAttackMonitor.Changes = []monitors.MonitorSummary{}
	result.BeforeAttackMonitor.Changes = []monitors.MonitorSummary{}
	attackerBefore := a.Mc.ListenAndFilter(
		a.Attacker.Tid, monitorfile.MonitorIdMap("攻击前"))
	for _, monitor := range attackerBefore {
		c := monitors.MonitorSummary{
			Name:    monitor.MID,
			Summary: monitor.Bubble,
		}
		result.PreAttackMonitor.Changes = append(result.PreAttackMonitor.Changes, c)
	}
	enemyBeAttack := a.Mc.ListenAndFilter(
		a.Targets.Tid, monitorfile.MonitorIdMap("被攻击"))
	for _, monitor := range enemyBeAttack {
		c := monitors.MonitorSummary{
			Name:    monitor.MID,
			Summary: monitor.Bubble,
		}
		result.BeforeAttackMonitor.Changes = append(result.BeforeAttackMonitor.Changes, c)
	}
	return
}

func (a *SingleAttack) Processer() {
	//TODO implement me
	panic("implement me")
}

func (a *SingleAttack) Later() {
	//TODO implement me
	panic("implement me")
}
