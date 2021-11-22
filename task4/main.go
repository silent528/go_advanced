package main

import (
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	const (
		seconds = 1
		quota   = 5
		total   = 100 // 模拟 1w 次请求
	)
	l := NewPeriodLimit(seconds, quota, rdb)
	var allowed, hitQuota, overQuota int
	for i := 0; i < total; i++ {
		val, err := l.Take("first")
		if err != nil {
			log.Fatal(err)
		}
		switch val {
		case Allowed:
			allowed++
		case HitQuota:
			hitQuota++
		case OverQuota:
			overQuota++
		default:
			log.Fatal("unknown status")
		}
		time.Sleep(time.Millisecond * 100)
	}
	log.Println(allowed)
	log.Println(hitQuota)
	log.Println(overQuota)
}
