package chanel

import "fmt"

func TestRotine(ch chan int) {
	for {
		select {
		case <-ch:
			for data := range ch {
				fmt.Println(data)
			}
		}
	}
}
