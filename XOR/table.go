package XOR

import (
	"fmt"
	"nn/math"
)

/*
	XOR function should return a structure like this
	XOR_TABLE = {
		input:
			[
				[0, 0], -> input[0] ->  [0, 0]
				[0, 1], -> input[1] ->  [0, 1]
				[1, 0], -> input[2] ->  [1, 0]
				[0, 0], -> input[3] ->  [0, 0]
			],
		output:
			[
				[0], -> output[0] -> [0]
				[1], -> output[1] -> [1]
				[1], -> output[2] -> [1]
				[0], -> output[3] -> [0]
			]
	}
*/

type Table struct {
	Input  []math.Matrix
	Output []math.Matrix
}

// math of the XOR table
func (t *Table) SetXOR() {
	input := make([]math.Matrix, 4)
	output := make([]math.Matrix, 4)

	/*
		 [0] = { - row 1
			[0, 0] = 0, - col 0
			[1, 0] = 0 - col 1
		}
	*/
	input[0] = make(math.Matrix)
	input[0].Add(0, 0, 0)
	input[0].Add(1, 0, 0)

	/*
		 [1] = { - row 2
			[0, 0] = 0, - col 0
			[1, 0] = 1 - col 1
		}
	*/
	input[1] = make(math.Matrix)
	input[1].Add(0, 0, 0)
	input[1].Add(1, 0, 1)

	/*
		 [2] = { - row 3
			[0, 0] = 1, - col 0
			[1, 0] = 0 - col 1
		}
	*/
	input[2] = make(math.Matrix)
	input[2].Add(0, 0, 1)
	input[2].Add(1, 0, 0)

	/*
		 [3] = { - row 4
			[0, 0] = 0, - col 0
			[1, 0] = 0 - col 1
		}
	*/
	input[3] = make(math.Matrix)
	input[3].Add(0, 0, 1)
	input[3].Add(1, 0, 1)

	// [0][0] = 0
	output[0] = make(math.Matrix)
	output[0].Add(0, 0, 0)

	// [1][0] = 1
	output[1] = make(math.Matrix)
	output[1].Add(0, 0, 1)

	// [2][0] = 1
	output[2] = make(math.Matrix)
	output[2].Add(0, 0, 1)

	// [3][0] = 0
	output[3] = make(math.Matrix)
	output[3].Add(0, 0, 0)

	t.Input = input
	t.Output = output
}

func (t *Table) Show() {
	input := t.Input
	output := t.Output
	i, j, k := 0, 0, 0

	fmt.Println("-- XOR --")
	for i = range input {
		for j = range input[i] {
			for k = range input[i][j] {
				if j == 0 {
					fmt.Printf("%1.f ", input[i][j][k])
				} else {
					fmt.Printf("%1.f ", input[i][j][k])
				}
			}

		}
		fmt.Printf(" = %1.f\n", output[i][k][k])
	}
	fmt.Println("-- XOR --\n")
}
