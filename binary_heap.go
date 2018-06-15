package main

import (
	"bytes"
	"fmt"
	"github.com/disiqueira/gotree"
	"strconv"
)

type BinaryHeap struct {
	arr []int
}

func NewBinaryHeap() *BinaryHeap {
	// leaving first position unused so bit shifting works well
	b := BinaryHeap{arr: []int{-1}}
	return &b
}

func (b *BinaryHeap) Add(x int) {
	b.arr = append(b.arr, x)
	ix := len(b.arr) - 1
	for ix > 1 && x < b.arr[ix>>1] {
		b.arr[ix] = b.arr[ix>>1]
		ix = ix >> 1
	}
	b.arr[ix] = x
}

func (b *BinaryHeap) heapify(ix int) {
	// get children
	lix := (ix << 1)
	rix := (ix << 1) + 1
	// for now, least is the current node
	least := ix
	// check if left node is least
	if lix < len(b.arr) && b.arr[lix] < b.arr[least] {
		least = lix
	}
	// check if right node is least
	if rix < len(b.arr) && b.arr[rix] < b.arr[least] {
		least = rix
	}
	// if one of children was least, swap and heapify from that node
	if least != ix {
		b.arr[ix], b.arr[least] = b.arr[least], b.arr[ix]
		b.heapify(least)
	}
}

func (b *BinaryHeap) Pop() int {
	// get root elem and replace with last element (shrink slice)
	x := b.arr[1]
	last_ix := len(b.arr) - 1
	b.arr[1] = b.arr[last_ix]
	b.arr = b.arr[:last_ix]
	// bubble element down to its place
	b.heapify(1)
	return x
}

func (b *BinaryHeap) Len() int {
	return len(b.arr) - 1
}

func (b *BinaryHeap) String() string {
	// for building string
	var buffer bytes.Buffer
	// print array
	buffer.WriteString(fmt.Sprintf("BHeap{%v}", b.arr[1:]))
	// if elements, print tree
	if len(b.arr) > 1 {
		buffer.WriteString("\n")
		// build tree using gotree package by descending recursively
		var addChildren func(node gotree.Tree, ix int)
		addChildren = func(node gotree.Tree, ix int) {
			rix := (ix << 1) + 1
			if rix < len(b.arr) {
				r := node.Add(strconv.Itoa(b.arr[rix]))
				addChildren(r, rix)
			}
			lix := (ix << 1)
			if lix < len(b.arr) {
				l := node.Add(strconv.Itoa(b.arr[lix]))
				addChildren(l, lix)
			}
		}
		tree := gotree.New(strconv.Itoa(b.arr[1]))
		addChildren(tree, 1)
		buffer.WriteString(tree.Print())
	}
	return buffer.String()
}
