
// Package indicators provides functions that can be used as 
// indicators for finance products.
package indicators


// Sma calculates simple moving average of a slice for a certain
// number of time periods.
func (slice mfloat) Sma(period int) []float64 {
	
	var sma_slice []float64
	
	for i:=period; i < len(slice); i++ {
		sma_slice = append(sma_slice, Sum(slice[i-period:i])/float64(period))
	}
	
	return sma_slice
}


// Ema calculates exponential moving average of a slice for a certain
// number of time periods.
func (slice mfloat) Ema(period int) []float64 {
	
	var ema_slice []float64

	k := 2 / (period+1)

	ema_slice = append(ema_slice, Sum(slice[0:period])/float64(period))

	for i:=period; i < len(slice); i++ {
		ema_slice = append(ema_slice, slice[i] * float64(k) + ema_slice[len(ema_slice)-1] * float64(1-k))
	}
	
	return ema_slice
}


// BollingerBands returns upper band, lower band and simple moving
// average of a slice.
func BollingerBands(slice mfloat, period int, n_std float64) ([]float64, []float64, []float64) {

	var upper_band, lower_band, middle_band mfloat

	middle_band = slice.Sma(period)
    std := Std(middle_band)
    upper_band = middle_band.AddToAll(std * n_std)
    lower_band = middle_band.AddToAll(-1.0 * std * n_std)

	return middle_band, upper_band, lower_band
}


// MACD returns a slice with z EMA of x and y EMAs corresponding
// to moving average convergence divergence of init prices slice.
func MACD(slice mfloat, x, y, z int) []float64 {
	
	var ret_slice, x_ema mfloat

	ret_slice = slice.Ema(y)
	x_ema = slice.Ema(x)
	
	aux := len(x_ema)-len(ret_slice)

	for i:=0; i< len(ret_slice); i++ {
		ret_slice[i] = ret_slice[i] - x_ema[i+aux]
	} 

	return ret_slice.Ema(z)
}