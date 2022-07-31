package redis_test

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func TestZSetNil(t *testing.T) {

	//z:=&redis.Z{
	//	Score:  22,
	//	Member: "aa",
	//}

	z := &redis.Z{}

	res, err := rdb.ZAdd(ctx, "myniltest", z).Result()
	ok, err := rdb.Expire(ctx, "myniltest", 60*time.Second).Result()

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res, ok)
}

func TestNilKey(t *testing.T) {

	//总结：判断key是否不存在，最靠谱的还是单独用Exists查询一次，
	// 其他某些特定情形也可反映无key：例如Get返回redis.Nil，集合类数据查询 **所有元素** 时返回[]，除此之外其他都得靠Exists判断了
	// redis.Nil也可能反映所查询的范围内无元素，只有全范围或string才能反映key不存在

	nilKeyName := "nilKeyName"

	//stirng
	res, err := rdb.Get(ctx, nilKeyName).Result()
	fmt.Println("Get", res, err, err == redis.Nil)

	//zset
	res2, err := rdb.ZRange(ctx, nilKeyName, 0, -1).Result() // 0 -1 表示查看第一个到最后一个元素
	fmt.Println("ZRange nil key", res2, err, err == redis.Nil)
	//查询全部时，返回[]能表示key不存在

	res2a, err := rdb.ZRange(ctx, "myzset", 5, 6).Result() //
	fmt.Println("ZRange out range", res2a, err, err == redis.Nil)
	//但是查询局部时，返回[]仅表示数组越界，无法区分key是否存在

	res2b, err := rdb.ZCard(ctx, nilKeyName).Result()
	fmt.Println("ZCard", res2b, err, err == redis.Nil)

	res2c, err := rdb.Exists(ctx, "myzset").Result()
	fmt.Println("Exists myzset", res2c, err, err == redis.Nil)

	res2d, err := rdb.Exists(ctx, nilKeyName).Result()
	fmt.Println("Exists nilKeyName", res2d, err, err == redis.Nil)

	//hash
	res3, err := rdb.HGet(ctx, nilKeyName, "aaa").Result()
	fmt.Println("HGet nilKeyName", res3, err, err == redis.Nil)

	res3a, err := rdb.HGetAll(ctx, nilKeyName).Result()
	fmt.Println("HGetAll", res3a, err, err == redis.Nil)

	res3b, err := rdb.HGet(ctx, "myhash", "aaa").Result()
	fmt.Println("HGet myhash", res3b, err, err == redis.Nil) //所查field上没有元素也报redis.Nil，无法和没有key区分

	res3c, err := rdb.Get(ctx, "myhash").Result()
	fmt.Println("get myhash", res3c, err, err == redis.Nil)

	//list
	res5, err := rdb.LRange(ctx, nilKeyName, 2, 4).Result()
	fmt.Println("LRange", res5, err, err == redis.Nil)

	/*
	      === RUN   TestZSetZRank
	      [] <nil> false
	      0 <nil> false
	       redis: nil true
	      --- PASS: TestZSetZRank (0.16s)

	   结论是，字符串的Get，才会以redis.Nil表示key不存在；zset的ZRange不使用这种表达
	*/

}

func TestSet(t *testing.T) {

	//置0的时候才是永不过期，写-1会报错
	s, err := rdb.Set(ctx, "setTestKey", "1", 0).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}

func TestDelKey(t *testing.T) {

	s, err := rdb.Set(ctx, "delTestKey", "aaa", 0).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)

	i, err := rdb.Del(ctx, "delTestKey").Result()
	fmt.Println(i, err)
}
