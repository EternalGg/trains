package hero

import "fmt"

type (
	Hero struct {
		Owner       int // 0 GamerA 1 GamerB 2 中立单位
		Id          int // HeroId
		Tid         int // 临时id
		Health      int // (最大生命值)总血量
		THealth     int // 临时血量
		Name        string
		AttackPoint int
		ActionPoint int         //行动点
		Price       int         //价格
		Speed       int         //速度
		GameTempo   map[int]int //游戏内临时属性
		RoundTempo  map[int]int //本轮次临时属性
	}
)

func (h Hero) Intro() {
	fmt.Println(h.Name, "生命值", h.THealth, ";", "攻击力", h.AttackPoint)
}
