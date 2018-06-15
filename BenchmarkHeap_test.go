package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkHeapAdd32Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 32; i++ {
		h.Add(i)
	}
}

func BenchmarkHeapPop32Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 32; i++ {
		h.Add(i)
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd32Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 32; i++ {
		h.Add(rand.Intn(32))
	}
}

func BenchmarkHeapPop32Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 32; i++ {
		h.Add(rand.Intn(32))
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd64Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 64; i++ {
		h.Add(i)
	}
}

func BenchmarkHeapPop64Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 64; i++ {
		h.Add(i)
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd64Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 64; i++ {
		h.Add(rand.Intn(64))
	}
}

func BenchmarkHeapPop64Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 64; i++ {
		h.Add(rand.Intn(64))
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd128Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 128; i++ {
		h.Add(i)
	}
}

func BenchmarkHeapPop128Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 128; i++ {
		h.Add(i)
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd128Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 128; i++ {
		h.Add(rand.Intn(128))
	}
}

func BenchmarkHeapPop128Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 128; i++ {
		h.Add(rand.Intn(128))
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd256Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 256; i++ {
		h.Add(i)
	}
}

func BenchmarkHeapPop256Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 256; i++ {
		h.Add(i)
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd256Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 256; i++ {
		h.Add(rand.Intn(256))
	}
}

func BenchmarkHeapPop256Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 256; i++ {
		h.Add(rand.Intn(256))
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd512Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 512; i++ {
		h.Add(i)
	}
}

func BenchmarkHeapPop512Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 512; i++ {
		h.Add(i)
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd512Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 512; i++ {
		h.Add(rand.Intn(512))
	}
}

func BenchmarkHeapPop512Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 512; i++ {
		h.Add(rand.Intn(512))
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd1024Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 1024; i++ {
		h.Add(i)
	}
}

func BenchmarkHeapPop1024Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 1024; i++ {
		h.Add(i)
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd1024Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 1024; i++ {
		h.Add(rand.Intn(1024))
	}
}

func BenchmarkHeapPop1024Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 1024; i++ {
		h.Add(rand.Intn(1024))
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd2048Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 2048; i++ {
		h.Add(i)
	}
}

func BenchmarkHeapPop2048Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 2048; i++ {
		h.Add(i)
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd2048Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 2048; i++ {
		h.Add(rand.Intn(2048))
	}
}

func BenchmarkHeapPop2048Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 2048; i++ {
		h.Add(rand.Intn(2048))
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd4096Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 4096; i++ {
		h.Add(i)
	}
}

func BenchmarkHeapPop4096Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 4096; i++ {
		h.Add(i)
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd4096Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 4096; i++ {
		h.Add(rand.Intn(4096))
	}
}

func BenchmarkHeapPop4096Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 4096; i++ {
		h.Add(rand.Intn(4096))
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd409632Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 4096*32; i++ {
		h.Add(i)
	}
}

func BenchmarkHeapPop409632Ordered(b *testing.B) {
	h := NewBinaryHeap()
	for i := 0; i < 4096*32; i++ {
		h.Add(i)
	}
	b.ResetTimer()
	for h.Len() > 0 {
		h.Pop()
	}
}

func BenchmarkHeapAdd409632Random(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	h := NewBinaryHeap()
	for i := 0; i < 4096*32; i++ {
		h.Add(rand.Intn(4096 * 32))
	}
}
