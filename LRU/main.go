package main

import lrucache "lru/lru"

func main() {
	cache := lrucache.NewLRUCache[string, string](3)

	cache.Put("1", "First Elemente")
	cache.Put("2", "Second Elemente")
	cache.Put("3", "Third Elemente")

	cache.Get("3")
	cache.Put("4", "Forth Elemente")
	cache.GetCache()
}
