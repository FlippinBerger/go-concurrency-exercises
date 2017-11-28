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
	"sync"
	"time"
)

var wg sync.WaitGroup

func producer(stream Stream, tweets chan *Tweet) {
	defer close(tweets)
	defer wg.Done()
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return
		}
		tweets <- tweet
	}
}

func consumer(tweets chan *Tweet) {
	defer wg.Done()
	for tweet := range tweets {
		if tweet.IsTalkingAboutGo() {
			fmt.Println(tweet.Username, "\ttweets about golang")
		} else {
			fmt.Println(tweet.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	ts := make(chan *Tweet, 6)

	start := time.Now()
	stream := GetMockStream()

	wg.Add(2)

	go producer(stream, ts)
	go consumer(ts)

	wg.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))
}
