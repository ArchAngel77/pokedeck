
package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries	 map[string]cacheEntry
	mutex 	 sync.Mutex
	duration time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	data	  []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
		duration: interval,
	}
	go c.reapLoop()
	return c
}
func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		data:	   val,
	}
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, ok := c.entries[key]
	    if ok {
		return val.data, true
	}
	    return nil, false
	}
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.duration)
	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.entries {
		if time.Since(entry.createdAt) > c.duration {
			delete(c.entries, key)
		}
	}
	c.mutex.Unlock()
}
}

