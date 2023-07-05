package main

import (
	"fmt"
	"time"

	"github.com/andredubov/cache/pkg/cache"
)

const (
	NAME = "Alex"
	AGE  = 27
)

func main() {
	memcache, timeout := cache.NewMemoryCache(), 2*time.Second

	memcache.Set(NAME, AGE, timeout)

	age, err := memcache.Get(NAME)
	if err != nil {
		fmt.Println(err)
	} else {
		message := fmt.Sprintf("%s is %d years old.\n", NAME, age)
		fmt.Println(message)
	}

	<-time.After(4 * time.Second)

	age, err = memcache.Get(NAME)
	if err != nil {
		fmt.Println(err)
	} else {
		message := fmt.Sprintf("%s is %d years old.\n", NAME, age)
		fmt.Println(message)
	}
}
