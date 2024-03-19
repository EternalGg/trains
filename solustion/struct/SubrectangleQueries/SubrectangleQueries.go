package SubrectangleQueries

type SubrectangleQueries struct {
	matrix [][]int
	height int
	weight int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	s := new(SubrectangleQueries)
	s.matrix = rectangle
	s.weight = len(rectangle[0])
	s.height = len(rectangle)
	return *s
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			this.matrix[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	if row > this.height-1 {
		return -1
	}
	if col > this.weight-1 {
		return -1
	}
	return this.matrix[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
