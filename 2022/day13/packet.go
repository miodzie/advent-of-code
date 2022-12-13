package main

import (
	"errors"
	"fmt"
)

var ErrPacketEmpty = errors.New("packet is empty")

type Packet struct {
	val         int
	children    []*Packet
	parent      *Packet
	unwrapIndex int
}

func P(n int) *Packet {
	return &Packet{val: n}
}

func List(n []int) *Packet {
	var items []*Packet
	for _, i := range n {
		items = append(items, P(i))
	}
	return &Packet{children: items}
}

func (p *Packet) UnwrapForce() *Packet {
	pkt, err := p.Unwrap()
	if err != nil {
		panic(err) // you dingus
	}
	return pkt
}

func (p *Packet) Unwrap() (*Packet, error) {
	if p.unwrapIndex == len(p.children) {
		return &Packet{}, ErrPacketEmpty
	}
	if len(p.children) == 0 {
		return &Packet{}, ErrPacketEmpty
	}
	var pkt *Packet
	pkt = p.children[p.unwrapIndex]
	p.unwrapIndex++
	return pkt, nil
}

func (p *Packet) IsList() bool {
	return len(p.children) != 0 || p.val == -1
}

func (p *Packet) Empty() bool {
	return len(p.children) == 0
}

func (p *Packet) AddN(n int) {
	p.children = append(p.children, P(n))
}

func (p *Packet) Add(pkt *Packet) {
	p.children = append(p.children, pkt)
}
func (p *Packet) dump() {
	fmt.Println(p)
}

func (p *Packet) String() string {
	var s string
	if !p.IsList() {
		if p.val == -1 {
			return ""
		}
		return fmt.Sprint(p.val)
	}
	s += "["
	for i, pkt := range p.children {
		s += pkt.String()
		if i != len(p.children)-1 {
			s += ","
		}
	}
	s += "]"
	return s
}

func (p *Packet) lessThanEqual(pkt *Packet) bool {
	fmt.Printf("Comparing %d and %d", p.val, pkt.val)
	return p.val <= pkt.val
}
