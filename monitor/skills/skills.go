package skills

import (
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

//Passive skills

// 被动tap 野兽 攻击力+1
func Beast() *monitors.Monitor {
	result := &monitors.Monitor{
		Name:          "野兽",
		MID:           monitorfile.MonitorIdMap("野兽"),
		Tid:           0,
		Froze:         false,
		Logs:          []string{},
		Bubble:        map[int]int{},
		IsForever:     true,
		RelianceOwner: true,
	}
	result.Bubble[monitorfile.BubbleIdMap("攻击力永久加成")] = 1
	return result
}

// 被动 夜行动物 夜间攻击力+1 移动速度-2
func NocturnalAnimal() *monitors.Monitor {
	result := &monitors.Monitor{
		Name:          "夜行动物",
		MID:           monitorfile.MonitorIdMap("夜行动物"),
		Tid:           0,
		Froze:         false,
		Logs:          []string{},
		Bubble:        map[int]int{},
		IsForever:     true,
		RelianceOwner: true,
	}
	// 距离为1的速度Buff
	result.Bubble[monitorfile.BubbleIdMap("攻击力永久加成")] = 1
	result.Bubble[monitorfile.BubbleIdMap("速度")] = -2
	//result.Bubble[monitorfile.BubbleIdMap("体力")] = 1
	return result
}

// 体力一 被动 体力+1
func Spirit1() *monitors.Monitor {
	result := &monitors.Monitor{
		Name:          "体力一",
		MID:           monitorfile.MonitorIdMap("体力一"),
		Tid:           0,
		Froze:         false,
		Logs:          []string{},
		Bubble:        map[int]int{},
		IsForever:     true,
		RelianceOwner: true,
	}
	// 距离为1的速度Buff
	result.Bubble[monitorfile.BubbleIdMap("行动点数")] = 1
	return result
}

// 巨型食草动物 当收到伤害时 攻击力+2 速度-2 体力+1
func GiantHerbivore() *monitors.Monitor {
	// 掉血，回合结束 monitor check
	result := &monitors.Monitor{
		Name:          "巨型食草动物",
		MID:           monitorfile.MonitorIdMap("巨型食草动物"),
		Tid:           0,
		Froze:         false,
		Logs:          []string{},
		Bubble:        map[int]int{},
		IsForever:     false,
		RelianceOwner: true,
	}
	selfBlood := monitors.MonitorLicense{
		ListenType: monitorfile.MonitorIdMap("掉血"),
		Subject:    monitorfile.OwnerMap("自己"),
	}
	result.ListenLicense = append(result.ListenLicense, selfBlood)
	result.Bubble[monitorfile.BubbleIdMap("体力")] = 1
	result.Bubble[monitorfile.BubbleIdMap("速度")] = -2
	result.Bubble[monitorfile.BubbleIdMap("攻击力永久加成")] = 2
	return result
}

// 铜墙铁壁 伤害减免+1
func HardSkin() *monitors.Monitor {
	result := &monitors.Monitor{
		Name:          "铜墙铁壁",
		MID:           monitorfile.MonitorIdMap("铜墙铁壁"),
		Tid:           0,
		Froze:         false,
		Logs:          []string{},
		Bubble:        map[int]int{},
		IsForever:     true,
		RelianceOwner: true,
	}
	selfAttack := monitors.MonitorLicense{
		ListenType: monitorfile.MonitorIdMap("被攻击"),
		Subject:    monitorfile.OwnerMap("自己"),
	}
	result.ListenLicense = append(result.ListenLicense, selfAttack)
	result.Bubble[monitorfile.BubbleIdMap("伤害减免")] = 1
	return result
}

// 每当有一个野兽/巨型食草动物在场 攻击力+1
func BeastKing() *monitors.Monitor {
	result := &monitors.Monitor{
		Name:          "野兽之王",
		MID:           monitorfile.MonitorIdMap("野兽之王"),
		Tid:           0,
		Froze:         false,
		Logs:          []string{},
		Bubble:        map[int]int{},
		IsForever:     true,
		RelianceOwner: true,
	}
	return result
}

// action skills
// cost,unit,money,distance
type (
	Skill struct {
		Id                 int    //id
		Name               string //名字
		OwnerId            int    //单位
		MovePoint          int    //行动点数
		Money              int    //花费金钱
		Distance           int    //距离
		TargetTeamMate     bool   //队友
		TargetEnemy        bool   // 敌人
		TargetMiddle       bool   // 中立
		TargetSelf         bool   // 自己
		TargetNoUnitPos    bool   // 无单位地形
		TargetNoMachinePos bool   // 无物品地形
		TargetUnitedPos    bool   // 有单位地形
		TargetMachinePos   bool   // 有机械地形
		NoTarget           bool   // 无目标可使用
	}
)

func StrToSkills(str string) *Skill {
	switch str {
	case "攻击":
		return Attack()
	case "移动":
		return Move()
	case "防守":
		return Defence()
	case "结束回合":
		return EndHeroTurn()
	case "野蛮冲撞":
		return CrushMove()
	case "巨角冲撞":
		return HornCrush()
	case "自然疗法":
		return NatureHealing()
	default:
		return nil
	}
}

// 默认近战攻击
func Attack() *Skill {
	att := &Skill{
		Id:           1,
		Name:         "攻击",
		MovePoint:    1,
		Money:        0,
		Distance:     1,
		TargetEnemy:  true,
		TargetMiddle: true,
		NoTarget:     false,
	}
	att.Id = monitorfile.SkillsMap(att.Name)
	return att
}

// 默认移动
func Move() *Skill {
	move := &Skill{
		Id:              2,
		Name:            "移动",
		MovePoint:       1,
		Money:           0,
		Distance:        1,
		TargetNoUnitPos: true,
		NoTarget:        false,
	}
	move.Id = monitorfile.SkillsMap(move.Name)
	return move
}

// 默认防守
func Defence() *Skill {
	defence := &Skill{
		Id:        3,
		Name:      "防守",
		MovePoint: 1,
		Money:     0,
		Distance:  0,

		NoTarget: true,
	}
	defence.Id = monitorfile.SkillsMap(defence.Name)
	return defence
}

func EndHeroTurn() *Skill {
	end := &Skill{
		Id:        4,
		Name:      "结束回合",
		MovePoint: 0,
		Money:     0,
		Distance:  0,
		NoTarget:  true,
	}
	return end
}

// 大象移动技能：行动后向距离为1的英雄单位造成1点伤害
func CrushMove() *Skill {
	end := &Skill{
		Id:              5,
		Name:            "巨兽踩踏",
		MovePoint:       1,
		Money:           0,
		Distance:        1,
		NoTarget:        false,
		TargetNoUnitPos: true,
	}
	return end
}

// 犀牛技能：巨角冲撞 移动后对附近1距离的随机敌人造成1.5倍的攻击伤害
func HornCrush() *Skill {
	end := &Skill{
		Id:              6,
		Name:            "巨角冲撞",
		MovePoint:       2,
		Money:           0,
		Distance:        1,
		NoTarget:        false,
		TargetNoUnitPos: true,
	}
	return end
}

func NatureHealing() *Skill {
	end := &Skill{
		Id:             7,
		Name:           "自然疗法",
		MovePoint:      0,
		Money:          0,
		Distance:       2,
		NoTarget:       false,
		TargetTeamMate: true,
		TargetSelf:     true,
	}
	return end
}
