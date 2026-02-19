package demos

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func MakeAnimalSpeak(a Animal) {
	println(a.Speak())
}

func InterfacesDemo() {
	dog := Dog{}
	cat := Cat{}

	MakeAnimalSpeak(dog) // Output: Woof!
	MakeAnimalSpeak(cat) // Output: Meow!
}
