package main

import (
    "bytes"
)

// Grid type definition
type Grid struct {
    width, height int
    cells [][]*Cell
}

// Grid constructor
func NewGrid(width, height int) *Grid {
    cells := make([][]*Cell, height)
    for i := range cells {
        cells[i] = make([]*Cell, width)
        for j := range cells[i] {
            cells[i][j] = NewCell()
        }
    }
    return &Grid{width: width, height:height, cells: cells}
}

// Function to get cells, just for code beauty
// Obs: Maybe I should check is (x,y) is Inboud here and return an error if not
func (g *Grid) GetCell(x, y int) *Cell {
    return g.cells[y][x]
}

// Function that checks is 'x' and 'y' combination is 
func (g *Grid) IsInbound(x, y int) bool {
    if x < 0 || x >= g.width {
        return false
    }
    if y < 0 || y >= g.height {
        return false
    }
    return true
}

// Function that counts the number of alive cells neighbors of (x,y)
func (g *Grid) CountAliveNeighbors(x, y int) int {
    count := 0
    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            // ignore (x,y) in this count, checks if is Inbound and if is alive
            if !(i == 0 && j == 0) && g.IsInbound(x + i, y + j) && g.GetCell(x + i, y + j).Alive {
                count++;
            }
        }
    }
    return count
}

// Grid stringerclea
func (g *Grid) String() string {
    var buffer bytes.Buffer
    // top border
    buffer.WriteString("┏")
    for i := 0; i < g.width; i++ {
        buffer.WriteString("━")
    }
    buffer.WriteString("┓\n")
    for y := range g.cells {
        // left border
        buffer.WriteString("┃")
        for x := range g.cells[y] {
            buffer.WriteString(g.GetCell(x, y).String())
        }
        // right border
        buffer.WriteString("┃\n")
    }
    // bottom border
    buffer.WriteString("┗")
    for i := 0; i < g.width; i++ {
        buffer.WriteString("━")
    }
    buffer.WriteString("┛\n")
    return buffer.String()
}
