package findingUsersActiveMinutes

import (
	"fmt"
	"testing"
)

func TestFindingUsersActiveMinutes(t *testing.T) {
	list := [][]int{{305589003, 4136}, {305589004, 4139}, {305589004, 4141}, {305589004, 4137}, {305589001, 4139}, {305589001, 4139}}
	fmt.Println(FindingUsersActiveMinutes(list, 6))
}
