package main

import (
	"fmt"
	"sync"
	"time"
)

func greet(phrase string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(4) // nambahin 4 tugas

	go greet("Nice to meet you!", &wg)
	go greet("How are you?", &wg)
	go slowGreet("How ... are ... you ...?", &wg)
	go greet("I hope you're liking the course!", &wg)

	wg.Wait()
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func greet(phrase string, doneChan chan bool) {
// 	fmt.Println("Hello!", phrase)
// 	doneChan <- true
// }

// func slowGreet(phrase string, doneChan chan bool) {
// 	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
// 	fmt.Println("Hello!", phrase)
// 	doneChan <- true
// 	close(doneChan) // di pakai kalau kamu tau ini akan di proses terakhir
// }

// func slowGreet(phrase string, doneChan chan bool) {
// 	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
// 	fmt.Println("Hello!", phrase)
// 	doneChan <- true
// }

// func main() {
// 	done := make(chan bool)
// 	go greet("Nice to meet you!", done)
// 	go greet("How are you?", done)
// 	go slowGreet("How ... are ... you ...?", done)
// 	go greet("I hope you're liking the course!", done)

// 	// for doneChan := range done {
// 	// 	fmt.Println(doneChan)
// 	// }
// 	for range done {
// 	}
// }

// func main() {
// 	done := make(chan bool)
// 	go greet("Nice to meet you!", done)
// 	go greet("How are you?", done)
// 	go slowGreet("How ... are ... you ...?", done)
// 	go greet("I hope you're liking the course!", done)

// 	// menerima 4 hasil
// 	for i := 0; i < 4; i++ {
// 		<-done
// 	}
// }

// func main() {
// 	dones := make([]chan bool, 4)

// 	dones[0] = make(chan bool)
// 	go greet("Nice to meet you!", dones[0])
// 	dones[1] = make(chan bool)
// 	go greet("How are you?", dones[1])
// 	dones[2] = make(chan bool)
// 	go slowGreet("How ... are ... you ...?", dones[2])
// 	dones[3] = make(chan bool)
// 	go greet("I hope you're liking the course!", dones[3])

// 	for _, done := range dones {
// 		<-done
// 	}
// }

// Multiple  chan
// func main() {
// 	done := make(chan bool)
// 	go greet("Nice to meet you!", done)
// 	go greet("How are you?", done)
// 	go slowGreet("How ... are ... you ...?", done)
// 	go greet("I hope you're liking the course!", done)

// 	<-done
// 	<-done
// 	<-done
// 	<-done
// }
