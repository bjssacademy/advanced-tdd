package main

import "testing"

func TestPlacesShip(t *testing.T) {
	// Arrange
	grid := NewGrid()

	// Act
	row := 2
	column := 3

	grid.PlaceShip(row, column)

	// Assert
	got := grid.isShipPresent(row, column)
	want := true

	if got != want {
		t.Error("Ship was not placed")
	}
}
