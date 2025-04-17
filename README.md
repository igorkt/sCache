Example usage:

Set, Get, Delete

	sCache := simplecache.NewScache()
	sCache.Set("name", "Volodya")
	fmt.Println(sCache.Get("name"))
	sCache.Delete("s2")
	fmt.Println(sCache.Get("name"))
