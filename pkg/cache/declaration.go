package cache

import "github.com/Hrukem/domz2_1"

type Ch struct {
	*domz2_1.Cache
}

func (c *Ch) NewCache() *domz2_1.Cache {
	return domz2_1.New(1)
}
