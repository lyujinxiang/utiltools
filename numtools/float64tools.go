package numtools

import "math"

func ComputeAVG(values []float64) float64 {
	var sum float64

	for _, value := range values {
		sum = sum + value
	}

	return (sum / float64(len(values)))
}

func ComputeSTDEVP(values []float64) float64 {
	var (
		AVG          float64
		vp           []float64
		stdDiv, temp float64
	)

	vp = make([]float64, len(values))

	/*Calculate mean of the data points*/
	AVG = ComputeAVG(values)
	/*Calculate standard deviation of individual data points*/
	for i, v := range values {
		temp = v - AVG
		vp[i] = (temp * temp)
	}

	/* Finally, Compute standard Deviation of data points
	 * by taking mean of individual std. Deviation.
	 */
	stdDiv = ComputeAVG(vp)
	return float64(math.Sqrt(float64(stdDiv)))
}
