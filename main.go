package main

import (
	"fmt"
	"simplecache/basic/simplecache"
)

func main() {
	sCache := simplecache.NewScache()
	sCache.Set("name", "Volodya")
	fmt.Println(sCache.Get("name"))
	sCache.Delete("s2")
	fmt.Println(sCache.Get("name"))
}
