package main

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDiagonalDifferences(t *testing.T){

	input1 := [][]int32{
		[]int32{3},
		[]int32{11, 2, 4},
		[]int32{4, 5, 6},
		[]int32{10, 8, -12},
	}

	output1 := int32(15)

	tests := []struct {
		input   [][]int32
		output int32
	}{
		{input1, output1},
	}

	for _, test := range tests{
		output := diagonalDifference(test.input)
		assert.Equal(t, test.output, output)
	}

}