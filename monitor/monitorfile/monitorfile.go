package monitorfile

// MonitorIdMap Monitor Map
func MonitorIdMap(str string) int {

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
	case "被攻击后":
		return 7
	case "攻击后":
		return 18

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
	case "被buff":
		return 16
	case "释放buff":
		return 17

	case "造成伤害":
		return 19
	case "增加攻击力":
		return 20
	case "增加生命值上限":
		return 21
	case "增加体力":
		return 22
	case "速度":
		return 23
	case "野兽":
		return 24
	case "夜行动物":
		return 25
	case "体力一":
		return 26
	case "巨型食草动物":
		return 27
	case "铜墙铁壁":
		return 28
	case "野兽之王":
		return 29
	default:
		return 0
	}
}

// BubbleIdMap 属性map
func BubbleIdMap(str string) int {
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
	case "伤害减免":
		return 7
	case "伤害减免百分比":
		return 8
	case "反弹":
		return 9
	case "伤害加深百分比":
		return 10
	case "伤害加深":
		return 11
	case "攻击力永久加成":
		return 12
	case "生命上限永久加成":
		return 13
	case "加血":
		return 14
	case "体力":
		return 15
	case "速度":
		return 16
	default:
		return 0
	}
}

// BubbleIdMap To Str
func BubbleIdMapToStr(n int) string {
	switch n {
	case 1:
		return "闪避率"
	case 2:
		return "暴击率"
	case 3:
		return "暴击倍率"
	case 4:
		return "攻击次数"
	case 5:
		return "攻击加成"
	case 6:
		return "固定攻击加成"
	case 7:
		return "伤害减免"
	case 8:
		return "伤害减免百分比"
	case 9:
		return "反弹"
	case 10:
		return "伤害加深百分比"
	case 11:
		return "伤害加深"
	case 12:
		return "攻击力永久加成"
	case 13:
		return "生命上限永久加成"
	case 14:
		return "加血"

	default:
		return "空，发生错误"
	}
}

// 经济map
func EconomyMapToint(str string) int {
	switch str {
	case "初始金钱":
		return 5000
	case "初始信用":
		return 3
	case "初始贷款利率":
		return 25
	case "初始利率":
		return 5
	default:
		return 0
	}
}

// 从属关系map
func OwnerMap(str string) int {
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

// 错误信息
func ErrorSession(str string) int {
	switch str {
	case "目标死亡":
		return 1
	case "攻击者死亡":
		return 2
	default:
		return 0
	}
}

// 错误信息转str
func ErrorSessionToStr(ints int) string {
	switch ints {
	case 1:
		return "目标死亡"
	case 2:
		return "攻击者死亡"
	default:
		return ""
	}
}

func HeroNameToint(str string) int {
	switch str {
	case "狂战士":
		return 1
	case "大象":
		return 2
	case "河马":
		return 3
	case "犀牛":
		return 4
	case "豹子":
		return 5
	case "鳄鱼":
		return 6
	case "士兵":
		return 7
	case "斑马":
		return 8
	case "主宰":
		return 9
	case "自然魔像":
		return 10
	case "燃烧魔像":
		return 11
	case "章鱼":
		return 12
	case "狮子":
		return 13
	case "苍蝇":
		return 14
	case "鹿":
		return 15
	default:
		return 0
	}
}
func HeroNameToStr(id int) string {
	switch id {
	case 1:
		return "狂战士"
	case 2:
		return "大象"
	case 3:
		return "河马"
	case 4:
		return "犀牛"
	case 5:
		return "豹子"
	case 6:
		return "鳄鱼"
	case 7:
		return "士兵"
	case 8:
		return "斑马"
	case 9:
		return "主宰"
	case 10:
		return "自然魔像"
	case 11:
		return "燃烧魔像"
	case 12:
		return "章鱼"
	case 13:
		return "狮子"
	case 14:
		return "苍蝇"
	case 15:
		return "鹿"
	default:
		return ""
	}
}

//action skill map

func SkillsMap(str string) int {
	switch str {
	case "攻击":
		return 1
	case "移动":
		return 2
	case "防守":
		return 3
	case "结束回合":
		return 4
	default:
		return 0
	}
}
func SkillsIntToStrMap(i int) string {
	switch i {
	case 1:
		return "攻击"
	case 2:
		return "移动"
	case 3:
		return "防守"
	case 4:
		return "结束回合"
	case 5:
		return "野蛮冲撞"
	case 6:
		return "巨角冲撞"
	case 7:
		return "自然疗法"
	default:
		return ""
	}
}

func SessionTypeToStr(i int) string {
	switch i {
	case 0:
		return "未登陆"
	case 1:
		return "已登陆信息"
	default:
		return ""
	}
}
