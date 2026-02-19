package patterns

import "fmt"

type State int

const (
	StateA State = iota
	StateB
	StateC
)

func (s State) String() string {
	return [...]string{"StateA", "StateB", "StateC"}[s]
}

func EnumsDemo() {
	var current State = StateB
	var label string = current.String()
	fmt.Printf("Current state: %s\n", label)
}
