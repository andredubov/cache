package cache_test

import (
	"testing"
	"time"

	"github.com/andredubov/cache"
)

const (
	KEY   = "Alex"
	VALUE = 27
)

func TestMemoryCache(t *testing.T) {
	memorycache, timeout := cache.NewMemoryCache(), 3*time.Second

	memorycache.Set(KEY, VALUE, timeout)

	value, err := memorycache.Get(KEY)
	if value == nil || err == cache.ErrItemNotFound {
		t.Error(err)
	}

	<-time.After(timeout)

	value, err = memorycache.Get(KEY)
	if value == nil && err != cache.ErrItemNotFound {
		t.Error("an item shouldn't found.")
	}
}
