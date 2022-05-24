package lru

import lru "github.com/hashicorp/golang-lru"

var (
	Cache *lru.Cache
	err   error
)

func init() {
	Cache, err = lru.New(5)
	if err != nil {
		panic(err)
	}
}
