package simplecache

import (
	"errors"
	"sync"
	"time"
)

type cacheEl struct {
	expiredTime int64
	value       any
}

type Scache struct {
	// mutex for avoid race condition
	mu       *sync.RWMutex
	elements map[string]cacheEl
}

func NewScache() *Scache {
	cache := Scache{
		elements: make(map[string]cacheEl),
		mu:       new(sync.RWMutex),
	}

	go cache.startBackgroundTask()

	return &cache
}

func (cache *Scache) startBackgroundTask() {
	for {
		go cache.deleteValueByExpiredTime()
	}
}

func (cache *Scache) deleteValueByExpiredTime() {
	for key, value := range cache.elements {
		if value.expiredTime > time.Now().UnixNano() {
			cache.Delete(key)
		}
	}
}

func (cache *Scache) Set(key string, value any, ttl time.Duration) {
	//TODO можно ли не блокировать всю мапу здесь и при чтении?
	if cache != nil {
		cache.mu.Lock()
		expiredTime := time.Now().Add(ttl).UnixNano()
		cache.elements[key] = cacheEl{expiredTime, value}
		cache.mu.Unlock()
	}
}

func (cache *Scache) Get(key string) (any, error) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	value := cache.elements[key].value
	if value == nil || cache.elements[key].expiredTime < time.Now().UnixNano() {
		return 0, errors.New("value not exist")
	}

	return value, nil
}

func (cache *Scache) Delete(key string) {
	cache.mu.Lock()
	delete(cache.elements, key)
	cache.mu.Unlock()
}
