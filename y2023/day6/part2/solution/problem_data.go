package solution

import "math"

type ProblemData struct {
	Timing   int
	Distance int
}

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

// FindCoefficientUsingRoots is a better version of FindCoefficient
func FindCoefficientUsingRoots(t, d int) int {
	x1, x2 := FindRoots(float64(t), float64(d))
	a, b := int(x1), int(x2)

	// if the rounded-down value of x2 (Math.Floor(x2)) is a solution to the DistanceEquation,
	// it means our distance is equal to their distance. Therefore, we ignore it.
	if DistanceEquation(b, t, d) == 0 {
		b -= 1
	}

	return b - a
}

// FindRoots calculates the roots of a quadratic equation and returns the results.
//
// Parameters:
//
//	t:     Time variable in the quadratic equation.
//	d: Initial distance or constant in the quadratic equation.
//
// Returns:
//
//	The roots of the quadratic equation
func FindRoots(t, d float64) (float64, float64) {
	// -k^2 + t*k - d = 0

	// -b +- sqrt(b^2 - 4*a*c)
	// -----------------------
	//			2*a

	a := -1.0
	b := t
	c := -d

	discriminant := b*b - 4*a*c

	x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	x2 := (-b - math.Sqrt(discriminant)) / (2 * a)

	return x1, x2
}
