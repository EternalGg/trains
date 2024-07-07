package dead

import (
	"train/monitor/hero"
)

type Death struct {
	Killer *hero.Hero
	Object *hero.Hero
}
