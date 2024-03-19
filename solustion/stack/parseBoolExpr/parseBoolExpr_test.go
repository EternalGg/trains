package parseBoolExpr

import (
	"fmt"
	"testing"
)

func TestParseBoolExpr(t *testing.T) {
	fmt.Println(ParseBoolExpr("|(f,f,f,t)"))
}
