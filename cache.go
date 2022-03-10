package cache

import (
	"sync"
	"time"

	"github.com/pkg/errors"
)

type Cache interface {
	Set(key string, value interface{}, ttl time.Duration)
	Get(key string) (interface{}, error)
	Delete(key string)
}

type cache struct {
	store map[string]interface{}
	mutex *sync.RWMutex
}

func New() Cache {
	return &cache{
		store: make(map[string]interface{}),
		mutex: new(sync.RWMutex),
	}
}

func (c *cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mutex.Lock()
	c.store[key] = value
	c.mutex.Unlock()

	go func(ttl time.Duration) {
		time.Sleep(ttl)
		c.Delete(key)
	}(ttl)
}

func (c *cache) Get(key string) (interface{}, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if _, exist := c.store[key]; !exist {
		return nil, errors.Errorf("there is no item that has the key [%s]", key)
	}
	return c.store[key], nil
}

func (c *cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.store, key)
}
