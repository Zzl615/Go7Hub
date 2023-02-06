package simplefactory

import (
	"testing"
	"fmt"
)

// TestType1 test get hiapi with factory
func TestHiAPI(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	fmt.Println("result of TestType1: ", s)
	if s != "Hi, Tom" {
		t.Fatal("Type1 test fail")
	}
}