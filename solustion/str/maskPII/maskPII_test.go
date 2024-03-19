package maskPII

import (
	"fmt"
	"testing"
)

func TestMaskPII(t *testing.T) {
	fmt.Println(MaskPII("+(501321)-50-23431"))
}
