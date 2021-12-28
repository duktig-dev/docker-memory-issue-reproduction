package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"redis_publisher/config"
	"redis_publisher/lib"
	"sync"
	"time"
)

var ctx = context.Background()
var wg = sync.WaitGroup{}
var logger lib.Logger
var cnf = config.Values
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var mut sync.Mutex

func main() {
	publish()
}

func publish() {

	rand.Seed(time.Now().Unix())

	fmt.Println("Connecting to Redis Server")
	logger.LogINFO("Connecting to Redis Server")

	i := 1

	f, err := os.OpenFile(fmt.Sprintf("/log/%s.log", cnf.PublisherId), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	for {
		select {
		case <-time.After(time.Millisecond * cnf.PublishInterval):

			mut.Lock()

			randomWord := randSeq(10)
			randomNumber := rand.Intn(100) + rand.Intn(100)
			//randomNumber2 := rand.Intn(100)
			randomIdNumber := rand.Intn(100) + rand.Intn(101) + rand.Intn(104)
			randomWordId := randSeq(5)

			payload := fmt.Sprintf(`{"event":"Golang-testing","service":"events.go.publisher.tests","published_time":"%s","data":{"go-publisher":"%s","random-number":%d, "number":%d,"strltr":"test-%s"},"appendix":{"event_id":"EVENT-%s%d"}}`,
				time.Now().Format(time.RFC3339),
				cnf.PublisherId,
				randomNumber,
				i,
				randomWord,
				randomWordId,
				randomIdNumber,
			)

			f.WriteString(payload + "\n")

			mut.Unlock()

			if cnf.PublishAmount > 0 && cnf.PublishAmount == i {
				logger.LogINFO("Finished Publishing from golang")
				os.Exit(0)
			}

			i++

		}
	}

}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
