package routines

import (
	"encoding/json"
	monitorcenter "train/monitor"
	"train/monitor/conn/room/notice"
	"train/monitor/hero/attribute"
	"train/monitor/hero/dead"
	"train/monitor/monitorfile"
	"train/monitor/routines/attack"
	"train/monitor/routines/beattack"
	"train/monitor/routines/damage/takedamage"
)

func Routines(types interface{}, log *monitorcenter.MonitorCenter) {
	switch types.(type) {
	case attack.Attack:
		attackData := types.(attack.Attack)
		// 攻击流程
		a := attackData.Processer()
		ja, _ := json.Marshal(a)
		log.MonitorLogs = append(log.MonitorLogs, notice.AttackResultMade(true, a.AttackerName, ja))
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
		jba, _ := json.Marshal(ba)
		//fmt.Println(ba.FinalDamage)
		log.MonitorLogs = append(log.MonitorLogs, notice.BeAttackResultMade(true, ba.Name, jba))

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
		jtdd, _ := json.Marshal(tdd)
		log.MonitorLogs = append(log.MonitorLogs, notice.TakeDamageResultMade(true, takeDamageData.Name, jtdd))

		// 反弹判断
		if tdd.HitBack != 0 {
			// 当攻击者或者被攻击者死亡，不再反弹
			if takeDamageData.Targets.THealth <= 0 || takeDamageData.Attacker.THealth <= 0 {
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
		if takeDamageData.Targets.THealth > 0 {
			return
		}
		// 死了
		death := dead.Death{
			Killer: takeDamageData.Attacker,
			Object: takeDamageData.Targets,
		}
		Routines(death, log)
	case dead.Death:
		//1.死亡monitor

		d := types.(dead.Death)
		deathMonitor := log.ListenAndFilter(d.Object.Tid,
			monitorfile.MonitorIdMap("死亡"))
		if len(deathMonitor) > 0 {
			log.MonitorsPublish(deathMonitor)
		}

		//2.log死亡信息
		dr := notice.DeadResultMade(d.Killer, d.Object)
		log.MonitorLogs = append(log.MonitorLogs, dr)

		//3.单位死亡 monitor pointer end
		for _, monitor := range log.MonitorMap {
			if monitor.RelianceOwner {
				log.MonitorMap[monitor.Tid] = nil
			}
		}
		// 4.单位死亡 time map end
		delete(log.Time.Time.HeroTime, d.Object)
		// 5. hero monitor map delete
		delete(log.HeroMonitorMap, d.Object)
		// 位置
		log.BattleFiled.Positions[d.Object.Pos].Hero = nil
		delete(log.HeroMap, d.Object.Tid)
		// time log
		// count log
	case attribute.Attribute:
		attr := types.(attribute.Attribute)
		attr.Publish()
	}
}

// Gates THE GATES To SERFDOM
func Gates(types interface{}, mc *monitorcenter.MonitorCenter) {
	Routines(types, mc)
	//mc.MonitorLogs = append(mc.MonitorLogs, logs)
	// should end?
}
