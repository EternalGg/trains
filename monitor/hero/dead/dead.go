package dead

import (
	"fmt"
	"train/monitor/hero"
)

type Death struct {
	Killer *hero.Hero
	Object *hero.Hero
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
