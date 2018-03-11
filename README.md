# Bloom filter

[Standard Bloom filter](https://en.wikipedia.org/wiki/Bloom_filter) implementation in Go.  
[MurmurHash3](https://en.wikipedia.org/wiki/MurmurHash) is used for hashing.

```go
// First parameter is the maximum size.
// The second is the false positive probability value that is acceptable for you.
// Based on these two values, optimal values are calculated for the hash count & bit array size.
bf := bloomfilter.New(1000, 0.025)
bf.Add("foo")
bf.Add("bar")

// "foo" has been found!
if bf.Check("foo") {
  fmt.Println("Found!")
}
```

```
BenchmarkAdd10000-8     	10000000	       147 ns/op
BenchmarkCheck10000-8   	10000000	       152 ns/op
```