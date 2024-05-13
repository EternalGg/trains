package dead

import (
	"fmt"
	monitorcenter "train/monitor"
)

type Death struct {
	Killer *monitorcenter.Hero
	Object *monitorcenter.Hero
}

func (d *Death) Checker() {

}

func (d *Death) Calculator() {

}

func (d *Death) Processer() {
	//
	fmt.Println("dead")
}

func (d *Death) Later() {
}
