package main

import (
	"errors"
	"fmt"
	"unicode"
)

func ComparePackets(first Packet, second Packet) bool {
	for ok := true; ok; {
		one, err := first.Unwrap()
		two, err2 := second.Unwrap()
		fmt.Println(one, two)
		// Both packets are empty.
		if err != nil && err2 != nil {
			return true
		}
		// Second packet is empty, they're in the right order.
		if err2 != nil && err == nil {
			return true
		}
		if !IsCorrectOrder(one, two) {
			return false
		}
	}

	return true
}

func IsCorrectOrder(first Packet, second Packet) bool {
	// If they're both not lists, just compare vals.
	if !first.IsList() && !second.IsList() {
		return first.val <= second.val
	}
	// If one of them isn't a list, and the other is,
	// convert the non list to a list.
	if !first.IsList() {
		first = List([]int{first.val})
	}
	if !second.IsList() {
		second = List([]int{second.val})
	}

	for !first.Empty() && !second.Empty() {
		if !IsCorrectOrder(first.UnwrapForce(), second.UnwrapForce()) {
			return false
		}
	}

	return true
}

var ErrPacketEmpty = errors.New("packet is empty")

type Packet struct {
	val   int
	items []Packet
}

func (p *Packet) UnwrapForce() Packet {
	pkt, err := p.Unwrap()
	if err != nil {
		panic(err) // you dingus
	}
	return pkt
}

func (p *Packet) Unwrap() (Packet, error) {
	if len(p.items) == 0 {
		return Packet{}, ErrPacketEmpty
	}
	var pkt Packet
	pkt, p.items = p.items[0], p.items[1:]
	return pkt, nil
}

func (p *Packet) IsList() bool {
	return len(p.items) != 0
}

func (p *Packet) Empty() bool {
	return len(p.items) == 0
}

func (p *Packet) AddN(n int) {
	p.items = append(p.items, P(n))
}

func (p *Packet) Add(pkt Packet) {
	p.items = append(p.items, pkt)
}
func (p *Packet) dump() {
	fmt.Println(p)
}

func (p *Packet) String() string {
	var s string
	if !p.IsList() {
		return fmt.Sprint(p.val)
	}
	s += "["
	for i, pkt := range p.items {
		s += pkt.String()
		if i != len(p.items)-1 {
			s += ","
		}
	}
	s += "]"
	return s
}

func P(n int) Packet {
	return Packet{val: n}
}

func List(n []int) Packet {
	var items []Packet
	for _, i := range n {
		items = append(items, P(i))
	}
	return Packet{items: items}
}

func parsePacket(line string) Packet {
	var packets []*Packet
	cur := -1
	for _, c := range line {
		if c == '[' {
			cur++
			packets = append(packets, &Packet{})
		}
		packet := packets[cur]
		if unicode.IsDigit(c) {
			packet.AddN(int(c - '0'))
		}
		// go to previous parent
		if c == ']' {
			prev := packets[cur]
			cur--
			if cur >= 0 {
				packets[cur].Add(*prev)
			}
		}
	}

	return *packets[0]
}
