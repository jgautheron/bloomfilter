package bloomfilter_test

import (
	"math/rand"
	"testing"
	"time"

	bloomfilter "github.com/jgautheron/bloomfilter"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestBloomFilter100(t *testing.T) {
	n := 100

	d := data(n)
	bf := bloomfilter.New(float64(n), 0.025)

	for _, k := range d {
		bf.Add(k)
	}

	for _, k := range d {
		if !bf.Check(k) {
			t.Errorf("Expected %s to be found", k)
		}
	}

	if bf.Check([]byte("foo123")) {
		t.Errorf("Expected %s to be NOT found", "foo123")
	}
}

func BenchmarkAdd100(b *testing.B) {
	benchmarkAdd(100, b)
}

func BenchmarkAdd10000(b *testing.B) {
	benchmarkAdd(10000, b)
}

func BenchmarkCheck100(b *testing.B) {
	benchmarkCheck(100, b)
}

func BenchmarkCheck10000(b *testing.B) {
	benchmarkCheck(10000, b)
}

func benchmarkCheck(i int, b *testing.B) {
	d := data(i)
	bf := bloomfilter.New(float64(i), 0.025)
	for _, k := range d {
		bf.Add(k)
	}

	for n := 0; n < b.N; n++ {
		bf.Check(d[rand.Intn(len(d))])
	}
}

func benchmarkAdd(i int, b *testing.B) {
	d := data(i)
	bf := bloomfilter.New(float64(i), 0.025)
	for n := 0; n < b.N; n++ {
		bf.Add(d[rand.Intn(len(d))])
	}
}

func data(length int) (out [][]byte) {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	n := 10
	for i := 0; i < length; i++ {
		b := make([]byte, n)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		out = append(out, b)
	}
	return
}
