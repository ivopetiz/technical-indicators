package indicators

import "math"

type mfloat []float64

// Avg returns 'data' average.
func Avg(data []float64) float64 {

	return Sum(data) / float64(len(data))
}

// Sum returns the sum of all elements of 'data'.
func Sum(data []float64) float64 {

	var sum float64

	for _, x := range data {
		sum += x
	}

	return sum
}

// Std returns standard deviation of a slice.
func Std(slice []float64) float64 {

	var result []float64

	mean := Avg(slice)

	for i := 0; i < len(slice); i++ {
		result = append(result, math.Pow(slice[i]-mean, 2))
	}

	return math.Sqrt(Sum(result) / float64(len(result)))
}

// AddToAll adds a value to all slice elements.
func (slice mfloat) AddToAll(val float64) []float64 {

	var addedSlice []float64

	for i := 0; i < len(slice); i++ {
		slice[i] += val
	}

	return addedSlice
}

// SubArrays subtracts two arrays.
func SubArrays(array1, array2 []float64) []float64 {

	var result []float64

	for i := 0; i < len(array1); i++ {
		result = append(result, array1[i]-array2[i])
	}

	return result
}
