package conn

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestConn(t *testing.T) {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("请输入文本（输入'exit'退出）: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "读取输入时出错:", err)
			break
		}

		// 去除换行符
		text = strings.TrimSpace(text)

		// 检查是否输入了'exit'来退出程序
		if text == "exit" {
			fmt.Println("程序已退出。")
			break
		}

		// 打印输入内容
		fmt.Printf("你输入了: %s\n", text)
	}
}

// 注意：上面的代码示例中忘记导入了"strings"包，
// 你需要在使用strings.TrimSpace函数之前导入它：
// import "strings"
