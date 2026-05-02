package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	posX    int32
	posY    int32
	size    int32
	isAlive bool
	color   rl.Color
}

func (c *Cell) makeAlive() {
	c.color = rl.RayWhite
	c.isAlive = true
}

func (c *Cell) makeDead() {
	c.color = rl.Black
	c.isAlive = false
}

func (c *Cell) draw() {
	rl.DrawRectangle(c.posX, c.posY, c.size, c.size, c.color)
}

func seedInitialState(cells [][]*Cell) {
	// R-pentomino
	// cells[10][20].makeAlive()
	// cells[10][21].makeAlive()
	// cells[11][19].makeAlive()
	// cells[11][20].makeAlive()
	// cells[12][20].makeAlive()

	// ----
	// Pulsar (period 3)
	cells[6][10].makeAlive()
	cells[6][11].makeAlive()
	cells[6][12].makeAlive()

	cells[8][8].makeAlive()
	cells[9][8].makeAlive()
	cells[10][8].makeAlive()

	cells[8][13].makeAlive()
	cells[9][13].makeAlive()
	cells[10][13].makeAlive()

	cells[11][10].makeAlive()
	cells[11][11].makeAlive()
	cells[11][12].makeAlive()

	// mirror blocks
	cells[13][10].makeAlive()
	cells[13][11].makeAlive()
	cells[13][12].makeAlive()

	cells[15][8].makeAlive()
	cells[16][8].makeAlive()
	cells[17][8].makeAlive()

	cells[15][13].makeAlive()
	cells[16][13].makeAlive()
	cells[17][13].makeAlive()

	cells[18][10].makeAlive()
	cells[18][11].makeAlive()
	cells[18][12].makeAlive()
}

type Matrix struct {
	cells  [][]*Cell
	nbRows int
	nbCols int
}

func NewMatrix(nbRows, nbCols int) *Matrix {
	var x, y int32

	cells := make([][]*Cell, nbRows)
	for i := range cells {
		cells[i] = make([]*Cell, nbCols)
	}

	for i := 0; i < nbRows; i++ {
		for j := 0; j < nbCols; j++ {
			cells[i][j] = &Cell{
				posX:    x,
				posY:    y,
				size:    20,
				isAlive: false,
				color:   rl.Black,
			}
			x += 20
		}
		x = 0
		y += 20
	}

	seedInitialState(cells)

	return &Matrix{cells: cells, nbRows: nbRows, nbCols: nbCols}
}

func (m *Matrix) aliveNeighboursCount(i, j int) int {
	n := 0

	if i != 0 && j != 0 && i < m.nbRows-1 && j < m.nbCols-1 {
		if m.cells[i-1][j].isAlive {
			n++
		}
		if m.cells[i-1][j-1].isAlive {
			n++
		}
		if m.cells[i-1][j+1].isAlive {
			n++
		}
		if m.cells[i][j-1].isAlive {
			n++
		}
		if m.cells[i][j+1].isAlive {
			n++
		}
		if m.cells[i+1][j-1].isAlive {
			n++
		}
		if m.cells[i+1][j].isAlive {
			n++
		}
		if m.cells[i+1][j+1].isAlive {
			n++
		}
	}

	if i == 0 && j == 0 {
		if m.cells[1][0].isAlive {
			n++
		}
		if m.cells[0][1].isAlive {
			n++
		}
		if m.cells[1][1].isAlive {
			n++
		}
	}

	if i == m.nbRows-1 && j == 0 {
		if m.cells[i-1][0].isAlive {
			n++
		}
		if m.cells[i-1][1].isAlive {
			n++
		}
		if m.cells[i][1].isAlive {
			n++
		}
	}

	if i == 0 && j == m.nbCols-1 {
		if m.cells[0][j-1].isAlive {
			n++
		}
		if m.cells[1][j-1].isAlive {
			n++
		}
		if m.cells[1][j].isAlive {
			n++
		}
	}

	if i == m.nbRows-1 && j == m.nbCols-1 {
		if m.cells[i-1][j].isAlive {
			n++
		}
		if m.cells[i-1][j-1].isAlive {
			n++
		}
		if m.cells[i][j-1].isAlive {
			n++
		}
	}

	if i >= 1 && i < m.nbRows-1 && j == 0 {
		if m.cells[i-1][0].isAlive {
			n++
		}
		if m.cells[i-1][1].isAlive {
			n++
		}
		if m.cells[i][1].isAlive {
			n++
		}
		if m.cells[i+1][0].isAlive {
			n++
		}
		if m.cells[i+1][1].isAlive {
			n++
		}
	}

	if i >= 1 && i < m.nbRows-1 && j == m.nbCols-1 {
		if m.cells[i-1][j].isAlive {
			n++
		}
		if m.cells[i-1][j-1].isAlive {
			n++
		}
		if m.cells[i][j-1].isAlive {
			n++
		}
		if m.cells[i+1][j].isAlive {
			n++
		}
		if m.cells[i+1][j-1].isAlive {
			n++
		}
	}

	if i == 0 && j >= 1 && j < m.nbCols-1 {
		if m.cells[0][j-1].isAlive {
			n++
		}
		if m.cells[0][j+1].isAlive {
			n++
		}
		if m.cells[1][j-1].isAlive {
			n++
		}
		if m.cells[1][j].isAlive {
			n++
		}
		if m.cells[1][j+1].isAlive {
			n++
		}
	}

	if i == m.nbRows-1 && j >= 1 && j < m.nbCols-1 {
		if m.cells[i][j-1].isAlive {
			n++
		}
		if m.cells[i][j+1].isAlive {
			n++
		}
		if m.cells[i-1][j-1].isAlive {
			n++
		}
		if m.cells[i-1][j].isAlive {
			n++
		}
		if m.cells[i-1][j+1].isAlive {
			n++
		}
	}

	return n
}

func (m *Matrix) draw() {
	next := make([][]bool, m.nbRows)
	for i := range next {
		next[i] = make([]bool, m.nbCols)
	}

	for i := 0; i < m.nbRows; i++ {
		for j := 0; j < m.nbCols; j++ {
			n := m.aliveNeighboursCount(i, j)
			a := m.cells[i][j].isAlive
			next[i][j] = (a && (n == 2 || n == 3)) || (!a && n == 3)
		}
	}

	for i := 0; i < m.nbRows; i++ {
		for j := 0; j < m.nbCols; j++ {
			if next[i][j] {
				m.cells[i][j].makeAlive()
			} else {
				m.cells[i][j].makeDead()
			}
			m.cells[i][j].draw()
		}
	}
}

func main() {
	rl.InitWindow(800, 400, "Conway Game of Life")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	matrix := NewMatrix(20, 40)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		matrix.draw()

		rl.EndDrawing()
	}
}
