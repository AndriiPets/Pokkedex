package cache

import (
	"sync"
	"time"
)

type Cache struct {
	Map map[string]CacheEntry
	Mux *sync.Mutex
}

type CacheEntry struct {
	CreatedAt time.Time
	Data      []byte
}

func NewCache(t time.Duration) Cache {
	cMap := make(map[string]CacheEntry)
	cache := Cache{
		Map: cMap,
		Mux: &sync.Mutex{},
	}
	go cache.ReapLoop(t)

	return cache
}

func (c *Cache) Add(k string, v []byte) {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	if _, ok := c.Map[k]; ok {
		c.Mux.Unlock()
		return
	}

	data := CacheEntry{
		CreatedAt: time.Now().UTC(),
		Data:      v,
	}

	c.Map[k] = data
}

func (c *Cache) Get(k string) (v []byte, ok bool) {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	val, ok := c.Map[k]

	return val.Data, ok
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.Reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) Reap(now time.Time, interval time.Duration) {
	c.Mux.Lock()
	defer c.Mux.Unlock()

	for k, v := range c.Map {
		if v.CreatedAt.Before(now.Add(-interval)) {
			delete(c.Map, k)
		}
	}
}
