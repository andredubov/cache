package cache_test

import (
	"testing"
	"time"

	"github.com/andredubov/cache"
)

const (
	NAME = "Alex"
	AGE  = 27
)

func TestMemoryCache(t *testing.T) {
	memorycache, expiredAt := cache.NewMemoryCache(), 3*time.Second

	memorycache.Set(NAME, AGE, expiredAt)

	age, err := memorycache.Get(NAME)
	if age == nil || err == cache.ErrItemNotFound {
		t.Error(err)
	}

	<-time.After(expiredAt)

	age, err = memorycache.Get(NAME)
	if age == nil && err != cache.ErrItemNotFound {
		t.Error("an item shouldn't found.")
	}
}
