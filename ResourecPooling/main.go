package main

import (
	"fmt"
	"resource-pooling/pool"
	"time"
)

func main() {
	fmt.Println("hello")
	pool := pool.NewResourcePool(2)
	resorce1 := pool.Acquire()
	resorce2 := pool.Acquire()
	fmt.Printf("Acquired Resource 1 ID=%d, Data=%s\n", resorce1.ID, resorce1.Data)
	fmt.Printf("Acquired Resource 2 ID=%d, Data=%s\n", resorce2.ID, resorce2.Data)

	go func() {
		resourec3 := pool.Acquire()
		fmt.Printf("Acquired Resource 3 ID=%d, Data=%s\n", resourec3.ID, resourec3.Data)
	}()

	pool.Release(resorce1)

	time.Sleep(1 * time.Second)
	return
}
