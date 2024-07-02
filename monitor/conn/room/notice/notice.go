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
		Data       []byte
	}

	// ActionData ID= 0 ActionName := NOTICE标记
	// 每当Action开始时存入 中间所有Ation都是这个Action的SUB
	ActionStartEnd struct {
		IsStart bool //是否开始
	}

	// ChoseCard ID= 1 ActionName := 卡牌选择
	CardChoseResult struct {
		Success   bool
		Str       string //成功选择
		ChoseCard int    //卡牌id
	}
	// Buy ID= 2 ActionName := 商店购买
	BuyResult struct {
		Success     bool
		ErrCode     int
		Str         string //成功/错误代码
		BuyCardId   int    //购买卡牌id
		SpendMoney  int    //花费金钱
		RemainMoney int    //剩余金钱
	}
	// Sale ID= 3 ActionName := 商店售卖
	SaleResult struct {
		Success     bool
		Str         string
		ErrCode     int
		SaleCardId  int //购买卡牌id
		SaleMoney   int //获得金钱
		RemainMoney int //剩余金钱
	}
	// Landing ID= 4 ActionName := 登场数据
	LandingResult struct {
		Success     bool
		Str         string     //是否成功登陆
		ErrCode     int        //错误代码
		LandingHero *hero.Hero //登场卡牌
		LandingPos  int        //登场位置
	}
	// MoveData ID = 5 ActionName = 移动数据
	MoveResult struct {
		Success     bool
		Str         string
		ErrCode     int
		LandingHero *hero.Hero //移动英雄
		StartPos    int        //初始位置
		LandingPos  int        //移动位置
	}
	// BubbleData ID = 6 ActionName = 泡泡数据
	BubbleResult struct {
		Success bool
		Str     string      //是否成功增加
		Hero    *hero.Hero  //登场卡牌
		Bubble  map[int]int //bubble增加的数据
	}
	// TakeDamageData Id = 7 ActionName = TakeDamage
	TakeDamageResult struct {
		Success    bool
		Str        string
		TakeDamage []byte
	}
	// BeAttackData Id = 8 ActionName = BeAttackData
	BeAttackResult struct {
		Success  bool
		Str      string
		BeAttack []byte
	}
	// AttackData Id = 9 ActionName = AttackData
	AttackResult struct {
		Success bool
		Str     string
		Attack  []byte
	}
	// TurnEndData id = 10 ActionName = TurnEndData
	TurnEndResult struct {
		Success  bool
		NextTurn int
	}
	// HeroDeadData id = 11 ActionName = DeadData
	DeadResult struct {
		Hero   *hero.Hero
		Killer *hero.Hero
	}
)

func MoveResultMade(s bool, str string, h *hero.Hero, start, landing, errcode int) *ActionData {
	m := MoveResult{
		Success:     s,
		Str:         str,
		ErrCode:     errcode,
		LandingHero: h,
		StartPos:    start,
		LandingPos:  landing,
	}
	jm, _ := json.Marshal(m)
	AD := ActionData{
		ID:         5,
		ActionName: "移动",
		Data:       jm,
	}
	return &AD
}
func ActionStart() *ActionData {
	s := ActionStartEnd{IsStart: true}
	js, _ := json.Marshal(s)
	return &ActionData{
		ID:         0,
		ActionName: "开始",
		Data:       js,
	}
}
func SaleResultMade(s bool, str string, err, cardid, salemoney, remainmoney int) *ActionData {
	SR := SaleResult{
		Success:     s,
		Str:         str,
		ErrCode:     err,
		SaleCardId:  cardid,
		SaleMoney:   salemoney,
		RemainMoney: remainmoney,
	}
	jsr, _ := json.Marshal(SR)
	AD := ActionData{
		ID:         3,
		ActionName: "出售卡牌",
		Data:       jsr,
	}
	return &AD
}
func BuyResultMade(s bool, str string, buycard, remainmoney, spendmoney, errcode int) *ActionData {
	BR := BuyResult{
		Success:     s,
		Str:         str,
		BuyCardId:   buycard,
		SpendMoney:  spendmoney,
		RemainMoney: remainmoney,
		ErrCode:     errcode,
	}
	jbr, _ := json.Marshal(BR)
	AD := ActionData{
		ID:         2,
		ActionName: "buy card",
		Data:       jbr,
	}
	return &AD
}
func LandingResultMade(s bool, str string, h *hero.Hero, err, pos int) *ActionData {
	lr := LandingResult{
		Success:     s,
		Str:         str,
		ErrCode:     err,
		LandingHero: h,
		LandingPos:  pos,
	}
	jlr, _ := json.Marshal(lr)
	AD := ActionData{
		ID:         4,
		ActionName: "登场数据",
		Data:       jlr,
	}
	return &AD
}
func BubbleResultMade(str string, h *hero.Hero, isSuccess bool, bubble map[int]int) *ActionData {
	n := BubbleResult{
		Success: isSuccess,
		Str:     str,
		Hero:    h,
		Bubble:  bubble,
	}
	jn, _ := json.Marshal(n)
	AD := ActionData{
		ID:         6,
		ActionName: "bubble add",
		Data:       jn,
	}
	return &AD
}
func TakeDamageResultMade(s bool, str string, td []byte) *ActionData {
	tdd := TakeDamageResult{
		Success:    s,
		Str:        str,
		TakeDamage: td,
	}
	jtd, _ := json.Marshal(tdd)
	AD := ActionData{
		ID:         7,
		ActionName: "take damage",
		Data:       jtd,
	}
	return &AD

}
func BeAttackResultMade(s bool, str string, ba []byte) *ActionData {
	bar := BeAttackResult{
		Success:  s,
		Str:      str,
		BeAttack: ba,
	}
	jbar, _ := json.Marshal(bar)
	AD := ActionData{
		ID:         8,
		ActionName: "BeAttackData",
		Data:       jbar,
	}
	return &AD
}
func AttackResultMade(s bool, str string, a []byte) *ActionData {
	ar := AttackResult{
		Success: s,
		Str:     str,
		Attack:  a,
	}
	jar, _ := json.Marshal(ar)
	AD := ActionData{
		ID:         9,
		ActionName: "AttackData",
		Data:       jar,
	}
	return &AD
}
func TurnEndResultMade(s bool, t int) *ActionData {
	te := TurnEndResult{
		Success:  s,
		NextTurn: t,
	}
	jte, _ := json.Marshal(te)
	return &ActionData{
		ID:         10,
		ActionName: "结束回合！",
		Data:       jte,
	}
}
func ActionToNotice(adlist []ActionData, str string, id int) *Notice {
	return &Notice{
		NoticeName: str,
		Id:         id,
		Actions:    adlist,
	}
}
func DeadResultMade(killer, object *hero.Hero) *ActionData {
	dr := DeadResult{
		Hero:   object,
		Killer: killer,
	}
	jdr, _ := json.Marshal(dr)
	return &ActionData{
		ID:         11,
		ActionName: "单位死亡！",
		Data:       jdr,
	}
}
