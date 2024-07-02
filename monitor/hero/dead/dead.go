package dead

import (
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

}

func (d *Death) Later() {
}
