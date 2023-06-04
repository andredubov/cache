# Cache

Simple in-memory cache

See it in action:

## Example

```go
package main

import (
	"fmt"
	"github.com/andredubov/cache"
)

const (
	NAME = "Alex"
	AGE  = 27
)

func main() {
	memcache, expiredAt := cache.NewMemoryCache(), 3 * time.Second

	memcache.Set(NAME, AGE, expiredAt)

	age, err := cache.Get(NAME)
	if err != nil {
		fmt.Println(err)
	} else {
		message := fmt.Sprintf("%s is %d years old.\n", NAME, age)
		fmt.Println(message)
	}

	<-time.After(expiredAt)

	age, err = cache.Get(NAME)
	if err != nil {
		fmt.Println(err)
	} else {
		message := fmt.Sprintf("%s is %d years old.\n", NAME, age)
		fmt.Println(message)
	}
}
```

