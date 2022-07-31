package redis_test

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

//
// redis-go的官方示例

// https://redis.uptrace.dev/guide/lua-scripting.html#redis-script

func TestCounter(t *testing.T) {

	var incrBy = redis.NewScript(`
	local key = KEYS[1]
	local change = ARGV[1]

	local value = redis.call("GET", key)
	if not value then
	value = 0
	end

	value = value + change
	redis.call("SET", key, value)

	return value`)

	keys := []string{"my_counter"}
	values := []interface{}{+3}
	num, err := incrBy.Run(ctx, rdb, keys, values...).Int()

	if err != nil {

		fmt.Println(err.Error())
	} else {
		fmt.Println(num)
	}
}

//Passing multiple values
func TestMultipleValues(t *testing.T) {

	var sum = redis.NewScript(`local key = KEYS[1]
	local sum = redis.call("GET", key)
	if not sum then
	sum = 0
	end

	local num_arg = #ARGV
	for i = 1, num_arg do
	sum = sum + ARGV[i]
	end

	redis.call("SET", key, sum)

	return sum`)

	res, err := sum.Run(ctx, rdb, []string{"my_sum"}, 1, 2, 3, 4, 5).Int()
	fmt.Println(res, err)
	// Output: 6 nil
}

//Loop continue

func TestLoopContinue(t *testing.T) {

	/*
		Lua does not support continue statements in loops,
		but you can emulate it with a nested repeat loop and a break statement:

	*/

	//var script:=redis.NewScript(`
	//local num_arg = #ARGV
	//
	//for i = 1, num_arg do
	//repeat
	//
	//if true then
	//do break end -- continue
	//end
	//
	//until true
	//end`)
}
