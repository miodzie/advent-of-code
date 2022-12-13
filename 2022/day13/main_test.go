package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePacket(t *testing.T) {
	packet := parsePacket("[1,[2,[3,[4,[5,6,7]]]],8,9]")
	fiveSixSeven := List([]int{5, 6, 7})
	four := Packet{items: []Packet{P(4), fiveSixSeven}}
	three := Packet{items: []Packet{P(3), four}}
	two := Packet{items: []Packet{P(2), three}}
	one := Packet{items: []Packet{P(1), two, P(8), P(9)}}
	//packet.dump()
	//one.dump()
	assert.Equal(t, one, packet)
}

func TestComparePackets3(t *testing.T) {
	// [1,[2,[3,[4,[5,6,7]]]],8,9] vs [1,[2,[3,[4,[5,6,0]]]],8,9]
	//first := Packet{
	//	items: []Packet{
	//		{val: 1, items: {}},
	//	}}
	//assert.Equal(t, 1, first.UnwrapForce())
	//first.UnwrapForce()
	//first.UnwrapForce()
	//first.UnwrapForce()
	//five67 := List([]int{5, 6, 7})
	//assert.Equal(t, five67, first.UnwrapForce())
}

func TestComparePackets2(t *testing.T) {
	// [9] vs [[8,7,6]]
	first := List([]int{9})
	second := Packet{items: []Packet{List([]int{8, 7, 6})}}

	assert.False(t, ComparePackets(first, second))
}

func TestComparePackets(t *testing.T) {
	// [[1],[2,3,4]]
	first := Packet{items: []Packet{List([]int{1}), List([]int{2, 3, 4})}}
	// [[1],4]
	second := Packet{items: []Packet{List([]int{1}), P(4)}}

	assert.True(t, ComparePackets(first, second))
}

func TestComparePackets_second_is_nil(t *testing.T) {
	// [3,4] and nil
	threeFour := Packet{items: []Packet{List([]int{3, 4})}}
	empty := Packet{}
	assert.True(t, ComparePackets(threeFour, empty))
}

func TestIsCorrectOrder(t *testing.T) {
	// [[1],[2,3,4]]
	first := Packet{items: []Packet{List([]int{1}), List([]int{2, 3, 4})}}
	// [[1],4]
	second := Packet{items: []Packet{List([]int{1}), P(4)}}

	// [1] and [1]
	result := IsCorrectOrder(first.UnwrapForce(), second.UnwrapForce())
	assert.True(t, result)

	// [2,3,4] and 4
	result = IsCorrectOrder(first.UnwrapForce(), second.UnwrapForce())
	assert.True(t, result)
}

func TestPacket_Unwrap(t *testing.T) {
	// [[4,4],4,4]
	twoFours := []Packet{P(4), P(4)}
	packet := Packet{items: []Packet{
		{items: twoFours},
		{val: 4},
		{val: 4},
	}}

	// [4,4]
	expected := Packet{items: twoFours}
	next, err := packet.Unwrap()
	assert.Nil(t, err)
	assert.Equal(t, expected, next)

	// 4,
	next, err = packet.Unwrap()
	assert.Nil(t, err)
	assert.Equal(t, P(4), next)
	// 4,
	next, err = packet.Unwrap()
	assert.Nil(t, err)
	assert.Equal(t, P(4), next)

	next, err = packet.Unwrap()
	assert.ErrorIs(t, err, ErrPacketEmpty)
}
