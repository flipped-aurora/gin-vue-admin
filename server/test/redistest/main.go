package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis"
)

// RedisSetPerformanceTest 测试Redis SET操作的性能
func RedisSetPerformanceTest(client *redis.Client, keyPrefix string, numOperations int) (int, error) {
	wg := &sync.WaitGroup{}
	//operations := make(chan struct{}, numOperations)
	var successfulOps int64

	// 设置每个操作的超时时间
	timeout := time.Second

	wg.Add(numOperations)
	// 启动一个goroutine来处理SET操作

	for i := 0; i < numOperations; i++ {
		go func() {
			defer wg.Done()
			key := fmt.Sprintf("%s:%d", keyPrefix, time.Now().UnixNano())
			err := client.Set(key, "value", timeout).Err()
			if err == nil {
				atomic.AddInt64(&successfulOps, 1)
			}
		}()
	}

	//go func() {
	//	for range operations {
	//		go func() {
	//			defer wg.Done()
	//			key := fmt.Sprintf("%s:%d", keyPrefix, time.Now().UnixNano())
	//			err := client.Set(key, "value", timeout).Err()
	//			if err == nil {
	//				atomic.AddInt64(&successfulOps, 1)
	//			}
	//		}()
	//	}
	//}()

	//defer close(operations)
	//// 发送操作请求
	//for i := 0; i < numOperations; i++ {
	//	operations <- struct{}{}
	//}
	//close(operations)

	// 等待所有操作完成
	wg.Wait()

	return int(successfulOps), nil
}

func main() {
	// 创建Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "81.69.19.231:6379",          // Redis地址
		Password: "jhkdjhkjdhsIUTYURTU_sKFAd4", // Redis密码，没有则留空
		DB:       1,                            // 使用默认DB
	})
	//start := time.Now()
	// 测试Redis连接
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("连接Redis失败: %v\n", err)
		return
	}
	fmt.Printf("连接Redis成功: %s\n", pong)
	keyPrefix := "test_key"
	c := 10000
	wg := &sync.WaitGroup{}
	wg.Add(c)
	now := time.Now()
	for i := 0; i < c; i++ {
		i := i
		go func() {
			defer wg.Done()
			k := fmt.Sprintf("%s:%v:%v", keyPrefix, i, time.Now().Unix())
			v := fmt.Sprintf("%v", i)
			set := client.Set(k, v, time.Second*60)
			if set.Err() != nil {
				fmt.Printf("%v: %v\n", i, set.Err())
			}
		}()
	}

	wg.Wait()
	since := time.Since(now)
	fmt.Println(since)
	//// 设置测试参数

	//numOperations := 10 // 假设我们要测试10000次SET操作

	//fmt.Println(keyPrefix, numOperations)
	//result, err := client.Set("key", "value", 0).Result()
	//if err != nil {
	//	panic(result)
	//}

	// 执行性能测试
	//successfulOps, err := RedisSetPerformanceTest(client, keyPrefix, numOperations)
	//if err != nil {
	//	fmt.Printf("性能测试失败: %v\n", err)
	//	return
	//}

	// 计算每秒执行的SET操作数
	//elapsed := time.Since(start)
	//opsPerSecond := float64(successfulOps) / elapsed.Seconds()
	//
	//fmt.Printf("在%v内成功执行了%d次SET操作，平均每秒%.2f次\n", elapsed, successfulOps, opsPerSecond)
}
