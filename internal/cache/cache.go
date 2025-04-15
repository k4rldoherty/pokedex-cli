package cache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheEntry
	mu   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		map[string]cacheEntry{},
		&sync.Mutex{},
	}
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			c.readLoop(interval)
		}
	}()
	return c
}

func (c *Cache) Add(k string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[k] = cacheEntry{createdAt: time.Now(), val: val}
	return nil
}

func (c *Cache) Get(k string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, ok := c.data[k]
	if !ok {
		return nil, false
	}
	return data.val, true
}

func (c *Cache) readLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	for k, v := range c.data {
		if now.Sub(v.createdAt) > interval {
			delete(c.data, k)
		}
	}
}
