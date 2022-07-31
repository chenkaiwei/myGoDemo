package dq_beanstalkd

import (
	"fmt"
	"github.com/zeromicro/go-queue/dq"
	"testing"
	"time"
)

func newPruducer() dq.Producer {

	return dq.NewProducer([]dq.Beanstalk{
		{
			Endpoint: "jp.chenkaiwei.com:11300",
			Tube:     "tube",
		},
		{
			Endpoint: "gz.chenkaiwei.com:11300",
			Tube:     "tube",
		},
	})
}

func TestProducerAt(t *testing.T) {

	producer := newPruducer()

	time.ParseInLocation("20060102", "20220518", time.Local)

	tt := time.Date(2022, 7, 4, 20, 13, 00, 00, time.Local)
	res, err := producer.At([]byte("aaa"), tt)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)

	//for i := 1000; i < 1005; i++ {
	//	_, err := producer.Delay([]byte(strconv.Itoa(i)), time.Second*5)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}
}

func TestProducerDelay(t *testing.T) {

	producer := newPruducer()

	res, err := producer.Delay([]byte("bbb"), 3*time.Second)
	//res, err := producer.Delay([]byte("bbb"),0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}

//延时消息撤回，用来做 未支付订单xx分钟后过期 功能，即若已支付则撤回对应延迟消息
func TestProducerRevoke(t *testing.T) {

	producer := newPruducer()

	cccres, err := producer.Delay([]byte("ccc"), 3*time.Second)
	if err != nil {
		fmt.Println(err.Error())
	}
	dddres, err := producer.Delay([]byte("ddd"), 3*time.Second)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(cccres)
	fmt.Println(dddres)

	time.Sleep(1 * time.Second)
	err = producer.Revoke(cccres) //添加消息时的返回值即roveke（撤回消息）时的入参
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("revoke id：", cccres)
}

func TestProducerMulti(t *testing.T) {

	producer := newPruducer()

	for i := 0; i < 100; i++ {
		res, err := producer.Delay([]byte("bbb"), 0)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(res)
	}

}
