package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	b := NewBinaryHeap()
	for i := 0; i < 10; i++ {
		b.Add(i * i)
	}
	fmt.Println(b)
}
