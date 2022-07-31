package bloom_test

import (
	"SecKillTrainTicket/common/commonUtil"
	"fmt"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"testing"
)

//var rdb = redis.NewClient(&redis.Options{
//	Addr:     "redis.chenkaiwei.com:6379",
//	Password: "Ckw_1988", //
//	DB:       2,  //
//})
func TestBloom(t *testing.T) {

	filter := bloom.New(redis.New("redis.chenkaiwei.com:6379", redis.WithPass("Ckw_1988")), "myBloomtest", 1024)

	for i := int64(100); i < 200; i++ {
		filter.Add(commonUtil.Int64ToBytes(i))
	}

	fmt.Println(filter.Exists(commonUtil.Int64ToBytes(2)))
	fmt.Println(filter.Exists(commonUtil.Int64ToBytes(3)))
}
