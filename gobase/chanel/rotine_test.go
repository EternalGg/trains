package chanel

import (
	"testing"
	"time"
)

func TestTestRotine(t *testing.T) {
	chan1 := make(chan int, 10)
	go TestRotine(chan1)
	for i := 0; i < 10; i++ {

		chan1 <- i
		chan1 <- i * 10
		time.Sleep(1 * time.Second)

	}
}
