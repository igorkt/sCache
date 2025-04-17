package simplecache

type Scache struct {
	elements map[string]any
}

func NewScache() Scache {
	return Scache{
		elements: make(map[string]any),
	}
}

func (cache Scache) Set(key string, value any) {
	cache.elements[key] = value
}

func (cache Scache) Get(key string) any {
	return cache.elements[key]
}

func (cache Scache) Delete(key string) {
	delete(cache.elements, key)
}
