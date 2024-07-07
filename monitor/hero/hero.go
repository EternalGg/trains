package hero

import (
	"fmt"
)

type (
	Hero struct {
		Owner          int // 0 GamerA 1 GamerB 2 中立单位
		Id             int // HeroId
		Tid            int // 临时id
		Health         int // (最大生命值)总血量
		THealth        int // 临时血量
		Pos            int // 位置id
		Name           string
		AttackPoint    int
		ActionPoint    int   //行动点
		TActionPoint   int   //回合行动点
		Price          int   //价格
		Speed          int   //速度
		PositiveSkills []int //skill id
		Logs           []string
	}
)

func (h Hero) Intro() {
	fmt.Println(h.Name, "生命值", h.THealth, ";", "攻击力", h.AttackPoint)
}
