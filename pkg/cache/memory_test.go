package cache_test

import (
	"testing"
	"time"

	"github.com/andredubov/cache/pkg/cache"
)

const (
	NAME = "Alex"
	AGE  = 27
)

func TestMemoryCache(t *testing.T) {
	memcache, expiredAt, timeout := cache.NewMemoryCache(), 2*time.Second, 4*time.Second

	memcache.Set(NAME, AGE, expiredAt)

	age, err := memcache.Get(NAME)
	if age == nil || err == cache.ErrItemNotFound {
		t.Error(err)
	}

	<-time.After(timeout)

	age, err = memcache.Get(NAME)
	if age == nil && err != cache.ErrItemNotFound {
		t.Error("an item shouldn't found in the cache")
	}
}
