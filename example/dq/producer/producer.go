package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tal-tech/go-queue/dq"
)

func main() {
	producer := dq.NewProducer([]dq.Beanstalk{
		{
			Endpoint: "172.18.133.218:11301",
			Tube:     "tube",
		},
		{
			Endpoint: "172.18.133.218:11302",
			Tube:     "tube",
		},
	})
	for i := 1000; i < 1005; i++ {
		dada, err := producer.Delay([]byte(strconv.Itoa(i)), time.Second*5)
		if err != nil {
			fmt.Println(dada, err)
		}

		if i%2 == 1 {
			err = producer.Revoke(dada)
			fmt.Println(err)

		}
	}
}
