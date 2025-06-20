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
	"time"
)

func producer(stream Stream, tweetCh chan *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			break
		}

		tweetCh <- tweet
	}
	close(tweetCh)
}

func consumer(tweetCh chan *Tweet) {
	for t := range tweetCh {
		go determineTweet(t)
	}
}

func determineTweet(t *Tweet) {
	if t.IsTalkingAboutGo() {
		fmt.Println(t.Username, "\ttweets about golang")
	} else {
		fmt.Println(t.Username, "\tdoes not tweet about golang")
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	tweetCh := make(chan *Tweet)

	// Producer
	go producer(stream, tweetCh)
	consumer(tweetCh)

	fmt.Printf("Process took %s\n", time.Since(start))
}
