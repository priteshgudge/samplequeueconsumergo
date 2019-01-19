package operators

import "fmt"

//Message is essentially a task. Can be made into an interface for flexibility
type Message struct {
	Action string
}

func (m *Message) performAction() {
	fmt.Printf("Action performed: %v\n", m.Action)
}
