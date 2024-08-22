package battlefiled

import (
	"fmt"
	"train/monitor/conn/room/notice"
	"train/monitor/hero"
)

type BattleFiled struct {
	Positions map[int]*Position
}

type Position struct {
	Id       int
	Hero     *hero.Hero  // 单位
	Machine  *hero.Hero  // 物品
	Distance map[int]int // 距离
}

type PositionStr struct {
	Id      string
	HeroId  string
	Machine string
}

func NormalGame() *BattleFiled {
	f := GameMap29(BattleFiled{Positions: map[int]*Position{}})
	return f
}

// 1v1游戏地图 29格
func GameMap29(bf BattleFiled) *BattleFiled {
	// first line
	bf.Positions[1] = &Position{
		Id:       1,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[2] = &Position{
		Id:       2,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[3] = &Position{
		Id:       3,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[4] = &Position{
		Id:       4,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[5] = &Position{
		Id:       5,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	// two line
	bf.Positions[10] = &Position{
		Id:       10,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[11] = &Position{
		Id:       11,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[12] = &Position{
		Id:       12,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[13] = &Position{
		Id:       13,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[14] = &Position{
		Id:       14,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[15] = &Position{
		Id:       15,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	// three line
	bf.Positions[20] = &Position{
		Id:       20,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[21] = &Position{
		Id:       21,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[22] = &Position{
		Id:       22,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[23] = &Position{
		Id:       23,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[24] = &Position{
		Id:       24,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[25] = &Position{
		Id:       25,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[26] = &Position{
		Id:       26,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	// four line
	bf.Positions[30] = &Position{
		Id:       30,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[31] = &Position{
		Id:       31,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[32] = &Position{
		Id:       32,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[33] = &Position{
		Id:       33,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[34] = &Position{
		Id:       34,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[35] = &Position{
		Id:       35,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	// five line
	bf.Positions[40] = &Position{
		Id:       40,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[41] = &Position{
		Id:       41,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[42] = &Position{
		Id:       42,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[43] = &Position{
		Id:       43,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[44] = &Position{
		Id:       4,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	// base
	bf.Positions[0] = &Position{
		Id:       0,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	bf.Positions[45] = &Position{
		Id:       45,
		Hero:     nil,
		Machine:  nil,
		Distance: map[int]int{},
	}
	Distance29(bf)
	for i := 0; i < 45; i++ {
		if bf.Positions[i] != nil {
			bf.Positions[i].Distance[i] = 0
		}
	}
	return &bf
}

func Distance29(bf BattleFiled) *BattleFiled {
	// line 1
	bf.Positions[0].Distance[1] = 1
	bf.Positions[0].Distance[10] = 1
	bf.Positions[0].Distance[11] = 1
	bf.Positions[0].Distance[12] = 1

	bf.Positions[1].Distance[0] = 1
	bf.Positions[1].Distance[2] = 1
	bf.Positions[1].Distance[11] = 1
	bf.Positions[1].Distance[12] = 1
	bf.Positions[1].Distance[13] = 1

	bf.Positions[2].Distance[1] = 1
	bf.Positions[2].Distance[3] = 1
	bf.Positions[2].Distance[12] = 1
	bf.Positions[2].Distance[13] = 1
	bf.Positions[2].Distance[14] = 1

	bf.Positions[3].Distance[2] = 1
	bf.Positions[3].Distance[4] = 1
	bf.Positions[3].Distance[14] = 1
	bf.Positions[3].Distance[13] = 1

	bf.Positions[4].Distance[3] = 1
	bf.Positions[4].Distance[5] = 1
	bf.Positions[4].Distance[14] = 1
	bf.Positions[4].Distance[15] = 1

	bf.Positions[5].Distance[4] = 1
	bf.Positions[5].Distance[15] = 1

	// line 2
	bf.Positions[10].Distance[0] = 1
	bf.Positions[10].Distance[11] = 1
	bf.Positions[10].Distance[21] = 1
	bf.Positions[10].Distance[20] = 1

	bf.Positions[11].Distance[0] = 1
	bf.Positions[11].Distance[1] = 1
	bf.Positions[11].Distance[10] = 1
	bf.Positions[11].Distance[12] = 1
	bf.Positions[11].Distance[20] = 1
	bf.Positions[11].Distance[21] = 1
	bf.Positions[11].Distance[22] = 1

	bf.Positions[12].Distance[0] = 1
	bf.Positions[12].Distance[1] = 1
	bf.Positions[12].Distance[2] = 1
	bf.Positions[12].Distance[11] = 1
	bf.Positions[12].Distance[13] = 1
	bf.Positions[12].Distance[21] = 1
	bf.Positions[12].Distance[22] = 1

	bf.Positions[13].Distance[1] = 1
	bf.Positions[13].Distance[2] = 1
	bf.Positions[13].Distance[3] = 1
	bf.Positions[13].Distance[11] = 1
	bf.Positions[13].Distance[14] = 1
	bf.Positions[13].Distance[22] = 1
	bf.Positions[13].Distance[23] = 1

	bf.Positions[14].Distance[2] = 1
	bf.Positions[14].Distance[3] = 1
	bf.Positions[14].Distance[4] = 1
	bf.Positions[14].Distance[13] = 1
	bf.Positions[14].Distance[23] = 1
	bf.Positions[14].Distance[24] = 1

	bf.Positions[15].Distance[4] = 1
	bf.Positions[15].Distance[5] = 1
	bf.Positions[15].Distance[24] = 1
	bf.Positions[15].Distance[25] = 1
	bf.Positions[15].Distance[26] = 1

	// line 3
	bf.Positions[20].Distance[10] = 1
	bf.Positions[20].Distance[11] = 1
	bf.Positions[20].Distance[21] = 1
	bf.Positions[20].Distance[30] = 1

	bf.Positions[21].Distance[10] = 1
	bf.Positions[21].Distance[11] = 1
	bf.Positions[21].Distance[12] = 1
	bf.Positions[21].Distance[20] = 1
	bf.Positions[21].Distance[22] = 1
	bf.Positions[21].Distance[30] = 1

	bf.Positions[22].Distance[11] = 1
	bf.Positions[22].Distance[12] = 1
	bf.Positions[22].Distance[13] = 1
	bf.Positions[22].Distance[21] = 1
	bf.Positions[22].Distance[31] = 1
	bf.Positions[22].Distance[30] = 1

	bf.Positions[23].Distance[13] = 1
	bf.Positions[23].Distance[14] = 1
	bf.Positions[23].Distance[31] = 1
	bf.Positions[23].Distance[32] = 1

	bf.Positions[24].Distance[14] = 1
	bf.Positions[24].Distance[15] = 1
	bf.Positions[24].Distance[25] = 1
	bf.Positions[24].Distance[32] = 1
	bf.Positions[24].Distance[33] = 1
	bf.Positions[24].Distance[34] = 1

	bf.Positions[25].Distance[15] = 1
	bf.Positions[25].Distance[24] = 1
	bf.Positions[25].Distance[26] = 1
	bf.Positions[25].Distance[35] = 1
	bf.Positions[25].Distance[33] = 1
	bf.Positions[25].Distance[34] = 1

	bf.Positions[26].Distance[15] = 1
	bf.Positions[26].Distance[25] = 1
	bf.Positions[26].Distance[34] = 1
	bf.Positions[25].Distance[33] = 1

	// line 4
	bf.Positions[30].Distance[20] = 1
	bf.Positions[30].Distance[21] = 1
	bf.Positions[30].Distance[22] = 1
	bf.Positions[30].Distance[40] = 1
	bf.Positions[30].Distance[41] = 1

	bf.Positions[31].Distance[22] = 1
	bf.Positions[31].Distance[23] = 1
	bf.Positions[31].Distance[32] = 1
	bf.Positions[31].Distance[41] = 1
	bf.Positions[31].Distance[42] = 1
	bf.Positions[31].Distance[43] = 1

	bf.Positions[32].Distance[23] = 1
	bf.Positions[32].Distance[24] = 1
	bf.Positions[32].Distance[33] = 1
	bf.Positions[32].Distance[31] = 1
	bf.Positions[32].Distance[42] = 1
	bf.Positions[32].Distance[43] = 1
	bf.Positions[32].Distance[44] = 1

	bf.Positions[33].Distance[24] = 1
	bf.Positions[33].Distance[25] = 1
	bf.Positions[33].Distance[32] = 1
	bf.Positions[33].Distance[34] = 1
	bf.Positions[33].Distance[45] = 1
	bf.Positions[33].Distance[43] = 1
	bf.Positions[33].Distance[44] = 1

	bf.Positions[34].Distance[24] = 1
	bf.Positions[34].Distance[25] = 1
	bf.Positions[34].Distance[26] = 1
	bf.Positions[34].Distance[33] = 1
	bf.Positions[34].Distance[35] = 1
	bf.Positions[34].Distance[45] = 1
	bf.Positions[34].Distance[44] = 1

	bf.Positions[35].Distance[25] = 1
	bf.Positions[35].Distance[26] = 1
	bf.Positions[35].Distance[34] = 1
	bf.Positions[35].Distance[45] = 1

	// line 5
	bf.Positions[40].Distance[30] = 1
	bf.Positions[40].Distance[41] = 1

	bf.Positions[41].Distance[30] = 1
	bf.Positions[41].Distance[31] = 1
	bf.Positions[41].Distance[40] = 1
	bf.Positions[41].Distance[42] = 1

	bf.Positions[42].Distance[32] = 1
	bf.Positions[42].Distance[31] = 1
	bf.Positions[42].Distance[41] = 1
	bf.Positions[42].Distance[43] = 1

	bf.Positions[43].Distance[32] = 1
	bf.Positions[43].Distance[31] = 1
	bf.Positions[43].Distance[33] = 1
	bf.Positions[43].Distance[44] = 1
	bf.Positions[43].Distance[42] = 1

	bf.Positions[44].Distance[32] = 1
	bf.Positions[44].Distance[34] = 1
	bf.Positions[44].Distance[33] = 1
	bf.Positions[44].Distance[43] = 1
	bf.Positions[44].Distance[45] = 1

	bf.Positions[45].Distance[35] = 1
	bf.Positions[45].Distance[34] = 1
	bf.Positions[45].Distance[33] = 1
	bf.Positions[45].Distance[44] = 1
	df := DistanceCompute29(&bf)

	return df
}

func DistanceCompute29(bf *BattleFiled) *BattleFiled {

	for k := 0; k < 50; k++ {
		for j := 1; j <= 10; j++ {
			// 假设最大距离为6
			// 遍历所有为距离为j的 当距离为j 则进入 j的map中 当他的附近1 则为相邻
			if bf.Positions[int(k)] == nil {
				continue
			}
			for pos, value := range bf.Positions[int(k)].Distance {
				if value != int(j) {
					continue
				} else {
					//fmt.Println(pos, value)
					for key, v := range bf.Positions[pos].Distance {
						if v == 1 && bf.Positions[int(k)].Distance[key] == 0 {
							bf.Positions[int(k)].Distance[key] = int(j + 1)
						}
					}
				}
			}
		}
	}

	return bf
}

func (b *BattleFiled) HeroMove(p1, p2 *Position, h *hero.Hero) *notice.ActionData {
	p2.Hero = p1.Hero
	p1.Hero = nil
	h.Pos = p2.Id
	fmt.Println(h.Pos, "英雄移动到了这里！")
	return notice.MoveResultMade(true, "移动成功", h, p1.Id, p2.Id, 0)
}
