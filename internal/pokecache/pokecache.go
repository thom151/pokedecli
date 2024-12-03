package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entry: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.entry[key] = cacheEntry{time.Now().UTC(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	chEntry, ok := c.entry[key]
	if !ok {
		return nil, false
	}
	return chEntry.val, true
}

func (c *Cache) reapLoop(t time.Duration) {
	ticker := time.NewTicker(t)
	for range ticker.C {
		c.reap(t)
	}
}

func (c *Cache) reap(t time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	timeLimit := time.Now().UTC().Add(-t)
	for key, val := range c.entry {
		if val.createdAt.Before(timeLimit) {
			delete(c.entry, key)
		}
	}

}
