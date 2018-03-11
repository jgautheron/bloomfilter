package bloomfilter

import (
	"math"

	"github.com/spaolacci/murmur3"
)

type Filter struct {
	fpp  float64
	size uint32
	hc   uint32
	bs   []byte
}

// New instantiates the filter with the given bit array size and the wanted
// false positive probability.
func New(cn, fpp float64) *Filter {
	s := size(cn, fpp)
	h := hc(s, cn)
	b := make([]byte, s)
	return &Filter{fpp, s, h, b}
}

// Add a new item to the filter
func (f *Filter) Add(item []byte) {
	var i uint32
	for i = 0; i < f.hc; i++ {
		digest := murmur3.Sum32WithSeed(item, i) % f.size
		f.bs[digest] = 1
	}
}

// Check if the given item has been added previously.
func (f *Filter) Check(item []byte) bool {
	var i uint32
	for i = 0; i < f.hc; i++ {
		digest := murmur3.Sum32WithSeed(item, i) % f.size
		if f.bs[digest] == 0 {
			return false
		}
	}
	return true
}

// size determines the optimal bit array size based on the given item count and fpp.
func size(cn, fpp float64) uint32 {
	return uint32(-(cn * math.Log(fpp)) / (math.Pow(math.Log(2), 2)))
}

// hc determines the optimal hascount based on the given size.
func hc(s uint32, cn float64) uint32 {
	return uint32((float64(s) / cn) * math.Log(2))
}
