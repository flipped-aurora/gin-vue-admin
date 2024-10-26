package myutil

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os/exec"
	"time"
)

// CheckAndStartRedis 检查 Redis 是否已启动，如果没有则启动
func CheckAndStartRedis() error {
	// 创建 Redis 客户端连接
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址和端口号
		Password: "",               // 如果 Redis 服务器设置了密码，请填写密码
		DB:       0,                // 使用的数据库编号，默认为 0
	})

	// 尝试连接 Redis 服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err == nil {
		// 如果连接成功，关闭 Redis 客户端连接并返回
		err = client.Close()
		if err != nil {
			return fmt.Errorf("error closing Redis connection: %v", err)
		}
		return nil
	}

	// 如果连接失败，则启动 Redis 服务器
	fmt.Println("Redis server is not running. Starting Redis server...")

	// 启动 Redis 服务器
	err = startRedisServer()
	if err != nil {
		return fmt.Errorf("error starting Redis server: %v", err)
	}

	fmt.Println("Redis server started successfully.")

	return nil
}

// startRedisServer 启动 Redis 服务器
func startRedisServer() error {
	// 指定 Redis 服务器的可执行文件路径
	redisPath := "D:/deskapp/Redis-x64-5.0.14.1/redis-server" // 这是 Redis 默认安装路径，请根据实际情况修改

	// 创建命令来启动 Redis 服务器
	cmd := exec.Command(redisPath)

	// 启动 Redis 服务器
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("error starting Redis server: %v", err)
	}

	// 等待 Redis 服务器启动完成
	time.Sleep(1 * time.Second)

	return nil
}
