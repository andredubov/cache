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
	KEY = "Alex"
	VALUE  = 27
)

func main() {
	memcache, timeout := cache.NewMemoryCache(), 3 * time.Second

	memcache.Set(KEY, VALUE, timeout)

	value, err := cache.Get(KEY)
	if err != nil {
		fmt.Println(err)
	} else {
		message := fmt.Sprintf("%s is %d years old.\n", KEY, value)
		fmt.Println(message)
	}

	<-time.After(timeout)

	value, err = cache.Get(KEY)
	if err != nil {
		fmt.Println(err)
	} else {
		message := fmt.Sprintf("%s is %d years old.\n", KEY, value)
		fmt.Println(message)
	}
}
```

