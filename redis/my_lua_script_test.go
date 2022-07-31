package redis_test

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

/*
在终端测试redis语句时

$ redis-cli

select 2   //切换到DB2，不要影响主库

 EVAL "local key = KEYS[1] local value = redis.call(\"GET\", key)  return value" 1 my_counter

*/

//这个可以作为调试脚本的模板
func Test1(t *testing.T) {

	script := redis.NewScript(`
	local key = KEYS[1]
	local value = redis.call("GET", key)
	return {value,#ARGV,ARGV[1]}
	`) //lua里，#表示长度
	//注意，下标从1开始

	res, err := script.Run(ctx, rdb, []string{"my_counter"}, 1, 1, 1, 1, 1).Result()

	fmt.Println(res, err)
}

//【demo】原子性减库存的lua脚本
func TestStorageReduce(t *testing.T) {

	//返回0时表示库存已经为0无法再减（返回0则是从1减下来的，属于成功）
	script := redis.NewScript(`
	local goodsLeftKey = KEYS[1]
	local goodsLeft = tonumber(redis.call("GET", goodsLeftKey))
	if (goodsLeft>0)
	then
	goodsLeft =redis.call("DECR",goodsLeftKey)
	else
	goodsLeft = 0
	end
	return goodsLeft
	`) //lua里，#表示长度
	//注意，下标从1开始

	res, err := script.Run(ctx, rdb, []string{"goodsLeft"}).Result()

	fmt.Println(res, err)
}

func TestKeyExists(t *testing.T) {

	//返回0时表示库存已经为0无法再减（返回0则是从1减下来的，属于成功）
	script := redis.NewScript(`
	
	local keyExist 
	keyExist=not (redis.call("EXISTS",KEYS[1])==0)

	return keyExist

	`) //lua里，#表示长度
	//注意，下标从1开始

	res, err := script.Run(ctx, rdb, []string{"myKey1"}).Result()

	fmt.Println(res, err)
}

//调试 批量获取小段购买权sectionItem，并生成订单号OrderNo
func TestSectionItemBuyRight(t *testing.T) {

	script := redis.NewScript(`

	local fields=ARGV
	local isEnough
	local orderNoGeneratorKey="orderNoGeneratorKey"

	local keyExist = (redis.call("EXISTS",KEYS[1])==1)
	if (not keyExist) then 
		return {keyExist}
	end 

    --订单号基准键是否存在
	local orderNoGeneratorExist =(redis.call("EXISTS",orderNoGeneratorKey)==1)
	if (not orderNoGeneratorExist) then
		return {keyExist,orderNoGeneratorExist}
	end

	-- 检查余量够不够
	local numArr=redis.call("HMGET",KEYS[1],unpack(fields))
	isEnough=true
	for i, num in ipairs(numArr) do
		if((0+num)<=0)
		then 
			isEnough=false
		return {keyExist,orderNoGeneratorExist,isEnough}
		end
	end

	--余量够了，对应小段购买权 -1 
	local resNumArr={}
	if isEnough
	then 
		for i,field in ipairs(fields)
		do
			resNumArr[i]=redis.call("HINCRBY",KEYS[1],field,-1)
		end
	end

	--扣减成功，生成对应订单号
	
	local orderNo=redis.call("INCRBY",orderNoGeneratorKey,1)

	return {keyExist,orderNoGeneratorExist,isEnough,orderNo} 

	`)
	/*相关lua语法：

		-- #表示长度
		-- 下标从1开始
		-- 脚本中批量查询（HMGET）：https://stackoverflow.com/questions/29594517/lua-script-in-redis-hmget-with-table
		-- bool对应：（返回后是 nil | 1 是正常现象，脚本里还是false|true）
			false	redis.Nil error
			true	int64(1)
	    --  非，不是 ！，而是 not
	*/

	fields := []interface{}{3, 2, 4}
	//fields:=[]interface{}{4,6,5}    //6 是0
	//fields:=[]interface{}{7,8}      //---[buyRight [5 6 7] [33 0 33] <nil>] <nil>

	res, err := script.Run(ctx, rdb, []string{"buyRight"}, fields...).Result()

	fmt.Println(res, err)
}
