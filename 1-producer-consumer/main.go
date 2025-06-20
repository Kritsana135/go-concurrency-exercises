//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
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

func producer(stream Stream) (tweets []*Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return tweets
		}

		tweets = append(tweets, tweet)
	}
}

func consumer(tweets []*Tweet, cWg *sync.WaitGroup) {
	for _, t := range tweets {
		cWg.Add(1)
		go determineTweet(t, cWg)
	}
}

func determineTweet(t *Tweet, cWg *sync.WaitGroup) {
	defer cWg.Done()
	if t.IsTalkingAboutGo() {
		fmt.Println(t.Username, "\ttweets about golang")
	} else {
		fmt.Println(t.Username, "\tdoes not tweet about golang")
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	// Producer
	tweets := producer(stream)

	// init consumer waitgroup
	var cWg sync.WaitGroup

	// Consumer
	consumer(tweets, &cWg)

	cWg.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))
}
