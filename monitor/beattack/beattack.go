package beattack

import mc "train/monitor"

type (
	BeAttack struct {
		Name     string
		Attacker *mc.Hero //攻击者
		Targets  *mc.Hero //攻击目标
		Damage   int      //固定技能
		Mc       mc.MonitorCenter
	}
)
