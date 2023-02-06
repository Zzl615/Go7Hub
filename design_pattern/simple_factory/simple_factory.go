package simplefactory

import "fmt"

// API is interface
type API interface {
	Say(name string) string
}

// HiAPI is one of API implement
type HiAPI struct {
	word string
}

func (a *HiAPI) Say(name string) string {
	return fmt.Sprintf("%s, %s", a.word, name)
}

// NewAPI return API instance by type
func NewAPI(t int) API {
	return &HiAPI{word: "Hi"}
}