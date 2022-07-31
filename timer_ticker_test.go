package myDemo_test

import (
	"fmt"
	"testing"
	"time"
)

//https://blog.csdn.net/qq_39445165/article/details/124285999
func TestTimer(t *testing.T) {

	//1. timer基本使用
	timer1 := time.NewTimer(4 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)
	fmt.Println("4s到了")

	//2. timer只能响应1次, 到截止时间会将时间发给通道
	//timer2 := time.NewTimer(time.Second)
	//for {
	//	<-timer2.C
	//	fmt.Println("时间到", t2)
	//}

	// 3.timer实现延时的功能
	time.Sleep(time.Second)
	fmt.Println(time.Now())
	timer3 := time.NewTimer(2 * time.Second)
	<-timer3.C
	fmt.Println("2秒到了")
	fmt.Println(time.Now())

	<-time.After(3 * time.Second)
	fmt.Println(time.Now())
	fmt.Println("3秒到了")

	// 4.停止定时器
	timer4 := time.NewTimer(4 * time.Second)
	b := timer4.Stop()
	if b {
		fmt.Println("timer4关闭")
	}

	// 5.重置定时器
	timer5 := time.NewTimer(5 * time.Second)
	timer5.Reset(3 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)
	time.Sleep(time.Second * 30)

}

func TestTicker(t *testing.T) {

	// 1.获取ticker对象
	ticker := time.NewTicker(2 * time.Second)
	i := 0
	// 子协程
	go func() {
		for {
			i++
			fmt.Println(<-ticker.C)
			if i == 4 {
				//停止ticker
				ticker.Stop()
			}
		}
	}()
	time.Sleep(time.Second * 20)
}
