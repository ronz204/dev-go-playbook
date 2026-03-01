package demos

func writter(out chan<- int) {
	defer close(out)

	for i := range 5 {
		out <- i
	}
}

func reader(in <-chan int) {
	for n := range in {
		println(n)
	}
}

func ConcurrencyDemo() {
	numbers := make(chan int)

	go writter(numbers)
	println("numbers are being writed")

	reader(numbers)
	println("all numbers are read")
}
