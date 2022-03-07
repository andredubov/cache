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
	cache := cache.New()
	cache.Set(NAME, AGE)
	age, err := cache.Get(NAME)
	if err != nil {
		fmt.Println(err)
	} else {
		message := fmt.Sprintf("%s is %d years old.\n", NAME, age)
		fmt.Println(message)
	}

	cache.Delete(NAME)
	age, err = cache.Get(NAME)
	if err != nil {
		fmt.Println(err)
	} else {
		message := fmt.Sprintf("%s is %d years old.\n", NAME, age)
		fmt.Println(message)
	}
}
```

