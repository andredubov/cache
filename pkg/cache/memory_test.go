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
	memorycache, timeout := cache.NewMemoryCache(), 3*time.Second

	memorycache.Set(NAME, AGE, timeout)

	age, err := memorycache.Get(NAME)
	if age == nil || err == cache.ErrItemNotFound {
		t.Error(err)
	}

	<-time.After(timeout)

	age, err = memorycache.Get(NAME)
	if age == nil && err != cache.ErrItemNotFound {
		t.Error("an item shouldn't found.")
	}
}
