package main

import (								
	"fmt"
	"strings"
)

func wordcount(s string,ch chan int){

	words := strings.Fields(s)

	ch<- len(words)

}
func main(){


	sentences := []string{
    "Go is a statically typed, compiled programming language designed at Google that is known for its simplicity and efficiency.",
    "Concurrency is not parallelism, and Go makes it easier to write programs that can handle many tasks at once using goroutines.",
    "The Go runtime scheduler efficiently multiplexes thousands of goroutines onto a small number of OS threads, enabling lightweight concurrency.",
    "Channels in Go are powerful tools for communication and synchronization between goroutines, making it easier to manage concurrent tasks.",
    "Understanding the core principles of Go such as interfaces, slices, and goroutines is essential to writing idiomatic and high-performance Go code.",
}


	ch := make(chan int)

	for _ ,sentence := range sentences{
		go wordcount(sentence,ch)
	}

	sum:=0
	for i:=0 ;i<len(sentences);i++{
		words := <- ch
		sum += words
		fmt.Println("Words in respective sentence:", words)
	}

	fmt.Println("Total words:",sum)

}