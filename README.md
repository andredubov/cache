# Cache

<img src="https://github.com/andredubov/cache/workflows/test/badge.svg?branch=main">

Simple in-memory cache

See it in action:

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/andredubov/cache/pkg/cache"
)

const (
	KEY = "Alex"
	VALUE  = 27
)

func main() {
	memcache, expiredAt, timeout := cache.NewMemoryCache(), 2*time.Second, 4*time.Second

	memcache.Set(KEY, VALUE, timeout)

	age, err := memcache.Get(NAME)
	if err != nil {
		fmt.Println(err)
	} else {
		message := fmt.Sprintf("%s is %d years old.\n", KEY, value)
		fmt.Println(message)
	}

	<-time.After(timeout)

	age, err = memcache.Get(NAME)
	if err != nil {
		fmt.Println(err)
	} else {
		message := fmt.Sprintf("%s is %d years old.\n", KEY, value)
		fmt.Println(message)
	}
}
```

