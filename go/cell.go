package main

// Cell type definition
type Cell struct {
    Alive bool
}

// Cell constructor
func NewCell() *Cell {
    return &Cell{Alive: false}
}

// Cell stringer
func (c *Cell) String() string {
    if c.Alive == true {
        return "◉"
    }
    return " "
}
