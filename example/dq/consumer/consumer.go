package main

import (
	"fmt"

	"github.com/tal-tech/go-queue/dq"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

func main() {
	consumer := dq.NewConsumer(dq.DqConf{
		Beanstalks: []dq.Beanstalk{
			{
				Endpoint: "172.18.133.218:11301",
				Tube:     "tube",
			},
			{
				Endpoint: "172.18.133.218:11302",
				Tube:     "tube",
			},
		},
		Redis: redis.RedisConf{
			Host: "localhost:6379",
			Type: redis.NodeType,
		},
	})
	consumer.Consume(func(body []byte) {
		fmt.Println(string(body))
	})
}
