package main_test

//package producer

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/cmdline"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type message struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Payload string `json:"message"`
}

/* 【↓废弃】，这段代码在main里跑就行，在test里跑就不行*/
func TestProduce1(t *testing.T) {

	pusher := kq.NewPusher([]string{
		"kafka.chenkaiwei.com:9092",
	}, "kq")

	ticker := time.NewTicker(time.Millisecond)
	for round := 0; round < 3; round++ {
		<-ticker.C

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

func TestMyProduce(t *testing.T) {
	pusher := kq.NewPusher([]string{
		"106.52.133.91:9092",
	}, "kq")

	if err := pusher.Push("aaaccc"); err != nil {
		log.Fatal("err")
		log.Fatal(err)
	}

	cmdline.EnterToContinue()
}
