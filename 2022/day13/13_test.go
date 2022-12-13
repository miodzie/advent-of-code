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
	// 5950 is too low.
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

func TestComparePackets5(t *testing.T) {
	first := ParsePacket("[7,7,7,7]")
	second := ParsePacket("[7,7,7]")
	assert.Equal(t, 1, IsCorrectOrder(first.children, second.children))
}

//func TestComparePackets_ExampleIndex2(t *testing.T) {
//	first := ParsePacket("[[1],[2,3,4]]")
//	second := ParsePacket("[[1],4]")
//
//	assert.True(t, IsCorrectOrder(first, second))
//}
//
//
//func TestComparePackets7(t *testing.T) {
//	first := ParsePacket("[[[]]]")
//	second := ParsePacket("[[]]")
//	assert.False(t, IsCorrectOrder(first, second))
//}
//
//func TestComparePackets3(t *testing.T) {
//	// [1,[2,[3,[4,[5,6,7]]]],8,9] vs [1,[2,[3,[4,[5,6,0]]]],8,9]
//	first := ParsePacket("[1,[2,[3,[4,[5,6,7]]]],8,9]")
//	first.dump()
//	second := ParsePacket("[1,[2,[3,[4,[5,6,0]]]],8,9]")
//	second.dump()
//	assert.False(t, IsCorrectOrder(first, second))
//}
//
//func TestComparePackets2(t *testing.T) {
//	// [9] vs [[8,7,6]]
//	first := List([]int{9})
//	second := &Packet{children: []*Packet{List([]int{8, 7, 6})}}
//
//	assert.False(t, IsCorrectOrder(first, second))
//}
//
//func TestComparePackets_second_is_nil(t *testing.T) {
//	// [3,4] and nil
//	threeFour := ParsePacket("[3,4]")
//	empty := ParsePacket("[]")
//	assert.False(t, IsCorrectOrder(threeFour, empty))
//}
//
//func TestIsCorrectOrder(t *testing.T) {
//	// [[1],[2,3,4]]
//	first := &Packet{children: []*Packet{List([]int{1}), List([]int{2, 3, 4})}}
//	// [[1],4]
//	second := &Packet{children: []*Packet{List([]int{1}), P(4)}}
//
//	// [1] and [1]
//	result := IsCorrectOrder(first.UnwrapForce(), second.UnwrapForce())
//	assert.True(t, result)
//
//	// [2,3,4] and 4
//	result = IsCorrectOrder(first.UnwrapForce(), second.UnwrapForce())
//	assert.True(t, result)
//}
