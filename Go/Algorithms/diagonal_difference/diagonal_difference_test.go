package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalDifferences(t *testing.T) {

	input1 := [][]int32{
		[]int32{3},
		[]int32{11, 2, 4},
		[]int32{4, 5, 6},
		[]int32{10, 8, -12},
	}

	output1 := int32(15)

	tests := []struct {
		input  [][]int32
		output int32
	}{
		{input1, output1},
	}

	for _, test := range tests {
		output := diagonalDifference(test.input)
		assert.Equal(t, test.output, output)
	}

}

func TestFindPrimaryCoords(t *testing.T) {

	input1 := [][]int32{
		[]int32{11, 2, 4},
		[]int32{4, 5, 6},
		[]int32{10, 8, -12},
	}

	output1 := []coord{
		coord{x: 0, y: 0},
		coord{x: 1, y: 1},
		coord{x: 2, y: 2},
	}

	tests := []struct {
		input  [][]int32
		output []coord
	}{
		{input1, output1},
	}

	for _, test := range tests {
		output := findPrimaryCoords(test.input)
		assert.Equal(t, test.output, output)
	}
}

func TestSecondaryCoords(t *testing.T) {
	input1 := [][]int32{
		[]int32{11, 2, 4},
		[]int32{4, 5, 6},
		[]int32{10, 8, -12},
	}

	output1 := []coord{
		coord{x: 0, y: 2},
		coord{x: 1, y: 1},
		coord{x: 2, y: 0},
	}

	tests := []struct {
		input  [][]int32
		output []coord
	}{
		{input1, output1},
	}

	for _, test := range tests {
		output := findSecondaryCoords(test.input)
		assert.Equal(t, test.output, output)
	}
}

func TestGetSumOfCoords(t *testing.T) {

	input := [][]int32{
		[]int32{11, 2, 4},
		[]int32{4, 5, 6},
		[]int32{10, 8, -12},
	}

	primary := []coord{
		coord{x: 0, y: 0},
		coord{x: 1, y: 1},
		coord{x: 2, y: 2},
	}
	secondary := []coord{
		coord{x: 0, y: 2},
		coord{x: 1, y: 1},
		coord{x: 2, y: 0},
	}

	primarysum := int32(4)
	secondarysum := int32(19)

	tests := []struct {
		input  [][]int32
		coords []coord
		output int32
	}{
		{input, primary, primarysum},
		{input, secondary, secondarysum},
	}

	for _, test := range tests {
		output := getSumOfCoords(test.input, test.coords)
		assert.Equal(t, test.output, output)
	}
}
