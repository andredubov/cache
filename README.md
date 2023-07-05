# Cache

Status of last deployment:<br>
<img src="https://github.com/andredubov/cache/workflows/check/badge.svg?branch=main"><br>

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
	NAME = "Alex"
	AGE  = 27
)

func main() {
	memcache, expiredAt, timeout := cache.NewMemoryCache(), 2*time.Second, 4*time.Second

	memcache.Set(NAME, AGE, expiredAt)

	age, err := memcache.Get(NAME)
	if err != nil {
		fmt.Println(err)
	} else {
		message := fmt.Sprintf("%s is %d years old.\n", NAME, age)
		fmt.Println(message)
	}

	<-time.After(timeout)

	age, err = memcache.Get(NAME)
	if err != nil {
		fmt.Println(err)
	} else {
		message := fmt.Sprintf("%s is %d years old.\n", NAME, age)
		fmt.Println(message)
	}
}
```

