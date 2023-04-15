package main

import (
	"fmt"
	"store/store"
)

func main() {
	cache := store.New()

	cache.Set("userId", 1)

	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")

	fmt.Println(cache)
}
