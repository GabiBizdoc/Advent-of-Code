package solution

// DistanceEquation calculates the distance based on the given coefficients and time.
//
// Parameters:
//
//	k: Coefficient in the equation.
//	t: Time variable in the equation.
//	d: Initial distance or constant in the equation.
//
// Returns:
//
//	Calculated distance based on the equation.
func DistanceEquation(k, t, d int) int {
	return k*(t-k) - d
}

// FindCoefficient calculates the coefficient (k) based on the given time and initial distance.
// It iterates over a range of values and counts the number of positive distances in the DistanceEquation.
//
// Parameters:
//
//	t: Time variable in the equation.
//	d: Initial distance or constant in the equation.
func FindCoefficient(t, d int) int {
	k := 0
	for i := 0; i < t; i++ {
		if DistanceEquation(i, t, d) > 0 {
			k += 1
		}
	}
	return k
}
