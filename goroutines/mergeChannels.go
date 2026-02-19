package main

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	merged := make(chan int)

// 	var wg sync.WaitGroup

// 	// send in channel 1
// 	go func() {
// 		for i := 0; i <= 5; i++ {
// 			ch1 <- i
// 		}
// 		close(ch1)
// 	}()

// 	// send in channel 2
// 	go func() {
// 		for i := 5; i <= 10; i++ {
// 			ch2 <- i
// 		}
// 		close(ch2)
// 	}()

// 	// add 2
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		for value := range ch1 {
// 			merged <- value
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		for value := range ch2 {
// 			merged <- value
// 		}
// 	}()

// 	go func() {
// 		wg.Wait()
// 		close(merged)
// 	}()

// 	for value := range merged {
// 		fmt.Println(value)
// 	}
// }
