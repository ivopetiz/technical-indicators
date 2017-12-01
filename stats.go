
package indicators

import "math"

type mfloat []float64

// Sum returns sum of a slice.
func Sum(slice []float64) float64 {
	var sum float64 = 0.0

	for i:=0; i < len(slice); i++ {
		sum += slice[i] 
	}

	return sum
}


// Avg returns slice mean.
func Avg(slice []float64) float64 {
	var result float64

	for i:=0; i < len(slice); i++ {
		result += slice[i]
	}
	return result
}

// Std returns standard deviation of a slice.
func Std(slice []float64) float64 {
	var result []float64
	
	mean := Avg(slice)
	
	for i:=0; i < len(slice); i++ {
		result = append(result, math.Pow(slice[i]-mean,2))
	}

	return  math.Sqrt(Sum(result)/float64(len(result)))
}

// AddToAll adds a value to all slice elements.
func (slice mfloat) AddToAll(val float64) []float64 {
	var added_slice []float64

	for i:=0; i < len(slice); i++ {
		slice[i] += val
	}

	return added_slice
}
