package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hhxsv5/go-redis-memory-analysis"
)

//1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
//2. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

var client redis.UniversalClient
var ctx context.Context

const (
	ip   string = "127.0.0.1"
	port uint16 = 6379
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%v:%v", ip, port),
		Password:     "",
		DB:           0,
		PoolSize:     128,
		MinIdleConns: 100,
		MaxRetries:   5,
	})

	ctx = context.Background()
}

func writeData(num int, key, value string) {
	for i := 0; i < num; i++ {
		k := fmt.Sprintf("%s:%v", key, i)
		cmd := client.Set(ctx, k, value, -1)
		if err := cmd.Err(); err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	writeData(5000, "len10_5k", generateValue(10))
	writeData(10000, "len10_10k", generateValue(10))

	writeData(5000, "len1000_5k", generateValue(1000))
	writeData(10000, "len1000_10k", generateValue(1000))

	writeData(5000, "len5000_5k", generateValue(5000))
	writeData(10000, "len5000_10k", generateValue(5000))
	fmt.Println("redis set end")
	analysis()
	fmt.Println("analysis end")
}

func generateValue(size int) string {
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		arr[i] = 'a'
	}
	return string(arr)
}

func analysis() {
	analysis, err := gorma.NewAnalysisConnection(ip, port, "")
	if err != nil {
		fmt.Println("something wrong:", err)
		return
	}
	defer analysis.Close()

	analysis.Start([]string{":"})

	err = analysis.SaveReports("./task5")
	if err == nil {
		fmt.Println("done")
	} else {
		fmt.Println("error:", err)
	}
}
