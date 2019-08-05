package actor

type Actor struct {
	Name string
	messageBox chan interface{}
}
