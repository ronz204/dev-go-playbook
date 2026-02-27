package demos

type Stack[Data any] struct {
	elements []Data
}

func (s *Stack[Data]) Push(element Data) {
	s.elements = append(s.elements, element)
}

func (s *Stack[Data]) Pop() Data {
	if len(s.elements) == 0 {
		var zero Data
		return zero
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

func (s Stack[Data]) Display() {
	// use blank identifier to ignore the index
	for _, element := range s.elements {
		println(element)
	}
}

func GenericsDemo() {
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	println("Popped from intStack:", intStack.Pop()) // Output: 30
	intStack.Display()                               // Output: 10, 20

	stringStack := Stack[string]{}
	stringStack.Push("Hello")
	stringStack.Push("World")

	println("Popped from stringStack:", stringStack.Pop()) // Output: World
	stringStack.Display()                                  // Output: Hello
}
