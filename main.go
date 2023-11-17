package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Tester interface {
	say(string) string
}

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func (i Item) say(str string) string {
	println(i.inStock)
	if str == "s" {
		return "ok"
	} else {
		return "error"
	}
}

type Observer struct {
	name string
}

func (o Observer) update(name string) {
	o.name = name
}
func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}
func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}
func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

var wg sync.WaitGroup
var lock sync.RWMutex
var x = 0

func main() {
	//r := gin.Default()
	//router.InitRouter(r)
	//r.Run(":8080")

	read := func(k *int) {
		lock.RLock()
		fmt.Println("read: ", *k)
		time.Sleep(time.Millisecond * 500)
		lock.RUnlock()
		wg.Done()
	}
	write := func(k *int) {
		lock.Lock()
		*k = *k + 1
		fmt.Println("write: ", *k)
		time.Sleep(time.Millisecond * 1000)
		lock.Unlock()
		wg.Done()
	}

	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go write(&x)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 500; i++ {
			wg.Add(1)
			go read(&x)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("end")
}

func synck() {
	var (
		stop = make(chan struct{})
		done = make(chan struct{})
	)
	delay := 100 * time.Millisecond
	var i = 0
	go func() {
		defer close(done)
		tiker := time.NewTicker(delay)
		defer tiker.Stop()
		for {
			select {
			case <-stop:
				return
			case <-tiker.C:
				fmt.Println("i : " + strconv.Itoa(i))
				i++
			}
		}

	}()
	for {
		fmt.Println("please inout a number")
		var input int
		n, err := fmt.Scan(&input)
		if err != nil || n <= 0 || input != 0 {
			continue
		}
		close(stop)
		fmt.Println(<-done)
		break
	}
}
