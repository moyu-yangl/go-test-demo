package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-test-demo/router"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

//var k = make(chan struct{})

func worker(ctx context.Context) {
Loop:
	for i := 0; ; i++ {
		fmt.Println("worker " + strconv.Itoa(i))
		time.Sleep(time.Millisecond * 500)
		wg.Add(1)
		go worker2(ctx)
		select {
		case <-ctx.Done():
			break Loop
		default:
		}
	}
	wg.Done()
	fmt.Println("worker done")
}
func worker2(ctx context.Context) {
Loop:
	for i := 0; ; i++ {
		fmt.Println("worker2 " + strconv.Itoa(i))
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break Loop
		default:
		}
	}
	wg.Done()
	fmt.Println("worker2 done")
}

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8080")

	//logger := log.GetLogger()
	//logger.Error("error")
}
