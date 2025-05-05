Example usage:

for Set use (keyName, keyValue, ttl)

Set, Get, Delete

	sCache := simplecache.NewScache()
	sCache.Set("name", "Volodya", time.Second*1)
	name, err := sCache.Get("name")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
	sCache.Delete("s2")
	time.Sleep(time.Second * 2)
	name, err = sCache.Get("name")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
