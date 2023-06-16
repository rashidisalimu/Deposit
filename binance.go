package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
)

var wg sync.WaitGroup

func version2(target int64) {
	var (
		apiKey    = "ezYRTjc2TUGT7IK3h5jcyMy3DBCeTLFiQjBz61deUqcU6Wz0JTyHm0LIjpB2MmOr"
		secretKey = "LI5kmSc8kKUAAUAV9bXbisqWtoA9IBk7vcgV7ns7q9HQxRqGOXjgnPL2O3PiFdSt"
	)
	client := binance.NewClient(apiKey, secretKey)

	start := time.Now()

	var sum float32 = 0
	for {
		server_time, _ := client.NewServerTimeService().Do(context.Background())
		fmt.Println(server_time)
		if target-server_time < 100 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	for {
		order, err := client.NewCreateOrderService().Symbol("IDUSDT").
			Side(binance.SideTypeBuy).Type(binance.OrderTypeLimit).
			TimeInForce(binance.TimeInForceTypeIOC).Quantity("0.2").
			Price("10000").Do(context.Background())
		if err != nil {
			sum += 1
			fmt.Println(err, sum)
			if sum >= 95 {
				fmt.Printf("%s took %v\n", time.Since(start))
				wg.Done()
			}
			// wg.Done()
		} else {
			fmt.Println("Ordertime:", order.TransactTime) // break
			wg.Done()
			continue
		}
	}
}

func main() {
	var (
		apiKey    = "ezYRTjc2TUGT7IK3h5jcyMy3DBCeTLFiQjBz61deUqcU6Wz0JTyHm0LIjpB2MmOr"
		secretKey = "LI5kmSc8kKUAAUAV9bXbisqWtoA9IBk7vcgV7ns7q9HQxRqGOXjgnPL2O3PiFdSt"
	)
	client := binance.NewClient(apiKey, secretKey)

	var target_time int64 = 167948640000

	for {
		server_time, _ := client.NewServerTimeService().Do(context.Background())
		fmt.Println(server_time)
		if target_time-server_time < 300 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("here")
	jobs := 2
	wg.Add(jobs)
	for i := 1; i <= jobs; i++ {
		go version2(target_time)
	}

	wg.Wait()
}
