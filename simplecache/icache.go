package simplecache

type icache interface {
	Set(key string, value any)
	Get(key string) any
	Delete(key string)
}
