package main

import (
	"bufio"
	"fmt"
	"github.com/miodzie/advent-of-code/util"
	"io"
)

func makeTunnel(r io.Reader) Tunnels {
	s, b := ParseInput(r)
	tunnels := make(Tunnels)
	tunnels.AddCoords(s)
	tunnels.AddCoords(b)
	return tunnels
}

func NonBeaconsAtY(tunnels []Coord, y int) (sum int) {
	var xMin, xMax int
	atY := make(map[Coord]bool)
	for _, t := range tunnels {
		if t.X <= xMin {
			xMin = t.X
		}
		if t.X >= xMax {
			xMax = t.X
		}
		if t.Y == y && t.Type == Beacon {
			atY[t] = true
		}
	}
	if xMin < 0 {
		xMin *= -1
	}
	distance := xMax + xMin
	fmt.Println(xMax, xMin, atY, len(atY))

	return distance - len(atY)
}

type Tunnels map[Coord]Type

func (t Tunnels) AddCoords(coords []Coord) {
	for _, c := range coords {
		t[c] = c.Type
	}
}

type Type int

const (
	Air Type = iota
	Sensor
	Beacon
)

type Coord struct {
	X, Y int
	Type Type
}

func (c Coord) TaxicabDistance(b Coord) Coord {
	return Coord{X: util.Abs(c.X - b.X), Y: util.Abs(c.Y - b.Y)}
}

func ParseInput(reader io.Reader) (sensors []Coord, beacons []Coord) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		s, b := parseLine(scanner.Text())
		sensors = append(sensors, s)
		beacons = append(beacons, b)
	}
	return
}

func parseLine(line string) (sensor Coord, beacon Coord) {
	sensor.Type = Sensor
	beacon.Type = Beacon
	fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
		&sensor.X, &sensor.Y, &beacon.X, &beacon.Y)
	return sensor, beacon
}
