package monitorcenter

import (
	"fmt"
	"testing"
)

func TestMonitorStart(t *testing.T) {

	//crazymonitor := crazy.CrazyMonitor(crazy)
	for i := 0; i < 100; i++ {
		fmt.Println(RandomNumber(3))
	}
}

func TestAddHeroMapping(t *testing.T) {

}
