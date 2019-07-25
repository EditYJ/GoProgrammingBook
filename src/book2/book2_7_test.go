// 作用域
package book2

import (
	"log"
	"os"
	"testing"
)

var cwd string

func TestScopeArea(t *testing.T) {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd Failed: %v", err)
	}
}
