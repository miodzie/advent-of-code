package main

import (
	"sort"
	"strconv"
)

type PacketSlice []*Packet

func (x PacketSlice) Len() int { return len(x) }
func (x PacketSlice) Less(i, j int) bool {
	return IsCorrectOrder(x[i].children, x[j].children) == -1
}
func (x PacketSlice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func decoderKey(pkts []*Packet) int {
	// Append divider packets
	pkts = append(pkts, ParsePacket("[[2]]"))
	pkts = append(pkts, ParsePacket("[[6]]"))

	sort.Sort(PacketSlice(pkts))

	var indices []int
	for i, pkt := range pkts {
		if pkt.String() == "[[2]]" || pkt.String() == "[[6]]" {
			indices = append(indices, i+1)
		}
	}

	return indices[0] * indices[1]
}

func CountOrderedPackets(pairs []Pair) (sum int) {
	var correct []string
	for _, pair := range pairs {
		if IsCorrectOrder(pair.First.children, pair.Second.children) == -1 {
			correct = append(correct, strconv.Itoa(pair.Index))
			sum += pair.Index
		}
	}
	//fmt.Printf("Correct Indices: %s", strings.Join(correct, ","))
	return
}

func IsCorrectOrder(left []*Packet, right []*Packet) int {
	for i := 0; i < len(left) && i < len(right); i++ {
		l, r := left[i], right[i]

		if !l.IsList() && !r.IsList() {
			if l.val < r.val {
				return -1
			}
			if l.val > r.val {
				return 1
			}
		} else {
			lChildren := l.children
			rChildren := r.children
			if !l.IsList() {
				lChildren = []*Packet{l}
			}
			if !r.IsList() {
				rChildren = []*Packet{r}
			}
			if res := IsCorrectOrder(lChildren, rChildren); res != 0 {
				return res
			}
		}
	}

	if len(left) < len(right) {
		return -1
	}
	if len(left) > len(right) {
		return 1
	}

	return 0
}
