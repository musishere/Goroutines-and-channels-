package main

import (
	"fmt"
	"time"
)

// func main() {
// 	// go doSomeMagic()
// 	// doSomeMagic()

// 	fmt.Println("Hello main")
// 	go func() {
// 		fmt.Println("Hello from anonymous goroutine")
// 	}()
// 	time.Sleep(time.Millisecond * 500)
// }

func doSomeMagic() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("Hi there. My name is mustafa ", i)
	}
}
