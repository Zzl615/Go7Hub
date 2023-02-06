package design_pattern

import "fmt"

// API is interface
type API interface {
	Say(name string) string
}

// HiAPI is one of API implement
type HiAPI struct {
	word string = "Hi!"
}

func (*HiAPI) Say(name string) string {
	return fmt.Sprintf("%s, %s", HiAPI.word, name)
}

// NewAPI return API instance by type
func NewAPI(t int) API {
	return &HiAPI{}
}