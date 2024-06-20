package monitors

import (
	"train/monitor/hero"
)

type (
	Monitor struct {
		Name  string
		Owner *hero.Hero // hero
		MID   int        // monitor id 根据id鉴别monitor
		Tid   int        // 临时id
		Froze bool       // 是否暂停中,暂停使用
		Logs  []string   // 发生了什么

		IsForever     bool // 是否是永久
		LifeTimeState int  // 持续时间扣减方式 0次数 1游戏时间 2单位回合
		LifeTime      int  // 持续时间

		// publishstate monitor，buff后续加成 foreverstate 直接buff加成 bubble 非加成到单位属性类加成
		PublishState  map[int]int        // 激活加成（被动加成)
		DeadState     map[int]int        // 当monitor死亡发生减少的属性（monitor还有额外的dead函数)
		Bubble        map[int]int        // 永久属性的增加与减少（包含看不见的）
		Ring          Ring               // 光环
		ListenerList  map[int]*hero.Hero // 对应Hero id的task map 多对一中的多
		ListenLicense []MonitorLicense
	}
	MonitorLicense struct {
		ListenType int // 订阅条件type 当canChange后 登场....
		Subject    int // 可订阅的英雄类型：属于谁的规则 0 全部 1 友军 2 中立 3 敌人 ...
	}
	MonitorSummary struct {
		Name    int         // monitor name
		Summary map[int]int // 属性
	}
	Ring struct {
		Effect   map[int]int
		Distance int
	}
)
