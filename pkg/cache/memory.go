package cache

import (
	"errors"
	"sync"
	"time"
)

var ErrItemNotFound = errors.New("item not found")

type item struct {
	value     any
	createdAt int64
	ttl       int64
}

type MemoryCache struct {
	sync.RWMutex
	cache map[any]*item
}

// NewMemoryCache uses map to store key:value data in-memory
func NewMemoryCache() *MemoryCache {
	c := &MemoryCache{cache: make(map[any]*item)}
	go c.setTtlTimer()

	return c
}

func (c *MemoryCache) Set(key, value any, ttl int64) error {
	c.Lock()
	c.cache[key] = &item{
		value:     value,
		createdAt: time.Now().Unix(),
		ttl:       ttl,
	}
	c.Unlock()

	return nil
}

func (c *MemoryCache) Get(key any) (any, error) {
	c.RLock()
	item, ex := c.cache[key]
	c.RUnlock()

	if !ex {
		return nil, ErrItemNotFound
	}

	return item.value, nil
}

func (c *MemoryCache) setTtlTimer() {
	for {
		c.Lock()
		for k, v := range c.cache {
			if time.Now().Unix()-v.createdAt > v.ttl {
				delete(c.cache, k)
			}
		}
		c.Unlock()

		<-time.After(time.Second)
	}
}
