package takedamage

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

type (
	TakeDamage struct {
		Name     string
		Attacker *hero.Hero //攻击者
		Targets  *hero.Hero //攻击目标
		Damage   int        //技能
		Mc       *mc.MonitorCenter
	}
)

func (d *TakeDamage) Checker() {

}

func (d *TakeDamage) Calculator() (result mc.TakeDamage, blood []*monitors.Monitor) {
	blood = d.Mc.ListenAndFilter(d.Targets.Tid,
		monitorfile.MonitorIdMap("掉血"))
	//fmt.Println("start")
	for _, monitor := range blood {
		//fmt.Println(monitor)
		for key, value := range monitor.Bubble {
			switch key {
			case monitorfile.BubbleIdMap("反弹"):
				//fmt.Println("来自", monitor.Owner.Name, "的反弹")
				result.HitBack += uint(value)
			}
		}
		m := monitors.MonitorSummary{
			Name:    monitor.MID,
			Summary: monitor.Bubble,
		}
		result.Sessions = append(result.Sessions, m)
	}
	//fmt.Println("end")
	return
}

func (d *TakeDamage) Processer() (result mc.TakeDamage) {
	// checker time
	d.Checker()
	// calculate time
	result, blood := d.Calculator()
	// processer time
	// 减少血量
	d.Targets.THealth -= d.Damage
	if len(blood) != 0 {
		d.Mc.MonitorsPublish(blood)
	}
	return result
}

func (d *TakeDamage) Later(td *mc.TakeDamage) {
}
