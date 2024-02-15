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
}
