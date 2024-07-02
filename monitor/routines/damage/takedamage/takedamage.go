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
	TakeDamageWithOutMC struct {
		Name     string
		Attacker *hero.Hero //攻击者
		Targets  *hero.Hero //攻击目标
		Damage   int        //技能
	}
	TakeDamageCalculate struct {
		TD      *TakeDamageWithOutMC
		Name    string
		Id      uint
		HitBack uint // 反弹伤害
	}
)

func (d *TakeDamage) Checker() {

}

func (d *TakeDamage) Calculator() (*TakeDamageCalculate, []*monitors.Monitor) {
	blood := d.Mc.ListenAndFilter(d.Targets.Tid,
		monitorfile.MonitorIdMap("掉血"))
	//fmt.Println("start")
	result := TakeDamageCalculate{
		TD:      nil,
		Name:    "被攻击！",
		Id:      0,
		HitBack: 0,
	}
	for _, monitor := range blood {
		//fmt.Println(monitor)
		for key, value := range monitor.Bubble {
			switch key {
			case monitorfile.BubbleIdMap("反弹"):
				//fmt.Println("来自", monitor.Owner.Name, "的反弹")
				result.HitBack += uint(value)
			}
		}
	}
	//fmt.Println("end")
	return &result, blood
}

func (d *TakeDamage) Processer() *TakeDamageCalculate {

	// checker time
	d.Checker()
	// calculate time
	result, blood := d.Calculator()
	// processer time
	// 减少血量
	d.Targets.THealth -= d.Damage
	// 如果blood的monitor长度>0 publish
	if len(blood) != 0 {
		d.Mc.MonitorsPublish(blood)
	}
	td := TakeDamageWithOutMC{
		Name:     d.Name,
		Attacker: d.Attacker,
		Targets:  d.Targets,
		Damage:   d.Damage,
	}
	result.TD = &td
	return result
}

func (d *TakeDamage) Later(td *TakeDamageCalculate) {
}
