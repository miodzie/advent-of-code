package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCoord_TaxicabDistance(t *testing.T) {
	s := Coord{2, 18, Sensor}
	b := Coord{-2, 15, Beacon}
	assert.Equal(t, Coord{X: 4, Y: 3}, s.TaxicabDistance(b))
}

func TestSolution1(t *testing.T) {
	f, _ := os.Open("input")
	s, b := ParseInput(f)
	s = append(s, b...)
	assert.Equal(t, 26, NonBeaconsAtY(s, 2_000_000))
	// 4893541 is too low.
	// 4893537 is too low
	// 4893541
}

func TestNumNonBeaconsAtY_Example1(t *testing.T) {
	example, _ := os.Open("example")
	s, b := ParseInput(example)
	s = append(s, b...)
	assert.Equal(t, 26, NonBeaconsAtY(s, 10))
}

func TestParseInput(t *testing.T) {
	f, _ := os.Open("example")
	sensors, beacons := ParseInput(f)
	assert.Equal(t, Coord{2, 18, Sensor}, sensors[0])
	assert.Equal(t, Coord{-2, 15, Beacon}, beacons[0])
	assert.Len(t, sensors, 14)
	assert.Len(t, beacons, 14)
}

func TestParseLine(t *testing.T) {
	l := "Sensor at x=2, y=18: closest beacon is at x=-2, y=15"
	sensor, beacon := parseLine(l)
	assert.Equal(t, Coord{2, 18, Sensor}, sensor)
	assert.Equal(t, Coord{-2, 15, Beacon}, beacon)
}
