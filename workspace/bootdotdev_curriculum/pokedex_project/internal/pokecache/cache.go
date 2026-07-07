package pokecache

import (
	"sync"
	"time"

)

type Cache struct {
	entries map[string]cacheEntery
	mu sync.Mutex
	interval time.Duration
}

type cacheEntery struct {
	createAt time.Time
	val []byte
}
func NewCahse(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntery),
		interval: interval, 
	}
	go cache.reapLoop()
	return  cache 
}
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntery{
		createAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string)([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if cachedVal, ok := c.entries[key]; ok {
		return cachedVal.val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap(){
	c.mu.Lock()
	defer c.mu.Unlock()
	currentTime := time.Now()
	for key, item := range c.entries {
		if currentTime.Sub(item.createAt) > c.interval {
			delete(c.entries, key)
		}
	}
}