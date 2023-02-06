package simplefactory

import (
	"testing"
	"fmt"
)

// TestHiAPI test get hiapi with factory
func TestHiAPI(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	fmt.Println("result of TestHiAPI: ", s)
	if s != "Hi, Tom" {
		t.Fatal("TestHiAPI test fail")
	}
}

// TestType1 test get helloAPI with factory
func TestHelloAPI(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	fmt.Println("result of TestHelloAPI: ", s)
	if s != "Hello, Tom" {
		t.Fatal("TestHelloAPI test fail")
	}
}