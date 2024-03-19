package pan

import "fmt"

func Pan() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("lmao")
}
