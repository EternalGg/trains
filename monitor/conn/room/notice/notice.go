package notice

import (
	"encoding/json"
	"train/monitor/hero"
)

type (
	// notice合集 通过notice中的id来判断 如果actions&&id为空为错误代码 而id和actions非空代表有动画
	Notice struct {
		NoticeName string
		Id         int
		Actions    []ActionData
	}
	// 不同的id代表不同action data 通过id来判断data的类型
	ActionData struct {
		ID         int
		ActionName string
		Data       string
	}

	// ActionData ID= 0 ActionName := NOTICE标记
	// 每当Action开始时存入 中间所有Ation都是这个Action的SUB
	ActionStartEnd struct {
		IsStart bool //是否开始
	}

	GameResult struct {
		Success     bool
		Errcode     int
		Str         string     // 成功/错误信息
		ChoseCard   int        // 选择的卡牌ID
		BuyCardId   int        // 购买的卡牌ID
		SpendMoney  int        // 花费的金钱
		RemainMoney int        // 剩余的金钱
		SaleCardId  int        // 售卖的卡牌ID
		SaleMoney   int        // 获得的金钱
		LandingHero *hero.Hero // 登场的卡牌
		LandingPos  int        // 登场位置
		StartPos    int        // 初始位置
		NextTurn    int
	}
)

func MoveResultMade(s bool, str string, h *hero.Hero, start, landing, errcode int) *ActionData {
	m := GameResult{
		Success:     s,
		Str:         str,
		Errcode:     errcode,
		LandingHero: h,
		StartPos:    start,
		LandingPos:  landing,
	}
	jm, _ := json.Marshal(m)
	AD := ActionData{
		ID:         5,
		ActionName: "移动",
		Data:       string(jm),
	}
	return &AD
}
func ActionStart() *ActionData {
	s := ActionStartEnd{IsStart: true}
	js, _ := json.Marshal(s)
	return &ActionData{
		ID:         0,
		ActionName: "开始",
		Data:       string(js),
	}
}
func SaleResultMade(s bool, str string, err, cardid, salemoney, remainmoney int) *ActionData {
	SR := GameResult{
		Success:     s,
		Str:         str,
		Errcode:     err,
		SaleCardId:  cardid,
		SaleMoney:   salemoney,
		RemainMoney: remainmoney,
	}
	jsr, _ := json.Marshal(SR)
	AD := ActionData{
		ID:         3,
		ActionName: "出售卡牌",
		Data:       string(jsr),
	}
	return &AD
}
func BuyResultMade(s bool, str string, buycard, remainmoney, spendmoney, errcode int) *ActionData {
	BR := GameResult{
		Success:     s,
		Str:         str,
		BuyCardId:   buycard,
		SpendMoney:  spendmoney,
		RemainMoney: remainmoney,
		Errcode:     errcode,
	}
	jbr, _ := json.Marshal(BR)
	AD := ActionData{
		ID:         2,
		ActionName: "buy card",
		Data:       string(jbr),
	}
	return &AD
}
func LandingResultMade(s bool, str string, h *hero.Hero, err, pos int) *ActionData {
	lr := GameResult{
		Success:     s,
		Str:         str,
		Errcode:     err,
		LandingHero: h,
		LandingPos:  pos,
	}
	jlr, _ := json.Marshal(lr)
	AD := ActionData{
		ID:         4,
		ActionName: "登场数据",
		Data:       string(jlr),
	}
	return &AD
}

//	func BubbleResultMade(str string, h *hero.Hero, isSuccess bool, bubble map[int]int) *ActionData {
//		n := GameResult{
//			Success: isSuccess,
//			Str:     str,
//			Hero:    h,
//			Bubble:  bubble,
//		}
//		jn, _ := json.Marshal(n)
//		AD := ActionData{
//			ID:         6,
//			ActionName: "bubble add",
//			Data:       jn,
//		}
//		return &AD
//	}
//
//	func TakeDamageResultMade(s bool, str string, td []byte) *ActionData {
//		tdd := TakeDamageResult{
//			Success:    s,
//			Str:        str,
//			TakeDamage: td,
//		}
//		jtd, _ := json.Marshal(tdd)
//		AD := ActionData{
//			ID:         7,
//			ActionName: "take damage",
//			Data:       jtd,
//		}
//		return &AD
//
// }
//
//	func BeAttackResultMade(s bool, str string, ba []byte) *ActionData {
//		bar := BeAttackResult{
//			Success:  s,
//			Str:      str,
//			BeAttack: ba,
//		}
//		jbar, _ := json.Marshal(bar)
//		AD := ActionData{
//			ID:         8,
//			ActionName: "BeAttackData",
//			Data:       jbar,
//		}
//		return &AD
//	}
//
//	func AttackResultMade(s bool, str string, a []byte) *ActionData {
//		ar := AttackResult{
//			Success: s,
//			Str:     str,
//			Attack:  a,
//		}
//		jar, _ := json.Marshal(ar)
//		AD := ActionData{
//			ID:         9,
//			ActionName: "AttackData",
//			Data:       jar,
//		}
//		return &AD
//	}
func TurnEndResultMade(s bool, t int) *ActionData {
	te := GameResult{
		Success:  s,
		NextTurn: t,
	}
	jte, _ := json.Marshal(te)
	return &ActionData{
		ID:         10,
		ActionName: "结束回合！",
		Data:       string(jte),
	}
}
func ActionToNotice(adlist []ActionData, str string, id int) *Notice {
	return &Notice{
		NoticeName: str,
		Id:         id,
		Actions:    adlist,
	}
}

//func DeadResultMade(killer, object *hero.Hero) *ActionData {
//	dr := GameResult{
//		Hero:   object,
//		Killer: killer,
//	}
//	jdr, _ := json.Marshal(dr)
//	return &ActionData{
//		ID:         11,
//		ActionName: "单位死亡！",
//		Data:       jdr,
//	}
//}
