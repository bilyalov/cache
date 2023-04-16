package main

import (
	"fmt"
	"store/store"
	"sync"
	"time"
)

func incTest(s store.Store) {
	fmt.Println("Started: incTest")

	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func(i int, wg *sync.WaitGroup) {
			s.Set("points", i, 0)
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()

	val, _ := s.Get("points")

	fmt.Printf("Last value is %v\n", val)

	fmt.Println("End: incTest")
}

func storeTest(s store.Store) {
	fmt.Println("Started: storeTest")

	fmt.Println("Stored: 'userId = 1' | TTL: 3 seconds")
	s.Set("userId", 1, time.Second*3)

	userId, _ := s.Get("userId")
	fmt.Printf("Instant get userId: %v\n", userId)

	time.Sleep(time.Second * 1)
	userId, _ = s.Get("userId")
	fmt.Printf("Get userId after 1 second: %v\n", userId)

	time.Sleep(time.Second * 3)
	userId, _ = s.Get("userId")
	fmt.Printf("Get userId after 3 seconds: %v\n", userId)

	fmt.Println("End: storeTest")
}

func main() {
	cache := store.New()

	incTest(cache)

	fmt.Println("-------------")

	storeTest(cache)
}
