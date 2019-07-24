package book2

import (
	"GoProgrammingBook/src/book2/tempconv"
	"fmt"
	"testing"
)
var pc [256]byte
func TestPackage(t *testing.T) {
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.FToC(tempconv.Fdegree))
	fmt.Println(tempconv.KToC(100))
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	fmt.Println(len(pc))
	fmt.Println(pc)
}
