package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]CacheEntry
	mux *sync.Mutex
}

type CacheEntry struct {
	val       []byte //http entries are just a bunch of bytes
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]CacheEntry),
		mux: &sync.Mutex{},
	}
	go c.PurgeRepeat(interval)
	//if not spawned in another goroutine will halt execution
	return c
}

func (c *Cache) AddCache(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = CacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	CacheE, ok := c.cache[key]
	return CacheE.val, ok
}

func (c *Cache) PurgeRepeat(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		//this code runs every interval minutes
		c.Purge(interval)
	}
}

func (c *Cache) Purge(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	timePassed := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(timePassed) {
			delete(c.cache, k)
		}
	}
}