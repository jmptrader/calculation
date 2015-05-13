package calc

// StdScore returns a standard score, that is, a dimensionless measure of number of standard deviations from the mean
// for a value x.
// See https://en.wikipedia.org/wiki/Standard_score.
// If the mean is sample mean and the stdDev is sample standard deviation, the result is called a t-score.
// But if you are using population mean and stdDev, the result is called a z-score.
// See http://www.statisticshowto.com/when-to-use-a-t-score-vs-z-score/.
func StdScore(x float64, mean float64, stdDev float64) float64 {
	return (x - mean) / stdDev
}
