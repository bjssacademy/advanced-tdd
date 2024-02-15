package main

const (
	ROWS = 7
	COLUMNS = 7

	SHIP = "SHIP"
)

type Grid struct {
	locations [ROWS][COLUMNS]string
}

func NewGrid() *Grid {
	return &Grid{}
}

func (g *Grid) PlaceShip(row int, col int) {
	g.locations[row][col] = SHIP
}

func (g *Grid) isShipPresent(row int, col int) bool {
	return g.locations[row][col] == SHIP
}