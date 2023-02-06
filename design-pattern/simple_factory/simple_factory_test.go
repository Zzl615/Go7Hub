package simplefactory

import (
	"testing"
	"github.com/Zzl615/Go7Hub/design-pattern/simplefactory"
)


// TestType1 test get hiapi with factory
func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	if s != "Hi, Tom" {
		t.Fatal("Type1 test fail")
	}
}