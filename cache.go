package cache

import (
	"errors"
	"sync"
)

type Cache interface {
	Set(key string, value interface{})
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

func (c *cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.store[key] = value
}

func (c *cache) Get(key string) (interface{}, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if _, exist := c.store[key]; !exist {
		return nil, errors.New("value not found")
	}
	return c.store[key], nil
}

func (c *cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.store, key)
}
