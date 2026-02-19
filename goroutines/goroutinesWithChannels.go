package main

// func main() {
// 	sendingChan := make(chan int)
// 	resultChan := make(chan int)

// 	var wg sync.WaitGroup
// 	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	go func() {
// 		for _, value := range nums {
// 			sendingChan <- value
// 		}
// 		close(sendingChan)
// 	}()

// 	for i := 0; i <= 3; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for value := range sendingChan {
// 				resultChan <- value * value
// 			}
// 		}()
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(resultChan)
// 	}()

// 	for result := range resultChan {
// 		fmt.Println(result)
// 	}
// }
