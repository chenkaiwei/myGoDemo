package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/cmdline"
)

type message struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Payload string `json:"message"`
}

func main() {
	pusher := kq.NewPusher([]string{
		"kafka.chenkaiwei.com:9092",
	}, "kq")

	//ticker := time.NewTicker(time.Millisecond)
	ticker := time.NewTicker(1 * time.Second) //

	for round := 0; round < 3; round++ {

		<-ticker.C //←保证每次循环有间隔？不知道图啥，注释掉好像也没事。

		count := rand.Intn(100)
		m := message{
			Key:     strconv.FormatInt(time.Now().UnixNano(), 10),
			Value:   fmt.Sprintf("%d,%d", round, count),
			Payload: fmt.Sprintf("%d,%d", round, count),
		}
		body, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))
		if err := pusher.Push(string(body)); err != nil {
			log.Fatal(err)
		}
	}

	cmdline.EnterToContinue()
}
