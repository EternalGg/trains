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
	case "攻击":
		return 6
	case "被攻击":
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
