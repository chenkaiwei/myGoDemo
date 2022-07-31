package redis_test

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

//这个可以作为调试脚本的模板
func TestMGet(t *testing.T) {

	script := redis.NewScript(`
	local key = KEYS[1]
	local fields = ARGV

	local numArr =redis.call("HMGET",KEYS[1],unpack(fields))
	return numArr
	`) //lua里，#表示长度
	//注意，下标从1开始

	res, err := script.Run(ctx, rdb, []string{"myhash"}, "a", "b", "c", "d").Result()

	fmt.Println(res, err)
}
