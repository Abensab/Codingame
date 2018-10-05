package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func (p1 Point) compare(p2 Point) [2]int {
	res := [2]int{0, 0}
	if p1.x > p2.x {
		res[0] = -1
	} else if p1.x < p2.x {
		res[0] = 1
	}
	if p1.y > p2.y {
		res[1] = -1
	} else if p1.y < p2.y {
		res[1] = 1
	}
	return res
}

func (p *Point) modify(m [2]int) {
	p.x = p.x + m[0]
	p.y = p.y + m[1]
}

var directions = map[[2]int]string{
	{0, -1}:  "N",
	{1, -1}:  "NE",
	{1, 0}:   "E",
	{1, 1}:   "SE",
	{0, 1}:   "S",
	{-1, 1}:  "SW",
	{-1, 0}:  "W",
	{-1, -1}: "NW",
}

func main() {
	// LX: the X position of the light of power
	// LY: the Y position of the light of power
	// TX: Thor's starting X position
	// TY: Thor's starting Y position
	var LX, LY, TX, TY int
	fmt.Scan(&LX, &LY, &TX, &TY)
	thunder := Point{x: LX, y: LY}
	thor := Point{x: TX, y: TY}

	for {
		move := thor.compare(thunder)
		fmt.Println(directions[move])
		thor.modify(move)

	}
}
