package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
    entries map[string]CacheEntry
    mu      sync.Mutex
    interval time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val []byte
}

func (c *Cache) reapLoop() {
    ticker := time.NewTicker(c.interval)
    for range ticker.C {
        c.mu.Lock()
        for k, v := range c.entries {
            if time.Since(v.createdAt) > c.interval {
                delete(c.entries, k)
            }
        }
        c.mu.Unlock()
    }
}

func NewCache(interval time.Duration) *Cache {
    cacheMap := make(map[string]CacheEntry)
    c := &Cache{
        entries:  cacheMap,
        interval: interval,
    }

    // Start reap loop in background
    go c.reapLoop()

    return c
}


func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val:      val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, exists := c.entries[key]; exists {
		return entry.val, true
	}
	return nil, false
}