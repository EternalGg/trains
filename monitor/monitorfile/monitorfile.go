package monitorfile

func MonitorIdMap(str string) uint {

	switch str {
	case "狂战士之血":
		return 1

	case "死亡":
		return 2
	case "放逐":
		return 3
	case "沉默":
		return 4

	case "攻击前":
		return 5
	case "被攻击":
		return 6
	case "攻击后":
		return 7

	case "掉血":
		return 8
	case "恢复生命值":
		return 9
	case "移动":
		return 10
	case "静止":
		return 11
	case "结束回合":
		return 12
	case "闪避":
		return 13
	case "被debuff":
		return 14
	case "释放debuff":
		return 15
	case "被buff":
		return 16
	case "释放buff":
		return 17
	default:
		return 0
	}
}
func MonitorStrMap(str uint) string {

	switch str {
	case 1:
		return "狂战士之血"

	case 2:
		return "死亡"
	case 3:
		return "放逐"
	case 4:
		return "沉默"

	case 5:
		return "攻击前"
	case 6:
		return "被攻击"
	case 7:
		return "攻击后"

	case 8:
		return "掉血"
	case 9:
		return "恢复生命值"
	case 10:
		return "移动"
	case 11:
		return "静止"
	case 12:
		return "结束回合"
	case 13:
		return "闪避"
	case 14:
		return "被debuff"
	case 15:
		return "释放debuff"
	case 16:
		return "被buff"
	case 17:
		return "释放buff"
	default:
		return ""
	}
}

func BubbleIdMap(str string) uint {
	switch str {
	case "闪避率":
		return 1
	case "暴击率":
		return 2
	case "暴击倍率":
		return 3
	case "攻击次数":
		return 4
	case "攻击加成":
		return 5
	case "固定攻击加成":
		return 6
	default:
		return 0
	}
}
func OwnerMap(str string) uint {
	switch str {
	case "己方单位不包含自己":
		return 0
	case "敌方单位":
		return 1
	case "中立单位":
		return 2
	case "自己":
		return 3
	default:
		return 4
	}
}
