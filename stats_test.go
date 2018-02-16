package indicators

import "testing"
import "reflect"

func TestAvg(t *testing.T) {

	var v float64

	v = Avg([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9})

	if v != 5 {
		t.Error("Expected 5, got ", v)
	}
}

func TestSum(t *testing.T) {

	var v float64

	v = Sum([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9})

	if v != 45 {
		t.Error("Expected 45, got ", v)
	}
}

func TestStd(t *testing.T) {

	var v float64

	v = Std([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9})

	if v > 2.738 && v < 2.739 {
		t.Error("Expected 2.738, got ", v)
	}
}

func TestAddToAll(t *testing.T) {

	var v []float64
	vTest := mfloat{1, 2, 3, 4, 5, 6, 7, 8, 9}
	vResult := []float64{3, 4, 5, 6, 7, 8, 9, 10, 11}

	v = vTest.AddToAll(2)

	if !reflect.DeepEqual(v, vResult) {
		t.Error("Expected [3,4,5,6,7,8,9,10,11], got ", v)
	}
}

func TestSubSlices(t *testing.T) {

	var testResult []float64

	slice1 := []float64{5, 4, 3, 2, 1}
	slice2 := []float64{5, 3, 4, 1, 10}
	result := []float64{0, 1, -1, 1, -9}

	testResult = SubSlices(slice1, slice2)

	if !reflect.DeepEqual(result, testResult) {
		t.Error("Expected [0, 1, -1, 1, -9], got ", testResult)
	}
}


func TestAddSlices(t *testing.T) {

	var testResult []float64

	slice1 := []float64{5, 4, 3, 2, 1}
	slice2 := []float64{5, 3, 4, 1, 10}
	result := []float64{10, 7, 7, 3, 11}

	testResult = AddSlices(slice1, slice2)

	if !reflect.DeepEqual(result, testResult) {
		t.Error("Expected [10, 7, 7, 3, 11], got ", testResult)
	}
}