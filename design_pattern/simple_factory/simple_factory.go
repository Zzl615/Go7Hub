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
	if t == 1 {
		return &HiAPI{word: "Hi"}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

type helloAPI struct{}

// Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}