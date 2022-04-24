package neuralNetwork

import (
	"math/rand"
	"nn/math"
)

// Public atributes
type NeuralNetwork struct{}

// Private atributes
var (
	bias_ih    math.Matrix
	bias_ho    math.Matrix
	weigths_ih math.Matrix
	weigths_ho math.Matrix
)

// Setup of the neural network
func (nn *NeuralNetwork) Init(Inodes int, Hnodes int, Onodes int) {
	bias_ih = make(math.Matrix)
	for i := 0; i < Hnodes; i++ {
		bias_ih.Add(i, 0, (rand.Float64()*2)-1)
	}

	bias_ho = make(math.Matrix)
	for i := 0; i < Onodes; i++ {
		for j := 0; j < 1; j++ {
			bias_ho.Add(i, j, (rand.Float64()*2)-1)
		}
	}

	weigths_ih = make(math.Matrix)
	for i := 0; i < Hnodes; i++ {
		for j := 0; j < Inodes; j++ {
			weigths_ih.Add(i, j, (rand.Float64()*2)-1)
		}
	}

	weigths_ho = make(math.Matrix)
	for i := 0; i < Onodes; i++ {
		for j := 0; j < Hnodes; j++ {
			weigths_ho.Add(i, j, (rand.Float64()*2)-1)
		}
	}
}

// Training the network
func (nn NeuralNetwork) Train(input, expected math.Matrix, learningRate float64) {
	// ### FEEDFORWARD ###
	hidden := math.Multiplay(weigths_ih, input)
	hidden = hidden.Sum(bias_ih)
	hidden.Sigmoid()

	// HIDDEN -> OUTPUT
	output := math.Multiplay(weigths_ho, hidden)
	output = output.Sum(bias_ho)
	output.Sigmoid()

	// ### BACKPROPAGATION ###

	// OUTPUT <- HIDDEN
	output_error := expected.Sub(output)
	d_output := math.Dsigmoid(output)
	// update gradient_OUTPUT
	gradient_OUTPUT := math.Hadamard(d_output, output_error)
	gradient_OUTPUT = math.ScalarMultiply(gradient_OUTPUT, learningRate)
	// Adjust Bias between OUTPUT <- HIDDEN
	bias_ho = math.Sum(bias_ho, gradient_OUTPUT)

	// Adjust Weigths between OUTPUT <- HIDDEN
	hidden_T := math.Transpose(hidden)
	weigths_ho_deltas := math.Multiplay(gradient_OUTPUT, hidden_T)
	weigths_ho = weigths_ho.Sum(weigths_ho_deltas)

	// INPUT <- HIDDEN
	weigths_ho_T := math.Transpose(weigths_ho)
	hidden_error := math.Multiplay(weigths_ho_T, output_error)
	d_hidden := math.Dsigmoid(hidden)
	input_T := math.Transpose(input)

	// update gradient_HIDDEN
	gradient_H := math.Hadamard(d_hidden, hidden_error)
	gradient_H = math.ScalarMultiply(gradient_H, learningRate)
	// Adjust Bias between OUTPUT <- HIDDEN
	bias_ih = bias_ih.Sum(gradient_H)

	// Adjust Weigths between INPUT <- HIDDEN
	weigths_ih_deltas := math.Multiplay(gradient_H, input_T)
	weigths_ih = weigths_ih.Sum(weigths_ih_deltas)
}

// Predicts the best decision given an entry
func (NeuralNetwork) Predict(input math.Matrix) math.Matrix {
	// INPUT -> HIDDEN
	hidden := math.Multiplay(weigths_ih, input)
	hidden = hidden.Sum(bias_ih)
	hidden.Sigmoid()

	// HIDDEN -> OUTPUT
	output := math.Multiplay(weigths_ho, hidden)
	output = output.Sum(bias_ho)
	output.Sigmoid()

	return output
}
