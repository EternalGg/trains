package robot

import "fmt"

func Robot(command string, obstacles [][]int, x int, y int) bool {
	// step 1 : create function calculate how much round going
	// step 2 : is going to target ?  not->return false
	// step 3 : is going to stone ?   yes->return false
	// step 4 : return true
	vectorStep, oneX, oneY := make([]bool, len(command)), 0, 0
	for i := 0; i < len(command); i++ {
		if command[i] == 'U' {
			vectorStep[i] = true
			oneY++
		} else {
			oneX++
		}
	}
	vector, target := []int{oneX, oneY}, []int{x, y}

	// step 1,2

	tsteps := minV(x, y, vector[0], vector[1])
	tx, ty := vector[0]*tsteps, vector[1]*tsteps
	fmt.Println(tx, ty)
	if tx == vector[0] && ty == vector[1] {

		fmt.Println("1")
	} else if !Move(tx, ty, vectorStep, target) {
		fmt.Println("2")
		return false
	}

	// step 3

	for i := 0; i < len(obstacles); i++ {
		if obstacles[i][0] > target[0] || obstacles[i][1] > target[1] {
			continue
		} else {
			sx, sy := obstacles[i][0], obstacles[i][1]
			steps := minV(sx, sy, vector[0], vector[1])
			ox, oy := vector[0]*steps, vector[1]*steps
			if ox == sx && oy == sy {
				fmt.Println("3", obstacles[i])
				return false
			} else if Move(ox, oy, vectorStep, obstacles[i]) {
				fmt.Println("4", obstacles[i])
				return false
			}
		}
	}

	// step 4
	return true
}

func oneRoundMove(x, y int, vector []int) int {
	if vector[0]/x != vector[1]/y {
		return -1
	} else {
		return vector[0] / x
	}
}

func minV(sx, sy, vx, vy int) int {
	minx := sx / vx
	miny := sy / vy
	if miny < minx {
		return miny
	} else {
		return minx
	}
}

func Move(sx, sy int, vectorStep []bool, targert []int) bool {
	for i := 0; i < len(vectorStep); i++ {
		if vectorStep[i] {
			sy++
		} else {
			sx++
		}
		if sx == targert[0] && sy == targert[1] {
			return true
		}
	}
	fmt.Println(sx, sy)
	return false
}
