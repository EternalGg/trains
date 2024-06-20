package main

import (
	"net/http"
	"train/monitor/conn"
)

// 其他玩家管理功能，如查找、移除等...
func main() {

	http.HandleFunc("/ws", conn.Conn)

	http.ListenAndServe(":8080", nil)

}
