package routines

import (
	monitorcenter "train/monitor"
	"train/monitor/hero/attribute"
	"train/monitor/hero/dead"
	"train/monitor/routines/attack"
	"train/monitor/routines/beattack"
	"train/monitor/routines/damage/takedamage"
)

func Routines(types interface{}, log *monitorcenter.Logs) {
	switch types.(type) {
	case attack.Attack:
		attackData := types.(attack.Attack)
		// 攻击流程
		a := attackData.Processer()
		log.SubEvent = append(log.SubEvent, a)
		// 错误退出
		if len(a.ErrorSession) != 0 {
			return
		}
		// 迎击 并打断后续，返回
		// 流程被攻击
		ba := beattack.BeAttack{
			Name:     attackData.Name,
			Attacker: attackData.Attacker,
			Targets:  attackData.Targets,
			Damage:   a.FinalDamage,
			Mc:       attackData.Mc,
		}
		Routines(ba, log)
	case beattack.BeAttack:
		beAttackData := types.(beattack.BeAttack)
		//fmt.Println(beAttackData.Damage)
		// processer time
		ba := beAttackData.Processer()
		//fmt.Println(ba.FinalDamage)
		log.SubEvent = append(log.SubEvent, ba)

		// 出现错误退出
		if len(ba.ErrorSession) != 0 {
			return
		}
		// 反击

		// 闪避退出 or 0伤害
		if ba.IsDoge || ba.FinalDamage == 0 {
			return
		}
		td := takedamage.TakeDamage{
			Name:     beAttackData.Name,
			Attacker: beAttackData.Attacker,
			Targets:  beAttackData.Targets,
			Damage:   int(ba.FinalDamage),
			Mc:       beAttackData.Mc,
		}
		Routines(td, log)
	case takedamage.TakeDamage:
		takeDamageData := types.(takedamage.TakeDamage)
		// processer time
		tdd := takeDamageData.Processer()
		log.SubEvent = append(log.SubEvent, takeDamageData)
		// 反弹判断

		if tdd.HitBack != 0 {
			// 当攻击者或者被攻击者死亡，不再反弹
			if takeDamageData.Targets.Health <= 0 || takeDamageData.Attacker.Health <= 0 {
				return
			}
			hitBack := attack.Attack{
				Name:     "反弹伤害",
				Attacker: takeDamageData.Targets,
				Targets:  takeDamageData.Attacker,
				Damage:   int(tdd.HitBack),
				Mc:       takeDamageData.Mc,
			}
			Routines(hitBack, log)
		}
		// 没死 返回
		if takeDamageData.Targets.Health > 0 {
			return
		}
		// 死了
		death := dead.Death{
			Killer: takeDamageData.Attacker,
			Object: takeDamageData.Targets,
		}
		Routines(death, log)
	case dead.Death:
		//buffs := types.(buff.Buff)
		//buffs.Processer()
		log.SubEvent = append(log.SubEvent)
	case attribute.Attribute:
		attr := types.(attribute.Attribute)
		attr.Positive()
	}
}

// Gates THE GATES To SERFDOM
func Gates(types interface{}, mc *monitorcenter.MonitorCenter) {
	logs := monitorcenter.Logs{
		MainEvent: types,
		SubEvent:  make([]interface{}, 0),
	}
	Routines(types, &logs)
	mc.MonitorLogs = append(mc.MonitorLogs, logs)
}
