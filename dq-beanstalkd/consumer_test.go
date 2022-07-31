package dq_beanstalkd

import (
	"github.com/zeromicro/go-queue/dq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"testing"
)

func TestConsumer(t *testing.T) {
	consumer := dq.NewConsumer(dq.DqConf{
		Beanstalks: []dq.Beanstalk{
			{
				Endpoint: "jp.chenkaiwei.com:11300",
				Tube:     "tube",
			},
			{
				Endpoint: "gz.chenkaiwei.com:11300",
				Tube:     "tube",
			},
		},
		Redis: redis.RedisConf{
			Host: "redis.chenkaiwei.com:6379",
			Type: redis.NodeType,
			Pass: "Ckw_1988",
		}, //生产者用一个随机数表示唯一，消费者拿到消息时把随机数往redis中存一份做幂等，用来确保不会重复消费
	})
	consumer.Consume(func(body []byte) {
		logx.Info(string(body))
	})
}
