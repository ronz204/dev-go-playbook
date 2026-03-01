package patterns

import "time"

func SelecterDemo() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "Data from Server 1 (Fast)"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		channel2 <- "Data from Server 2 (Slow)"
	}()

	select {
	case msg1 := <-channel1:
		println("Received:", msg1)
	case msg2 := <-channel2:
		println("Received:", msg2)
	case <-time.After(2 * time.Second):
		println("Timeout: No response received within 2 seconds")
	}

	println("all done")
}
