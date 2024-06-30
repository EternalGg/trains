package data

import (
	"train/monitor/hero"
	"train/monitor/monitors"
)

type AttackWithOutMc struct {
	Name     string
	Attacker *hero.Hero //攻击者
	Targets  *hero.Hero //攻击目标
}
type AttackCalculate struct {
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
