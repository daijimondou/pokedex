package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(t time.Duration) Cache {
	newCache := Cache{
		entry: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	go newCache.ReapLoop(t)
	return newCache
}

func (cache *Cache) ReapLoop(t time.Duration) {
	ticker := time.NewTicker(t)
	for range ticker.C {
		cache.mu.Lock()
		for k, v := range cache.entry {
			if t < time.Since(v.createdAt) {
				delete(cache.entry, k)
			}
		}
		cache.mu.Unlock()
	}
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.entry[key] = entry
	cache.mu.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	e, ok := cache.entry[key]
	if !ok {
		return nil, false
	}
	return e.val, true
}
