package skills

import (
	"train/monitor/hero"
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

// action skills
// cost,unit,money,distance
type Skill struct {
	Id        int        //id
	Name      string     //名字
	Owner     *hero.Hero //单位
	MovePoint int        //行动点数
	Money     int        //花费金钱
	Distance  int        //距离
	Targets   []int      //0队友 1敌方单位 2中立单位 3自己
}

func StrToSkills(str string) *Skill {
	switch str {
	case "攻击":
		return Attack()
	case "移动":
		return Move()
	case "防守":
		return Defence()
	default:
		return nil
	}
}

// 默认近战攻击
func Attack() *Skill {
	att := &Skill{
		Name:      "攻击",
		MovePoint: 1,
		Money:     0,
		Distance:  1,
	}
	att.Id = monitorfile.SkillsMap(att.Name)
	return att
}

// 默认移动
func Move() *Skill {
	move := &Skill{
		Name:      "移动",
		MovePoint: 1,
		Money:     0,
		Distance:  1,
	}
	move.Id = monitorfile.SkillsMap(move.Name)
	return move
}

// 默认防守
func Defence() *Skill {
	defence := &Skill{
		Name:      "防守",
		MovePoint: 1,
		Money:     0,
		Distance:  0,
	}
	defence.Id = monitorfile.SkillsMap(defence.Name)
	return defence
}
