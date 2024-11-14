package main

import (
	"fmt"
	"memory-cache/cache"
)

func main() {
	newCach := cache.New()

	newCach.Set("userId1", 12)
	newCach.Set("userId2", "441")
	newCach.Set("userId3", "413")

	fmt.Println(newCach.Get("userId1"))

	newCach.Delete("91992913213123")
	err := newCach.Delete("b")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newCach)
}
