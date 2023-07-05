# Cache

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
	memcache, timeout := cache.NewMemoryCache(), 3*time.Second

	memcache.Set(NAME, AGE, timeout)

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

