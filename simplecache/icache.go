package simplecache

import "time"

type icache interface {
	Set(key string, value any, ttl time.Duration)
	Get(key string) (any, error)
	Delete(key string)
}
