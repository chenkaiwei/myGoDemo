package consumer

import (
	"fmt"
	"testing"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
)

func TestConsumer(t *testing.T) {
	var c kq.KqConf
	conf.MustLoad("config.yaml", &c)

	q := kq.MustNewQueue(c, kq.WithHandle(func(k, v string) error {
		fmt.Printf("=> %s\n", v)
		return nil
	}))
	defer q.Stop()
	q.Start()
}
