package main

import (
	"fmt"

	"github.com/swamp0407/gopher_dojo/cache"
)

func main() {
	cache := cache.NewCacheSlice(10)
	cache.Set(1, 1)
	cache.Set(2, 2)
	cache.Set(3, 3)
	cache.Set(4, 4)
	fmt.Println(cache.Get(1))
	cache.Delete(1)
}
