package main

import "fmt"

// define the download() function
func download(size int, ch chan int) {
	var total int
	for i := 1; i <= size; i++ {
		total += i
	}
	ch <- total
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	var s1, s2, s3 int
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	fmt.Scanln(&s3)

	go download(s1, ch1)
	go download(s2, ch2)
	go download(s3, ch3)

	// output the sum of all results
	fmt.Println(<-ch1 + <-ch2 + <-ch3)
}
