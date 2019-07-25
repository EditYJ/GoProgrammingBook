package book3

import (
	"fmt"
	"testing"
)

type Weekday int

const(
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)



func TestConstType(t *testing.T) {
	fmt.Println(Saturday)
}
