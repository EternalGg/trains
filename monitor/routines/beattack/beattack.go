package beattack

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

type (
	BeAttack struct {
		Name     string
		Attacker *hero.Hero //攻击者
		Targets  *hero.Hero //攻击目标
		Damage   int        //固定技能
		Mc       *mc.MonitorCenter
	}
	BeAttackWithOutMc struct {
		Name     string
		Attacker *hero.Hero //攻击者
		Targets  *hero.Hero //攻击目标
	}
	BeAttackCalculate struct {
		BA               *BeAttackWithOutMc
		Name             string
		Id               int
		DogeRate         int                       // 闪避率
		DamageReduce     int                       // 伤害减免
		DamageReduceRate int                       // 伤害减免百分比
		IsDoge           bool                      // 是否闪避
		FinalDamage      int                       // 最终伤害
		Sessions         []monitors.MonitorSummary // 信息
		ErrorSession     []int                     // 错误信息
		FightBack        bool                      // 反击
		DamageDepthRate  int                       // 伤害加深率
		DamageDepth      int                       // 伤害加深数值
	}
)

func (B *BeAttack) Checker() (result []int) {
	if B.Targets.Health <= 0 {
		result = append(result, monitorfile.ErrorSession("目标死亡"))
	}
	if B.Attacker.Health <= 0 {
		result = append(result, monitorfile.ErrorSession("攻击者死亡"))
	}
	return
}

// Calculator 计算闪避率,伤害减免,伤害减免百分比
func (B *BeAttack) Calculator() (result *BeAttackCalculate, beAttack []*monitors.Monitor) {
	result = &BeAttackCalculate{
		BA:               nil,
		Name:             "",
		Id:               0,
		DogeRate:         0,
		DamageReduce:     0,
		DamageReduceRate: 0,
		IsDoge:           false,
		FinalDamage:      0,
		Sessions:         nil,
		ErrorSession:     nil,
		FightBack:        false,
		DamageDepthRate:  0,
		DamageDepth:      0,
	}
	beAttack = B.Mc.ListenAndFilter(
		B.Targets.Tid,
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
				result.DamageDepth += (value)
			case monitorfile.BubbleIdMap("伤害加深百分比"):
				result.DamageDepthRate += (value)
			}
		}
		m := monitors.MonitorSummary{
			Name:    monitor.MID,
			Summary: monitor.Bubble,
		}
		result.Sessions = append(result.Sessions, m)
	}
	result.Name = B.Name
	return
}

func (B *BeAttack) Processer() (BA *BeAttackCalculate) {
	// check time
	//BA.ErrorSession = B.Checker()

	// calculate time
	BA, monitors := B.Calculator()
	// process time
	// 反击
	// 迎击
	if mc.RandomByTime(BA.DogeRate) {
		// 如果闪避，则返回闪避以及publish闪避
		BA.IsDoge = true
		doge := B.Mc.ListenAndFilter(B.Targets.Tid,
			monitorfile.MonitorIdMap("闪避"))
		B.Mc.MonitorsPublish(doge)
	} else {
		// 如果没有闪避，则publish
		// 伤害增加计算
		//fmt.Println(BA.DamageReduce, B.Damage)
		B.Damage -= ((B.Damage * BA.DamageReduceRate) / 100) +
			BA.DamageReduce
		//fmt.Println(BA.DamageReduce, B.Damage)
		// 伤害加深计算
		B.Damage += ((B.Damage * int(BA.DamageDepthRate)) / 100) +
			int(BA.DamageDepth)
		// 如果伤害大于0 进入掉血环节
		BA.FinalDamage = B.Damage
		B.Mc.MonitorsPublish(monitors)
	}

	BA.Name = B.Name
	B.Later()
	BA.BA = &BeAttackWithOutMc{
		Name:     B.Name,
		Attacker: B.Attacker,
		Targets:  B.Targets,
	}
	return
}

// 被攻击后,没有被闪避
func (B *BeAttack) Later() {

}

// 当单位死亡 自己monitor的所有monitorend 以及所有依赖于自己的from
// 先把检查自己monitor的做好
// 兽王：每当monitorcount变动 就检查一遍 监控monitor count变动
