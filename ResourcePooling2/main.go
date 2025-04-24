package main

import (
	"fmt"
	"resource-pooling/pool"
	"time"
)

func main() {
	fmt.Println("I am from main")
	resourcePool := pool.NewResourcePool(2)

	acqureResource := func(id int) {
		fmt.Printf("Goroutine %d trying to acquire resource \n", id)
		res, err := resourcePool.Acquire()
		if err != nil {
			fmt.Println("Failed: GoRoutine ", id)
			fmt.Println("error: ", err)
			return
		}
		rID, rData := res.GetDataID()
		fmt.Printf("Acquired: Goroutine  %d, resource Id %d, Data %s \n", id, rID, rData)
		time.Sleep(3 * time.Second)

		resourcePool.Release(res)
		fmt.Printf("Realeased: Goroutine %d, Realeased resource Id %d, Data %s \n", id, rID, rData)
	}

	for i := 0; i <= 4; i++ {
		go acqureResource(i)
	}

	time.Sleep(20 * time.Second)

	// r1, _ := resourcePool.Acquire()
	// r2, _ := resourcePool.Acquire()

	// go func() {
	// 	r3, err := resourcePool.Acquire()
	// 	if err != nil {
	// 		fmt.Println("error while acquiring resource")
	// 	}
	// 	resourcePool.Release(r3)
	// }()

	// time.Sleep(7 * time.Second)
	// resourcePool.Release(r1)
	// resourcePool.Release(r2)

}
