package main

func worker(ch chan int, message int) {
	ch <- message
}

// func main() {
// 	ch := make(chan int)

// 	for i := 0; i <= 10; i++ {
// 		go worker(ch, i)
// 		message := <-ch
// 		fmt.Println(message)
// 	}

// }
