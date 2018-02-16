
// Package indicators provides functions that can be used as
// indicators for finance products.
package indicators

// Sma calculates simple moving average of a slice for a certain
// number of time periods.
func (slice mfloat) SMA(period int) []float64 {

	var smaSlice []float64

	for i := period; i < len(slice); i++ {
		smaSlice = append(smaSlice, Sum(slice[i-period:i])/float64(period))
	}

	return smaSlice
}

// Ema calculates exponential moving average of a slice for a certain
// number of tiSmame periods.
func (slice mfloat) EMA(period int) []float64 {

	var emaSlice []float64

	k := 2 / (period + 1)

	emaSlice = append(emaSlice, Sum(slice[0:period])/float64(period))

	for i := period; i < len(slice); i++ {
		emaSlice = append(emaSlice, slice[i]*float64(k)+emaSlice[len(emaSlice)-1]*float64(1-k))
	}

	return emaSlice
}

// BollingerBands returns upper band, lower band and simple moving
// average of a slice.
func BollingerBands(slice mfloat, period int, nStd float64) ([]float64, []float64, []float64) {

	var upperBand, lowerBand, middleBand mfloat

	middleBand = slice.SMA(period)
	std := Std(middleBand)
	upperBand = middleBand.AddToAll(std * nStd)
	lowerBand = middleBand.AddToAll(-1.0 * std * nStd)

	return middleBand, upperBand, lowerBand
}

// MACD stands for moving average convergence divergence.
func MACD(data mfloat, ema ...int) ([]float64, []float64) {

	var macd, ema1, ema2, ema3 mfloat

	if len(ema) < 3 {
		ema = []int{12, 26, 9}
	}

	ema1 = data.EMA(ema[0])
	ema2 = data.EMA(ema[1])
	macd = SubArrays(ema1, ema2)
	ema3 = macd.EMA(ema[2])

	return macd, ema3
}

// OBV means On Balance Volume.
func OBV(priceData, volumeData mfloat) []float64 {

	obv := []float64{volumeData[0]}

	for i, vol := range volumeData[1:] {
		if priceData[i] > priceData[i-1] {
			obv = append(obv, obv[i-1]+vol)
		} else if priceData[i] < priceData[i-1] {
			obv = append(obv, obv[i-1]-vol)
		} else {
			obv = append(obv, obv[i-1])
		}
	}

	return obv
}


// Ichimoku Cloud.
func IchimokuCloud(priceData, lowData, highData mfloat, configs []int) ([]float64, []float64, []float64,[]float64, []float64) {

	conversionLine, baseLine, leadSpanA, leadSpanB, lagSpan []float64

	conversionLine = (highData.SMA(9) - lowData.SMA(9))/2
	baseLine	   = (highData.SMA(26) - lowData.SMA(26))/2
	leadSpanA	   = (conversionLine + baseLine)/2
	leadSpanB	   = (highData.SMA(52) - lowData.SMA(52))/2
	lagSpan		   = priceData[0:len(priceData)-26]
	
	return conversionLine, baseLine, leadSpanA, leadSpanB, lagSpan
}