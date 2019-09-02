package main

// LinearRegression :
func LinearRegression(XTrain, YTrain Matrix, learningRate float64, epoch int64) (Matrix, float64) {
	XTrain = XTrain.Transpose(true)
	YTrain = YTrain.Transpose(true)

	var nX = XTrain.Row()
	var trainingExamples = XTrain.Col()
	var parameters = Random(1, nX, 0.001)
	var bias float64 // bias term (y = wx + b)
	var cost float64 // variable to store cost

	for j := int64(0); j < epoch; j++ {

		var prediction = Dot(parameters, XTrain)
		// PrintMatrix(&parameters)
		prediction.Add(bias, true)
		var loss = SubElementwise(prediction, YTrain)
		m := make(Matrix, 1, 1)
		for i := 0; i < nX; i++ {
			m[0] = XTrain[i]
			loss = MulElementwise(loss, m)
			cost = Sum(loss, 0)[0][0] / float64(trainingExamples) // (1/m)
			parameters[0][i] = parameters[0][i] - learningRate*cost
		}
		bias = bias - learningRate*cost
	}
	return parameters, bias
}

// Predict :
func Predict(parameters Matrix, bias float64, XTest Matrix) Matrix {
	XTest = XTest.Transpose(true)
	var prediction = Dot(parameters, XTest)
	prediction.Add(bias, true)
	return prediction
}
