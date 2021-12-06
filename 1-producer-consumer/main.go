//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(link chan<- *Tweet, stream Stream) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(link)
			break
		}
		link <- tweet
	}
}

func consumer(link <-chan *Tweet, done chan<- bool) {
	for t := range link {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
	done <- true
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	link := make(chan *Tweet, 10)
	done := make(chan bool)

	// Producer
	go producer(link, stream)

	// Consumer
	go consumer(link, done)

	<-done
	fmt.Printf("Process took %s\n", time.Since(start))
}
