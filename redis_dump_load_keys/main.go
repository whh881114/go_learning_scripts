package main

// redis_dump_load_keys，连入redis或redis集群并发取key，写到指定的key文件；恢复key时，读取指定的文件导入到redis中。
// 以上是一个初步的思路，实现逻辑还有待确认。

import (
	"github.com/go-redis/redis/v8"
)

func main() {
	opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)

}
