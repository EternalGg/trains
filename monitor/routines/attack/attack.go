package attack

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
	"train/monitor/routines/attack/normalattack"
	"train/monitor/routines/attack/solidattack"
)

// Attack 总攻击类

type (
	Attack struct {
		Name     string
		Attacker *hero.Hero //攻击者
		Targets  *hero.Hero //攻击目标
		Damage   int
		Mc       *mc.MonitorCenter
	}
	AttackWithOutMc struct {
		Name     string
		Attacker *hero.Hero //攻击者
		Targets  *hero.Hero //攻击目标
	}
	AttackCalculate struct {
		A                  AttackWithOutMc
		AttackerName       string
		TargetName         string
		Name               string
		Id                 uint
		BaseDamage         int                       // 基础攻击力
		DamageAddition     int                       // 攻击加成
		CriticalHitRate    int                       // 暴击概率
		CriticalStrikeRate int                       // 暴击倍率加成 0为100%暴击率（无暴击）
		OtherDamage        int                       // 固定伤害加成
		IsCritical         bool                      // 是否暴击
		FinalDamage        int                       // 最终伤害
		Sessions           []monitors.MonitorSummary // 信息
		ErrorSession       []int                     // 错误信息

	}
)

func (a *Attack) Checker() (result []int) {
	if a.Attacker.Health <= 0 {
		result = append(result, monitorfile.ErrorSession("攻击者死亡"))
	}
	if a.Targets.Health <= 0 {
		result = append(result, monitorfile.ErrorSession("目标死亡"))
	}
	return
}
func (a *Attack) Calculator() (result AttackCalculate) {
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

func (a *Attack) Processer() (ad AttackCalculate) {
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

//func (a *Attack) Later(ba mc.BeAttackCalculate, td mc.TakeDamage) {
//	//if td.HitBack != 0 {
//	//	hb := Attack{
//	//		Name:     "反弹伤害",
//	//		Attacker: a.Targets,
//	//		Targets:  a.Attacker,
//	//		Damage:   td.HitBack,
//	//		Mc:       a.Mc,
//	//	}
//	//	hb.Processer()
//	//}
//}
