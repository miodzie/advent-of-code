package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSolution1(t *testing.T) {
	f, _ := os.Open("input")
	pairs := ParseInput(f)
	pairs[0].First.dump()
	pairs[0].Second.dump()
	// 5950 is too low.
	// Must be parsing wrong.
	fmt.Println(CountOrderedPackets(pairs))
}

func TestCountOrderedPackets(t *testing.T) {
	f, _ := os.Open("example")
	pairs := ParseInput(f)
	assert.Equal(t, 13, CountOrderedPackets(pairs))
}

func TestComparePackets_ExampleIndex1(t *testing.T) {
	// [1,1,3,1,1] vs [1,1,5,1,1]
	left := ParsePacket("[1,1,3,1,1]")
	right := ParsePacket("[1,1,5,1,1]")

	assert.Equal(t, -1, IsCorrectOrder(left.children, right.children))
}

func TestComparePackets_ExampleIndex2(t *testing.T) {
	first := ParsePacket("[[1],[2,3,4]]")
	second := ParsePacket("[[1],4]")

	assert.Equal(t, -1, IsCorrectOrder(first.children, second.children))
}

func TestComparePackets_ExampleIndex3(t *testing.T) {
	// [9] vs [[8,7,6]]
	first := List([]int{9})
	second := &Packet{children: []*Packet{List([]int{8, 7, 6})}}

	assert.Equal(t, 1, IsCorrectOrder(first.children, second.children))
}

func TestComparePackets_ExampleIndex5(t *testing.T) {
	first := ParsePacket("[7,7,7,7]")
	second := ParsePacket("[7,7,7]")
	assert.Equal(t, 1, IsCorrectOrder(first.children, second.children))
}

func TestComparePackets_ExampleIndex7(t *testing.T) {
	first := ParsePacket("[[[]]]")
	second := ParsePacket("[[]]")
	assert.Equal(t, 1, IsCorrectOrder(first.children, second.children))
}

func TestComparePackets_ExampleIndex8(t *testing.T) {
	first := ParsePacket("[1,[2,[3,[4,[5,6,7]]]],8,9]")
	second := ParsePacket("[1,[2,[3,[4,[5,6,0]]]],8,9]")
	assert.Equal(t, 1, IsCorrectOrder(first.children, second.children))
}

func TestComparePackets_second_is_empty(t *testing.T) {
	// [3,4] and nil
	threeFour := ParsePacket("[3,4]")
	empty := ParsePacket("[]")
	assert.Equal(t, 1, IsCorrectOrder(threeFour.children, empty.children))
}
