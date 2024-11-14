package main

import (
	"fmt"
	"memory-cache/cache"
)

func main() {
	newCach := cache.New()

	newCach.Set("userId1", 12)
	newCach.Set("userId2", "413")
	newCach.Set("userId3", "413")

	fmt.Println(newCach.Get("userId1"))

	newCach.Delete("userId2")
	fmt.Println(newCach)
}
