package conn

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
	"train/monitor/conn/room"
)

// 玩家conn
// 行动：移动 登陆 技能....
// 银行：借贷 存钱
// 商店：购买 出售
// 早上：官方商店
// 中午：官方商店 银行
// 晚上：野外商店
// 问题：交流 conn pattern
// 早上：购买，部署
// 早上-中午：行动
// 中午：银行 购买 部署
// 中午-晚上：行动
// 晚上：购买，部署
// 晚上-早上：行动
type (
	PlayerManager struct {
		mu          sync.RWMutex
		players     map[int]*Player
		playerState map[*Player]int //1 大厅 2 匹配 3 对局
	}
	Player struct {
		ID   int
		Name string
	}
	ServerSession struct {
		Token []byte // token验证
		Type  int    // 0未登陆 1登陆player信息 2大厅等待信息 3卡牌选择信息 4游戏内时间与State
		// 5 single chose card
		Data []byte
	}
	Lobby struct {
		WaitingTime int      //等待时间
		GameMod     []string //可以选择的游戏模式 1自由
	}
)

func StrToSession(str string) *room.ClientSession {
	//fmt.Println(str == "登陆", str[:2] == "登陆")
	result := room.ClientSession{
		Token: nil,
		Type:  0,
		Data:  nil,
	}
	// 1.登陆
	// 2.登陆 返回登陆总时长 后选择游戏模式
	// 3.进入游戏模式后 发送信息包括游戏时间
	switch str {
	case "登陆":
		result.Type = 1
		player := Player{
			ID:   1,
			Name: "test1",
		}
		result.Data, _ = json.Marshal(player)
		return &result
	default:
		return &result
	}
}

var PM *PlayerManager
var upGrader = &websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		players:     make(map[int]*Player),
		playerState: map[*Player]int{},
	}
}
func (pm *PlayerManager) AddPlayer(p *Player) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.players[p.ID] = p

}
func (pm *PlayerManager) ChangePlayerState(p *Player, state int) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.playerState[p] = state
}
func init() {
	PM = NewPlayerManager()
}

func WaitLogin(conn *websocket.Conn, user *Player) bool {
	_, p, err := conn.ReadMessage()

	Sessions := room.ClientSession{}
	json.Unmarshal(p, &Sessions)
	switch Sessions.Type {
	case 1:
		player := Player{}
		json.Unmarshal(Sessions.Data, &player)
		user.ID = player.ID
		user.Name = player.Name
		PM.AddPlayer(user)
		PM.ChangePlayerState(user, 1)
		conn.WriteMessage(1, LoginSession(user))
		return false
	}
	if err != nil {
		fmt.Println(err)
	}
	return true
}
func WaitLobby(conn *websocket.Conn, lobby *Lobby) {
	time.Sleep(5000000000)
	lobby.WaitingTime += 5
	lm, _ := json.Marshal(lobby)
	session := room.ClientSession{
		Type: 2,
		Data: lm,
	}
	fmt.Println(lobby.WaitingTime)
	WaitlobbySession, _ := json.Marshal(session)
	conn.WriteMessage(1, WaitlobbySession)
}
func LoginSession(player *Player) (result []byte) {
	session := room.ClientSession{}
	session.Type = 0
	switch player.ID {
	case 0:
		session.Type = 0
		session.Data = []byte("未登陆！")
		result, _ = json.Marshal(session)
		return
	default:
		session.Type = 1
		sd, _ := json.Marshal(player)
		session.Data = sd
		// 应该返还token 先跳过
		result, _ = json.Marshal(session)
		return
	}
}

func Conn(w http.ResponseWriter, r *http.Request) {
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 用户信息
	NowPlayer := &Player{}
	// 读取客户端消息
	// 如果未登陆 则卡在登陆里

	if NowPlayer.ID == 0 {
		for WaitLogin(conn, NowPlayer) {
			conn.WriteMessage(1, LoginSession(NowPlayer))
		}
	}

	// 进入大厅，发送大厅信息以及时间
	IsLobby := true
	lobby := Lobby{
		WaitingTime: 0,
		GameMod:     []string{"free"},
	}

	go func() {
		for IsLobby {
			WaitLobby(conn, &lobby)
		}
	}()
	Game := room.TestingGameInit(NowPlayer.ID)
	defer func() {
		conn.Close()
	}()
	for true {
		_, M, _ := conn.ReadMessage()
		Sessions := room.ClientSession{}
		json.Unmarshal(M, &Sessions)

		switch Sessions.Type {
		case 4:
			//fmt.Println("zzzzzz", PM.playerState[NowPlayer])

			if PM.playerState[NowPlayer] == 1 {
				IsLobby = false
				PM.ChangePlayerState(NowPlayer, 3)
				go func() {
					go Game.GameStart()
					// 休眠 先让GameStart里的Is over=false
					time.Sleep(1000000)

					for Game.GameState.GameState != "end" {
						select {
						case sessions := <-Game.ChOut:
							result, _ := json.Marshal(sessions)
							conn.WriteMessage(1, result)
						}
					}
				}()
			}
		case 5:
			conn.Close()
			return
		case 6:
			// input 游戏数据游戏开始 进入
			roomsession := room.ClientSession{}
			roomsession.Data = Sessions.Data
			Game.Ch <- &roomsession
		}

	}

}
