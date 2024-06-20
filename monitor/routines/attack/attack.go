package attack

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/routines/attack/normalattack"
	"train/monitor/routines/attack/solidattack"
)

// Attack 总攻击类
type Attack struct {
	Name     string
	Attacker *hero.Hero //攻击者
	Targets  *hero.Hero //攻击目标
	Damage   int
	Mc       *mc.MonitorCenter
}

func (a *Attack) Checker() (result []int) {
	if a.Attacker.Health <= 0 {
		result = append(result, monitorfile.ErrorSession("攻击者死亡"))
	}
	if a.Targets.Health <= 0 {
		result = append(result, monitorfile.ErrorSession("目标死亡"))
	}
	return
}
func (a *Attack) Calculator() (result mc.AttackCalculate) {
	if a.Damage == 0 {
		sa := normalattack.SingleAttack{
			Attacker: a.Attacker,
			Targets:  a.Targets,
			Mc:       a.Mc,
		}
		result = sa.Calculator()
	} else {
		sa := solidattack.SolidAttack{
			Name:     a.Name,
			Attacker: a.Attacker,
			Targets:  a.Targets,
			Damage:   a.Damage,
			Mc:       a.Mc,
		}
		result = sa.Calculator()

	}
	result.AttackerName = a.Attacker.Name
	result.TargetName = a.Targets.Name
	return
}

func (a *Attack) Processer() (ad mc.AttackCalculate) {
	// checker time
	if len(a.Checker()) >= 1 {
		return
	}
	// calculate time
	ad = a.Calculator()
	ad.Name = a.Name
	// processer time
	return
}

func (a *Attack) Later(ba mc.BeAttackCalculate, td mc.TakeDamage) {
	//if td.HitBack != 0 {
	//	hb := Attack{
	//		Name:     "反弹伤害",
	//		Attacker: a.Targets,
	//		Targets:  a.Attacker,
	//		Damage:   td.HitBack,
	//		Mc:       a.Mc,
	//	}
	//	hb.Processer()
	//}
}
