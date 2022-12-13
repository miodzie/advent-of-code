package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPacket_UnwrapDoubleEmpty(t *testing.T) {
	pkt := ParsePacket("[[[]]]")
	expected := ParsePacket("[[]]")
	inside := pkt.UnwrapForce()
	assert.Equal(t, expected.String(), inside.String())
}

func TestPacket_Unwrap(t *testing.T) {
	// [[4,4],4,4]
	twoFours := []*Packet{P(4), P(4)}
	packet := Packet{children: []*Packet{
		{children: twoFours},
		{val: 4},
		{val: 4},
	}}

	// [4,4]
	expected := &Packet{children: twoFours}
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
