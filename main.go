package main

import (
	"fmt"
	"math/rand"
	"nn/XOR"
	"nn/neuralNetwork"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

const (
	I_NODES      = 2
	H_NODES      = 2
	O_NODES      = 1
	learningRate = 0.1

	MIN = 0
	MAX = 3

	TRAINING_COUNT = 10000
)

func main() {
	var nn = neuralNetwork.NeuralNetwork{}

	// Table XOR
	table := XOR.Table{}
	table.SetXOR()
	table.Show()

	nn.Init(I_NODES, H_NODES, O_NODES)
	stop := true

	fmt.Printf("Learning to perform the XOR operation\n")
	for stop {
		for i := 0; i < TRAINING_COUNT; i++ {
			index := rand.Intn(MAX-MIN+1) + MIN
			nn.Train(table.Input[index], table.Output[index], learningRate)
		}

		nnResponse1 := nn.Predict(table.Input[3])[0][0]
		nnResponse2 := nn.Predict(table.Input[1])[0][0]

		if nnResponse1 <= 0.009 && nnResponse2 >= 0.98 {
			stop = false
			fmt.Printf("\n\nConvergence criterion reached!!!\n")
			fmt.Printf("%f < 0.009\n", nnResponse1)
			fmt.Printf("%f > 0.98\n", nnResponse2)
			fmt.Printf("[%1.f][%1.f] = %f (right answer %1.f)\n", table.Input[3][0][0], table.Input[3][1][0], nnResponse1, table.Output[3][0][0])
			fmt.Printf("[%1.f][%1.f] = %f (right answer %1.f)\n", table.Input[1][0][0], table.Input[1][1][0], nnResponse2, table.Output[1][0][0])
		} else {
			fmt.Printf("[%1.f][%1.f] = %f (right answer %1.f)\n", table.Input[3][0][0], table.Input[3][1][0], nnResponse1, table.Output[3][0][0])
			fmt.Printf("[%1.f][%1.f] = %f (right answer %1.f)\n\n", table.Input[1][0][0], table.Input[1][1][0], nnResponse2, table.Output[1][0][0])
		}
	}
}
