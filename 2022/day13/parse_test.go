package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseInput(t *testing.T) {
	f, err := os.Open("example")
	if err != nil {
		panic(err)
	}
	expected := Pair{1,
		ParsePacket("[1,1,3,1,1]"),
		ParsePacket("[1,1,5,1,1]")}
	pairs := ParseInput(f)

	for _, p := range pairs {
		p.First.dump()
		p.Second.dump()
		fmt.Println()
	}

	assert.Len(t, pairs, 8)
	assert.Equal(t, expected, pairs[0])
}

func TestParsePacket(t *testing.T) {
	packet := ParsePacket("[1,[2,[3,[4,[5,6,7]]]],8,9]")
	fiveSixSeven := List([]int{5, 6, 7})
	four := &Packet{children: []*Packet{P(4), fiveSixSeven}}
	three := &Packet{children: []*Packet{P(3), four}}
	two := &Packet{children: []*Packet{P(2), three}}
	one := &Packet{children: []*Packet{P(1), two, P(8), P(9)}}
	assert.Equal(t, one.String(), packet.String())
}

func TestParsePacket2(t *testing.T) {
	pkt := ParsePacket("[[1],[2,3,4]]")
	expected := Packet{children: []*Packet{List([]int{1}), List([]int{2, 3, 4})}}

	assert.Equal(t, expected.String(), pkt.String())
}

func TestParsePacketOne(t *testing.T) {
	three := ParsePacket("[3]")
	assert.Equal(t, "[3]", three.String())
}
