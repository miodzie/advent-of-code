package main

import (
	"io"
	"strings"
	"unicode"
)

type Pair struct {
	Index         int
	First, Second *Packet
}

func ParseInput(r io.Reader) (pairs []Pair) {
	buf := new(strings.Builder)
	if _, err := io.Copy(buf, r); err != nil {
		panic(err)
	}
	index := 0
	for _, split := range strings.Split(buf.String(), "\n\n") {
		index++
		pkts := strings.Split(split, "\n")
		pairs = append(pairs,
			Pair{
				index,
				ParsePacket(pkts[0]),
				ParsePacket(pkts[1])})
	}

	return
}

func ParsePacket(line string) *Packet {
	var current = &Packet{val: -1}
	for i, c := range line {
		// open a new packet.
		if c == '[' && i != 0 {
			parent := current
			current = &Packet{parent: parent, val: -1}
		}
		if unicode.IsDigit(c) {
			current.AddN(int(c - '0'))
		}
		if c == ']' {
			// If we're inside another packet, unwrap.
			if current.parent != nil {
				current.parent.Add(current)
				current = current.parent
			}
		}
	}

	return current
}
