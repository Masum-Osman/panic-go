package concurrent

import (
	"fmt"
	"time"
)

func sleepAndTalk(secs time.Duration, msg string, c chan string) {
	time.Sleep(secs * time.Second)
	c <- msg
}

func RunChannels() {
	c := make(chan string)

	go sleepAndTalk(0, "Hello", c)
	go sleepAndTalk(1, "Gophers!", c)
	go sleepAndTalk(2, "What's", c)
	go sleepAndTalk(7, "up?", c)

	for i := 0; i < 4; i++ {
		fmt.Println(<-c)
	}
}
