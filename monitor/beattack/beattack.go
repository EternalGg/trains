package beattack

import (
	mc "train/monitor"
	"train/monitor/monitorfile"
)

type (
	BeAttack struct {
		Name     string
		Attacker *mc.Hero //攻击者
		Targets  *mc.Hero //攻击目标
		Damage   int      //固定技能
		Mc       *mc.MonitorCenter
	}
)

func (B *BeAttack) Checker() (result []uint) {
	if B.Targets.Health <= 0 {
		result = append(result, monitorfile.ErrorSession("目标死亡"))
	}
	if B.Attacker.Health <= 0 {
		result = append(result, monitorfile.ErrorSession("攻击者死亡"))
	}
	return
}

// Calculator 计算闪避率,伤害减免,伤害减免百分比
func (B *BeAttack) Calculator() (result mc.BeAttackCalculate, beAttack []*mc.Monitor) {
	beAttack = B.Mc.ListenAndFilter(
		B.Attacker.Id,
		monitorfile.MonitorIdMap("被攻击"))
	for _, monitor := range beAttack {
		for key, value := range monitor.Bubble {
			switch key {
			case monitorfile.BubbleIdMap("闪避率"):
				result.DogeRate += value
			case monitorfile.BubbleIdMap("伤害减免"):
				result.DamageReduce += value
			case monitorfile.BubbleIdMap("伤害减免百分比"):
				result.DamageReduceRate += value
			case monitorfile.BubbleIdMap("伤害加深"):
				result.DamageDepth += uint(value)
			case monitorfile.BubbleIdMap("伤害加深百分比"):
				result.DamageDepthRate += uint(value)
			}
		}
		m := mc.MonitorSummary{
			Name:    monitor.MID,
			Summary: monitor.Bubble,
		}
		result.Sessions = append(result.Sessions, m)
	}
	result.Name = B.Name
	return
}

func (B *BeAttack) Processer() (BA mc.BeAttackCalculate) {
	// check time
	//BA.ErrorSession = B.Checker()
	if len(BA.ErrorSession) >= 1 {
		return
	}
	// calculate time
	beAttackData, monitors := B.Calculator()
	// process time
	// 反击
	// 迎击
	if mc.RandomByTime(beAttackData.DogeRate) {
		// 如果闪避，则返回闪避以及publish闪避
		BA.IsDoge = true
		doge := B.Mc.ListenAndFilter(B.Targets.Id,
			monitorfile.MonitorIdMap("闪避"))
		B.Mc.Publish(doge)
	} else {
		// 如果没有闪避，则publish
		// 伤害增加计算
		B.Damage = B.Damage -
			(((B.Damage * BA.DamageReduceRate) / 100) -
				BA.DamageReduce)
		// 伤害加深计算
		B.Damage += ((B.Damage * int(BA.DamageDepthRate)) / 100) +
			int(BA.DamageDepth)
		// 如果伤害大于0 进入掉血环节
		BA.FinalDamage = B.Damage
		B.Mc.Publish(monitors)
	}

	BA.Name = B.Name
	B.Later()
	return
}

// 被攻击后,没有被闪避
func (B *BeAttack) Later() {

}
