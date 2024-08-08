package main

import (
	"fmt"

	"github.com/oarkflow/xsync/cache"
)

func main() {
	c := cache.New[string, string](cache.WithCapacity(0))
	result, err := c.Fetch("key", func() (string, error) {
		return "value", nil
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
