package ParkingSystem

import "fmt"

type ParkingSystem struct {
	big    int
	small  int
	medium int
	bN     int
	sN     int
	mN     int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	newPark := ParkingSystem{big: big, medium: medium, small: small}
	return newPark
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		fmt.Println(this.sN+1, this.small)
		if this.sN+1 == this.small {
			return false
		} else {
			this.sN++
		}
	case 2:
		if this.mN+1 == this.medium {
			return false
		} else {
			this.mN++
		}
	case 3:
		if this.bN+1 == this.big {
			return false
		} else {
			this.bN++
		}
	}
	return true
}
