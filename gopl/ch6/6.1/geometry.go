package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

type colorpoint struct {
	*point
	color int
}

func distance(p, q point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p point) distance(q point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p *point) update() {
	p.x = 0
	p.y = 0
}

func main() {
	p := point{1, 2}
	q := point{2, 3}

	p.update()
	fmt.Printf("x=%f, y=%f\n", p.x, p.y)

	r := colorpoint{&point{3, 3}, 1}
	r.update()
	fmt.Println(r.distance(q))
}
