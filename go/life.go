package main

import(
    "math/rand"
)

// Life type definition
type Life struct {
    width, height, limit, steps int
    current, next *Grid
}

// Life constructor
func NewLife(width, height, limit int) *Life {
    current := NewGrid(width, height)
    next := NewGrid(width, height)
    return &Life{width: width, height: height, limit: limit, current: current, next: next, steps: 0}
}

// Function that generates a random initial life
func (life *Life) Generate() {
    for i:= 0; i < life.width * life.height / 2 + 1; i++ {
        life.current.GetCell(rand.Intn(life.width), rand.Intn(life.height)).Alive = true
    }
}

// Function that calculates thge next generation and swaps current and next buffers
func (life *Life) NextGeneration() bool {
    if life.limit != 0 && life.steps >= life.limit {
        return false
    }
    for x := 0; x < life.width; x++ {
        for y := 0; y < life.height; y++ {
            count := life.current.CountAliveNeighbors(x, y)
            if life.current.GetCell(x, y).Alive && count >= 2 && count <= 3 {
                life.next.GetCell(x, y).Alive = true
            } else if !life.current.GetCell(x, y).Alive && count == 3 {
                life.next.GetCell(x, y).Alive = true
            } else {
                life.next.GetCell(x, y).Alive = false
            }
        }
    }
    life.current, life.next = life.next, life.current
    life.steps++
    return true
}

// Life stringer
func (life *Life) String() string {
    return life.current.String()
}
