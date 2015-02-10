// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package functions

import (
	"math"

	"github.com/gonum/floats"
)

// Beale implements the Beale's function.
//
// Standard starting points:
//  Easy: [1, 1]
//  Hard: [1, 4]
//
// References:
//  - Beale, E.: On an Iterative Method for Finding a Local Minimum of a
//    Function of More than One Variable. Technical Report 25, Statistical
//    Techniques Research Group, Princeton University (1958)
//  - More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained
//    optimization software. ACM Trans Math Softw 7 (1981), 17-41
type Beale struct{}

func (Beale) F(x []float64) float64 {
	if len(x) != 2 {
		panic("dimension of the problem must be 2")
	}

	f1 := 1.5 - x[0]*(1-x[1])
	f2 := 2.25 - x[0]*(1-x[1]*x[1])
	f3 := 2.625 - x[0]*(1-x[1]*x[1]*x[1])
	return f1*f1 + f2*f2 + f3*f3
}

func (Beale) Df(x, grad []float64) {
	if len(x) != 2 {
		panic("dimension of the problem must be 2")
	}
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	t1 := 1 - x[1]
	t2 := 1 - x[1]*x[1]
	t3 := 1 - x[1]*x[1]*x[1]

	f1 := 1.5 - x[0]*t1
	f2 := 2.25 - x[0]*t2
	f3 := 2.625 - x[0]*t3

	grad[0] = -2 * (f1*t1 + f2*t2 + f3*t3)
	grad[1] = 2 * x[0] * (f1 + 2*f2*x[1] + 3*f3*x[1]*x[1])
}

func (Beale) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{3, 0.5},
			F:      0,
			Global: true,
		},
	}
}

// BiggsEXP2 implements the the Biggs' EXP2 function.
//
// Standard starting point:
//  [1, 2]
//
// Reference:
//  Biggs, M.C.: Minimization algorithms making use of non-quadratic properties
//  of the objective function. IMA J Appl Math 8 (1971), 315-327; doi:10.1093/imamat/8.3.315
type BiggsEXP2 struct{}

func (BiggsEXP2) F(x []float64) (sum float64) {
	if len(x) != 2 {
		panic("dimension of the problem must be 2")
	}

	for i := 1; i <= 10; i++ {
		z := float64(i) / 10
		y := math.Exp(-z) - 5*math.Exp(-10*z)
		f := math.Exp(-x[0]*z) - 5*math.Exp(-x[1]*z) - y
		sum += f * f
	}
	return sum
}

func (BiggsEXP2) Df(x, grad []float64) {
	if len(x) != 2 {
		panic("dimension of the problem must be 2")
	}
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	for i := range grad {
		grad[i] = 0
	}
	for i := 1; i <= 10; i++ {
		z := float64(i) / 10
		y := math.Exp(-z) - 5*math.Exp(-10*z)
		f := math.Exp(-x[0]*z) - 5*math.Exp(-x[1]*z) - y

		dfdx0 := -z * math.Exp(-x[0]*z)
		dfdx1 := 5 * z * math.Exp(-x[1]*z)

		grad[0] += 2 * f * dfdx0
		grad[1] += 2 * f * dfdx1
	}
}

func (BiggsEXP2) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{1, 10},
			F:      0,
			Global: true,
		},
	}
}

// BiggsEXP3 implements the the Biggs' EXP3 function.
//
// Standard starting point:
//  [1, 2, 1]
//
// Reference:
//  Biggs, M.C.: Minimization algorithms making use of non-quadratic properties
//  of the objective function. IMA J Appl Math 8 (1971), 315-327; doi:10.1093/imamat/8.3.315
type BiggsEXP3 struct{}

func (BiggsEXP3) F(x []float64) (sum float64) {
	if len(x) != 3 {
		panic("dimension of the problem must be 3")
	}

	for i := 1; i <= 10; i++ {
		z := float64(i) / 10
		y := math.Exp(-z) - 5*math.Exp(-10*z)
		f := math.Exp(-x[0]*z) - x[2]*math.Exp(-x[1]*z) - y
		sum += f * f
	}
	return sum
}

func (BiggsEXP3) Df(x, grad []float64) {
	if len(x) != 3 {
		panic("dimension of the problem must be 3")
	}
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	for i := range grad {
		grad[i] = 0
	}
	for i := 1; i <= 10; i++ {
		z := float64(i) / 10
		y := math.Exp(-z) - 5*math.Exp(-10*z)
		f := math.Exp(-x[0]*z) - x[2]*math.Exp(-x[1]*z) - y

		dfdx0 := -z * math.Exp(-x[0]*z)
		dfdx1 := x[2] * z * math.Exp(-x[1]*z)
		dfdx2 := -math.Exp(-x[1] * z)

		grad[0] += 2 * f * dfdx0
		grad[1] += 2 * f * dfdx1
		grad[2] += 2 * f * dfdx2
	}
}

func (BiggsEXP3) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{1, 10, 5},
			F:      0,
			Global: true,
		},
	}
}

// BiggsEXP4 implements the the Biggs' EXP4 function.
//
// Standard starting point:
//  [1, 2, 1, 1]
//
// Reference:
//  Biggs, M.C.: Minimization algorithms making use of non-quadratic properties
//  of the objective function. IMA J Appl Math 8 (1971), 315-327; doi:10.1093/imamat/8.3.315
type BiggsEXP4 struct{}

func (BiggsEXP4) F(x []float64) (sum float64) {
	if len(x) != 4 {
		panic("dimension of the problem must be 4")
	}

	for i := 1; i <= 10; i++ {
		z := float64(i) / 10
		y := math.Exp(-z) - 5*math.Exp(-10*z)
		f := x[2]*math.Exp(-x[0]*z) - x[3]*math.Exp(-x[1]*z) - y
		sum += f * f
	}
	return sum
}

func (BiggsEXP4) Df(x, grad []float64) {
	if len(x) != 4 {
		panic("dimension of the problem must be 4")
	}
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	for i := range grad {
		grad[i] = 0
	}
	for i := 1; i <= 10; i++ {
		z := float64(i) / 10
		y := math.Exp(-z) - 5*math.Exp(-10*z)
		f := x[2]*math.Exp(-x[0]*z) - x[3]*math.Exp(-x[1]*z) - y

		dfdx0 := -z * x[2] * math.Exp(-x[0]*z)
		dfdx1 := z * x[3] * math.Exp(-x[1]*z)
		dfdx2 := math.Exp(-x[0] * z)
		dfdx3 := -math.Exp(-x[1] * z)

		grad[0] += 2 * f * dfdx0
		grad[1] += 2 * f * dfdx1
		grad[2] += 2 * f * dfdx2
		grad[3] += 2 * f * dfdx3
	}
}

func (BiggsEXP4) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{1, 10, 1, 5},
			F:      0,
			Global: true,
		},
	}
}

// BiggsEXP5 implements the the Biggs' EXP5 function.
//
// Standard starting point:
//  [1, 2, 1, 1, 1]
//
// Reference:
//  Biggs, M.C.: Minimization algorithms making use of non-quadratic properties
//  of the objective function. IMA J Appl Math 8 (1971), 315-327; doi:10.1093/imamat/8.3.315
type BiggsEXP5 struct{}

func (BiggsEXP5) F(x []float64) (sum float64) {
	if len(x) != 5 {
		panic("dimension of the problem must be 5")
	}

	for i := 1; i <= 11; i++ {
		z := float64(i) / 10
		y := math.Exp(-z) - 5*math.Exp(-10*z) + 3*math.Exp(-4*z)
		f := x[2]*math.Exp(-x[0]*z) - x[3]*math.Exp(-x[1]*z) + 3*math.Exp(-x[4]*z) - y
		sum += f * f
	}
	return sum
}

func (BiggsEXP5) Df(x, grad []float64) {
	if len(x) != 5 {
		panic("dimension of the problem must be 5")
	}
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	for i := range grad {
		grad[i] = 0
	}
	for i := 1; i <= 11; i++ {
		z := float64(i) / 10
		y := math.Exp(-z) - 5*math.Exp(-10*z) + 3*math.Exp(-4*z)
		f := x[2]*math.Exp(-x[0]*z) - x[3]*math.Exp(-x[1]*z) + 3*math.Exp(-x[4]*z) - y

		dfdx0 := -z * x[2] * math.Exp(-x[0]*z)
		dfdx1 := z * x[3] * math.Exp(-x[1]*z)
		dfdx2 := math.Exp(-x[0] * z)
		dfdx3 := -math.Exp(-x[1] * z)
		dfdx4 := -3 * z * math.Exp(-x[4]*z)

		grad[0] += 2 * f * dfdx0
		grad[1] += 2 * f * dfdx1
		grad[2] += 2 * f * dfdx2
		grad[3] += 2 * f * dfdx3
		grad[4] += 2 * f * dfdx4
	}
}

func (BiggsEXP5) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{1, 10, 1, 5, 4},
			F:      0,
			Global: true,
		},
	}
}

// BiggsEXP6 implements the the Biggs' EXP6 function.
//
// Standard starting point:
//  [1, 2, 1, 1, 1, 1]
//
// References:
//  - Biggs, M.C.: Minimization algorithms making use of non-quadratic
//    properties of the objective function. IMA J Appl Math 8 (1971), 315-327;
//    doi:10.1093/imamat/8.3.315
//  - More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained
//    optimization software. ACM Trans Math Softw 7 (1981), 17-41
type BiggsEXP6 struct{}

func (BiggsEXP6) F(x []float64) (sum float64) {
	if len(x) != 6 {
		panic("dimension of the problem must be 6")
	}

	for i := 1; i <= 13; i++ {
		z := float64(i) / 10
		y := math.Exp(-z) - 5*math.Exp(-10*z) + 3*math.Exp(-4*z)
		f := x[2]*math.Exp(-x[0]*z) - x[3]*math.Exp(-x[1]*z) + x[5]*math.Exp(-x[4]*z) - y
		sum += f * f
	}
	return sum
}

func (BiggsEXP6) Df(x, grad []float64) {
	if len(x) != 6 {
		panic("dimension of the problem must be 6")
	}
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	for i := range grad {
		grad[i] = 0
	}
	for i := 1; i <= 13; i++ {
		z := float64(i) / 10
		y := math.Exp(-z) - 5*math.Exp(-10*z) + 3*math.Exp(-4*z)
		f := x[2]*math.Exp(-x[0]*z) - x[3]*math.Exp(-x[1]*z) + x[5]*math.Exp(-x[4]*z) - y

		dfdx0 := -z * x[2] * math.Exp(-x[0]*z)
		dfdx1 := z * x[3] * math.Exp(-x[1]*z)
		dfdx2 := math.Exp(-x[0] * z)
		dfdx3 := -math.Exp(-x[1] * z)
		dfdx4 := -z * x[5] * math.Exp(-x[4]*z)
		dfdx5 := math.Exp(-x[4] * z)

		grad[0] += 2 * f * dfdx0
		grad[1] += 2 * f * dfdx1
		grad[2] += 2 * f * dfdx2
		grad[3] += 2 * f * dfdx3
		grad[4] += 2 * f * dfdx4
		grad[5] += 2 * f * dfdx5
	}
}

func (BiggsEXP6) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{1, 10, 1, 5, 4, 3},
			F:      0,
			Global: true,
		},
		{
			X: []float64{1.7114159947956764, 17.68319817846745, 1.1631436609697268,
				5.1865615510738605, 1.7114159947949301, 1.1631436609697998},
			F:      0.005655649925499929,
			Global: false,
		},
		{
			// X: []float64{1.22755594752403, X[1] >> 0, 0.83270306333466, X[3] << 0, X[4] = X[0], X[5] = X[2]},
			X:      []float64{1.22755594752403, 1000, 0.83270306333466, -1000, 1.22755594752403, 0.83270306333466},
			F:      0.306366772624790,
			Global: false,
		},
	}
}

// Box3D implements the Box' three-dimensional function.
//
// Standard starting point:
//  [0, 10, 20]
//
// References:
//  - Box, M.J.: A comparison of several current optimization methods, and the
//    use of transformations in constrained problems. Comput J 9 (1966), 67-77
//  - More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained
//    optimization software. ACM Trans Math Softw 7 (1981), 17-41
type Box3D struct{}

func (Box3D) F(x []float64) (sum float64) {
	if len(x) != 3 {
		panic("dimension of the problem must be 3")
	}

	for i := 1; i <= 10; i++ {
		c := -float64(i) / 10
		y := math.Exp(c) - math.Exp(10*c)
		f := math.Exp(c*x[0]) - math.Exp(c*x[1]) - x[2]*y
		sum += f * f
	}
	return sum
}

func (Box3D) Df(x, grad []float64) {
	if len(x) != 3 {
		panic("dimension of the problem must be 3")
	}
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	grad[0] = 0
	grad[1] = 0
	grad[2] = 0
	for i := 1; i <= 10; i++ {
		c := -float64(i) / 10
		y := math.Exp(c) - math.Exp(10*c)
		f := math.Exp(c*x[0]) - math.Exp(c*x[1]) - x[2]*y
		grad[0] += 2 * f * c * math.Exp(c*x[0])
		grad[1] += -2 * f * c * math.Exp(c*x[1])
		grad[2] += -2 * f * y
	}
}

func (Box3D) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{1, 10, 1},
			F:      0,
			Global: true,
		},
		{
			X:      []float64{10, 1, -1},
			F:      0,
			Global: true,
		},
		{
			// Any point at the line {a, a, 0}.
			X:      []float64{1, 1, 0},
			F:      0,
			Global: true,
		},
	}
}

// BrownBadlyScaled implements the Brown's badly scaled function.
//
// Standard starting point:
//  [1, 1]
//
// References:
//  - More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained
//    optimization software. ACM Trans Math Softw 7 (1981), 17-41
type BrownBadlyScaled struct{}

func (BrownBadlyScaled) F(x []float64) float64 {
	if len(x) != 2 {
		panic("dimension of the problem must be 2")
	}

	f1 := x[0] - 1e6
	f2 := x[1] - 2e-6
	f3 := x[0]*x[1] - 2
	return f1*f1 + f2*f2 + f3*f3
}

func (BrownBadlyScaled) Df(x, grad []float64) {
	if len(x) != 2 {
		panic("dimension of the problem must be 2")
	}
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	f1 := x[0] - 1e6
	f2 := x[1] - 2e-6
	f3 := x[0]*x[1] - 2
	grad[0] = 2*f1 + 2*f3*x[1]
	grad[1] = 2*f2 + 2*f3*x[0]
}

func (BrownBadlyScaled) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{1e6, 2e-6},
			F:      0,
			Global: true,
		},
	}
}

// BrownAndDennis implements the Brown and Dennis function.
//
// Standard starting point:
//  [25, 5, -5, -1]
//
// References:
//  - Brown, K.M., Dennis, J.E.: New computational algorithms for minimizing a
//    sum of squares of nonlinear functions. Research Report Number 71-6, Yale
//    University (1971)
//  - More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained
//    optimization software. ACM Trans Math Softw 7 (1981), 17-41
type BrownAndDennis struct{}

func (BrownAndDennis) F(x []float64) (sum float64) {
	if len(x) != 4 {
		panic("dimension of the problem must be 4")
	}

	for i := 1; i <= 20; i++ {
		c := float64(i) / 5
		f1 := x[0] + c*x[1] - math.Exp(c)
		f2 := x[2] + x[3]*math.Sin(c) - math.Cos(c)
		f := f1*f1 + f2*f2
		sum += f * f
	}
	return sum
}

func (BrownAndDennis) Df(x, grad []float64) {
	if len(x) != 4 {
		panic("dimension of the problem must be 4")
	}
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	for i := range grad {
		grad[i] = 0
	}
	for i := 1; i <= 20; i++ {
		c := float64(i) / 5
		f1 := x[0] + c*x[1] - math.Exp(c)
		f2 := x[2] + x[3]*math.Sin(c) - math.Cos(c)
		f := f1*f1 + f2*f2
		grad[0] += 4 * f * f1
		grad[1] += 4 * f * f1 * c
		grad[2] += 4 * f * f2
		grad[3] += 4 * f * f2 * math.Sin(c)
	}
}

func (BrownAndDennis) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{-11.594439904886773, 13.203630051265385, -0.40343948776868443, 0.2367787746745986},
			F:      85822.20162635627,
			Global: true,
		},
	}
}

// ExtendedPowellSingular implements the extended Powell's function.
// Its Hessian matrix is singular at the minimizer.
//
// Standard starting point:
//  [3, -1, 0, 3, 3, -1, 0, 3, ..., 3, -1, 0, 3]
//
// References:
//  - Spedicato E.: Computational experience with quasi-Newton algorithms for
//    minimization problems of moderatly large size. Towards Global
//    Optimization 2 (1978), 209-219
//  - More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained
//    optimization software. ACM Trans Math Softw 7 (1981), 17-41
type ExtendedPowellSingular struct{}

func (ExtendedPowellSingular) F(x []float64) (sum float64) {
	if len(x)%4 != 0 {
		panic("dimension of the problem must be a multiple of 4")
	}

	for i := 0; i < len(x); i += 4 {
		f1 := x[i] + 10*x[i+1]
		f2 := x[i+2] - x[i+3]
		t := x[i+1] - 2*x[i+2]
		f3 := t * t
		t = x[i] - x[i+3]
		f4 := t * t
		sum += f1*f1 + 5*f2*f2 + f3*f3 + 10*f4*f4
	}
	return sum
}

func (ExtendedPowellSingular) Df(x, grad []float64) {
	if len(x)%4 != 0 {
		panic("dimension of the problem must be a multiple of 4")
	}
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	for i := 0; i < len(x); i += 4 {
		f1 := x[i] + 10*x[i+1]
		f2 := x[i+2] - x[i+3]
		t1 := x[i+1] - 2*x[i+2]
		f3 := t1 * t1
		t2 := x[i] - x[i+3]
		f4 := t2 * t2
		grad[i] = 2*f1 + 40*f4*t2
		grad[i+1] = 20*f1 + 4*f3*t1
		grad[i+2] = 10*f2 - 8*f3*t1
		grad[i+3] = -10*f2 - 40*f4*t2
	}
}

func (ExtendedPowellSingular) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{0, 0, 0, 0},
			F:      0,
			Global: true,
		},
		{
			X:      []float64{0, 0, 0, 0, 0, 0, 0, 0},
			F:      0,
			Global: true,
		},
		{
			X:      []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			F:      0,
			Global: true,
		},
	}
}

// ExtendedRosenbrock implements the extended, multidimensional Rosenbrock
// function.
//
// Standard starting point:
//  Easy: [-1.2, 1, -1.2, 1, ...]
//  Hard: any point far from the minimum
//
// References:
//  - Rosenbrock, H.H.: An Automatic Method for Finding the Greatest or Least
//    Value of a Function. Computer J 3 (1960), 175-184
//  - http://en.wikipedia.org/wiki/Rosenbrock_function
type ExtendedRosenbrock struct{}

func (ExtendedRosenbrock) F(x []float64) (sum float64) {
	for i := 0; i < len(x)-1; i++ {
		a := 1 - x[i]
		b := x[i+1] - x[i]*x[i]
		sum += a*a + 100*b*b
	}
	return sum
}

func (ExtendedRosenbrock) Df(x, grad []float64) {
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	dim := len(x)
	for i := range grad {
		grad[i] = 0
	}
	for i := 0; i < dim-1; i++ {
		grad[i] -= 2 * (1 - x[i])
		grad[i] -= 400 * (x[i+1] - x[i]*x[i]) * x[i]
	}
	for i := 1; i < dim; i++ {
		grad[i] += 200 * (x[i] - x[i-1]*x[i-1])
	}
}

func (ExtendedRosenbrock) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{1, 1},
			F:      0,
			Global: true,
		},
		{
			X:      []float64{1, 1, 1},
			F:      0,
			Global: true,
		},
		{
			X:      []float64{1, 1, 1, 1},
			F:      0,
			Global: true,
		},
		{
			X: []float64{-0.7756592265653526, 0.6130933654850433,
				0.38206284633839305, 0.14597201855219452},
			F:      3.701428610430017,
			Global: false,
		},
		{
			X:      []float64{1, 1, 1, 1, 1},
			F:      0,
			Global: true,
		},
		{
			X: []float64{-0.9620510206947502, 0.9357393959767103,
				0.8807136041943204, 0.7778776758544063, 0.6050936785926526},
			F:      3.930839434133027,
			Global: false,
		},
		{
			X: []float64{-0.9865749795709938, 0.9833982288361819, 0.972106670053092,
				0.9474374368264362, 0.8986511848517299, 0.8075739520354182},
			F:      3.973940500930295,
			Global: false,
		},
		{
			X: []float64{-0.9917225725614055, 0.9935553935033712, 0.992173321594692,
				0.9868987626903134, 0.975164756608872, 0.9514319827049906, 0.9052228177139495},
			F:      3.9836005364248543,
			Global: false,
		},
		{
			X:      []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			F:      0,
			Global: true,
		},
	}
}

// Gaussian implements the Gaussian function.
// The function has one global minimum and a number of false local minima
// caused by the finite floating point precision.
//
// Standard starting point:
//  [0.4, 1, 0]
//
// Reference:
//  More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained optimization
//  software. ACM Trans Math Softw 7 (1981), 17-41
type Gaussian struct{}

func (Gaussian) y(i int) (yi float64) {
	switch i {
	case 1, 15:
		yi = 0.0009
	case 2, 14:
		yi = 0.0044
	case 3, 13:
		yi = 0.0175
	case 4, 12:
		yi = 0.0540
	case 5, 11:
		yi = 0.1295
	case 6, 10:
		yi = 0.2420
	case 7, 9:
		yi = 0.3521
	case 8:
		yi = 0.3989
	}
	return yi
}

func (g Gaussian) F(x []float64) (sum float64) {
	if len(x) != 3 {
		panic("dimension of the problem must be 3")
	}

	for i := 1; i <= 15; i++ {
		c := 0.5 * float64(8-i)
		b := c - x[2]
		d := b * b
		e := math.Exp(-0.5 * x[1] * d)
		f := x[0]*e - g.y(i)
		sum += f * f
	}
	return sum
}

func (g Gaussian) Df(x, grad []float64) {
	if len(x) != 3 {
		panic("dimension of the problem must be 3")
	}
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	grad[0] = 0
	grad[1] = 0
	grad[2] = 0
	for i := 1; i <= 15; i++ {
		c := 0.5 * float64(8-i)
		b := c - x[2]
		d := b * b
		e := math.Exp(-0.5 * x[1] * d)
		f := x[0]*e - g.y(i)
		grad[0] += 2 * f * e
		grad[1] -= f * e * d * x[0]
		grad[2] += 2 * f * e * x[0] * x[1] * b
	}
}

func (Gaussian) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{0.398956137837997, 1.0000190844805048, 0},
			F:      1.12793276961912e-08,
			Global: true,
		},
	}
}

// GulfResearchAndDevelopment implements the Gulf Research and Development function.
//
// Standard starting point:
//  [5, 2.5, 0.15]
//
// References:
//  - Cox, R.A.: Comparison of the performance of seven optimization algorithms
//    on twelve unconstrained minimization problems. Ref. 1335CNO4, Gulf
//    Research and Development Company, Pittsburg (1969)
//  - More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained
//    optimization software. ACM Trans Math Softw 7 (1981), 17-41
type GulfResearchAndDevelopment struct{}

func (GulfResearchAndDevelopment) F(x []float64) (sum float64) {
	if len(x) != 3 {
		panic("dimension of the problem must be 3")
	}

	for i := 1; i <= 99; i++ {
		arg := float64(i) / 100
		r := math.Pow(-50*math.Log(arg), 2.0/3.0) + 25 - x[1]
		t1 := math.Pow(math.Abs(r), x[2]) / x[0]
		t2 := math.Exp(-t1)
		t := t2 - arg
		sum += t * t
	}
	return sum
}

func (GulfResearchAndDevelopment) Df(x, grad []float64) {
	if len(x) != 3 {
		panic("dimension of the problem must be 3")
	}
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	for i := range grad {
		grad[i] = 0
	}
	for i := 1; i <= 99; i++ {
		arg := float64(i) / 100
		r := math.Pow(-50*math.Log(arg), 2.0/3.0) + 25 - x[1]
		t1 := math.Pow(math.Abs(r), x[2]) / x[0]
		t2 := math.Exp(-t1)
		t := t2 - arg
		s1 := t1 * t2 * t
		grad[0] += s1
		grad[1] += s1 / r
		grad[2] -= s1 * math.Log(math.Abs(r))
	}
	grad[0] *= 2 / x[0]
	grad[1] *= 2 * x[2]
	grad[2] *= 2
}

func (GulfResearchAndDevelopment) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{50, 25, 1.5},
			F:      0,
			Global: true,
		},
		{
			X:      []float64{99.89529935174151, 60.61453902799833, 9.161242695144592},
			F:      32.8345,
			Global: false,
		},
		{
			X:      []float64{201.662589489426, 60.61633150468155, 10.224891158488965},
			F:      32.8345,
			Global: false,
		},
	}
}

// HelicalValley implements the helical valley function of Fletcher and Powell.
// Function is not defined at x[0] = 0.
//
// Standard starting point:
//  [-1, 0, 0]
//
// References:
//  - Fletcher, R., Powell, M.J.D.: A rapidly convergent descent method for
//    minimization. Comput J 6 (1963), 163-168
//  - More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained
//    optimization software. ACM Trans Math Softw 7 (1981), 17-41
type HelicalValley struct{}

func (HelicalValley) F(x []float64) float64 {
	if len(x) != 3 {
		panic("dimension of the problem must be 3")
	}
	if x[0] == 0 {
		panic("function not defined at x[0] = 0")
	}

	theta := 0.5 * math.Atan(x[1]/x[0]) / math.Pi
	if x[0] < 0 {
		theta += 0.5
	}
	f1 := 10 * (x[2] - 10*theta)
	f2 := 10 * (math.Hypot(x[0], x[1]) - 1)
	f3 := x[2]
	return f1*f1 + f2*f2 + f3*f3
}

func (HelicalValley) Df(x, grad []float64) {
	if len(x) != 3 {
		panic("dimension of the problem must be 3")
	}
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}
	if x[0] == 0 {
		panic("function not defined at x[0] = 0")
	}

	theta := 0.5 * math.Atan(x[1]/x[0]) / math.Pi
	if x[0] < 0 {
		theta += 0.5
	}
	h := math.Hypot(x[0], x[1])
	r := 1 / h
	q := r * r / math.Pi
	s := x[2] - 10*theta
	grad[0] = 200 * (5*s*q*x[1] + (h-1)*r*x[0])
	grad[1] = 200 * (-5*s*q*x[0] + (h-1)*r*x[1])
	grad[2] = 2 * (100*s + x[2])
}

func (HelicalValley) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{1, 0, 0},
			F:      0,
			Global: true,
		},
	}
}

// Linear implements a linear function.
type Linear struct{}

func (Linear) F(x []float64) float64 {
	return floats.Sum(x)
}

func (Linear) Df(x, grad []float64) {
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	for i := range grad {
		grad[i] = 1
	}
}

// PenaltyI implements the first penalty function by Gill, Murray and Pitfield.
//
// Standard starting point:
//  [1, ..., n]
//
// References:
//  - Gill, P.E., Murray, W., Pitfield, R.A.: The implementation of two revised
//    quasi-Newton algorithms for unconstrained optimization. Report NAC 11,
//    National Phys Lab (1972), 82-83
//  - More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained
//    optimization software. ACM Trans Math Softw 7 (1981), 17-41
type PenaltyI struct{}

func (PenaltyI) F(x []float64) (sum float64) {
	for _, v := range x {
		sum += (v - 1) * (v - 1)
	}
	sum *= 1e-5

	var s float64
	for _, v := range x {
		s += v * v
	}
	sum += (s - 0.25) * (s - 0.25)
	return sum
}

func (PenaltyI) Df(x, grad []float64) {
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	s := -0.25
	for _, v := range x {
		s += v * v
	}
	for i, v := range x {
		grad[i] = 2 * (2*s*v + 1e-5*(v-1))
	}
}

func (PenaltyI) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{0.2500074995875379, 0.2500074995875379, 0.2500074995875379, 0.2500074995875379},
			F:      2.2499775008999372e-05,
			Global: true,
		},
		{
			X: []float64{0.15812230111311634, 0.15812230111311634, 0.15812230111311634,
				0.15812230111311634, 0.15812230111311634, 0.15812230111311634,
				0.15812230111311634, 0.15812230111311634, 0.15812230111311634, 0.15812230111311634},
			F:      7.087651467090369e-05,
			Global: true,
		},
	}
}

// PenaltyII implements the second penalty function by Gill, Murray and Pitfield.
//
// Standard starting point:
//  [0.5, ..., 0.5]
//
// References:
//  - Gill, P.E., Murray, W., Pitfield, R.A.: The implementation of two revised
//    quasi-Newton algorithms for unconstrained optimization. Report NAC 11,
//    National Phys Lab (1972), 82-83
//  - More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained
//    optimization software. ACM Trans Math Softw 7 (1981), 17-41
type PenaltyII struct{}

func (PenaltyII) F(x []float64) (sum float64) {
	dim := len(x)
	s := -1.0
	for i, v := range x {
		s += float64(dim-i) * v * v
	}
	for i := 1; i < dim; i++ {
		yi := math.Exp(float64(i+1)/10) + math.Exp(float64(i)/10)
		f := math.Exp(x[i]/10) + math.Exp(x[i-1]/10) - yi
		sum += f * f
	}
	for i := 1; i < dim; i++ {
		f := math.Exp(x[i]/10) - math.Exp(-1.0/10)
		sum += f * f
	}
	sum *= 1e-5
	sum += (x[0] - 0.2) * (x[0] - 0.2)
	sum += s * s
	return sum
}

func (PenaltyII) Df(x, grad []float64) {
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	dim := len(x)
	s := -1.0
	for i, v := range x {
		s += float64(dim-i) * v * v
	}
	for i, v := range x {
		grad[i] = 4 * s * float64(dim-i) * v
	}
	for i := 1; i < dim; i++ {
		yi := math.Exp(float64(i+1)/10) + math.Exp(float64(i)/10)
		f := math.Exp(x[i]/10) + math.Exp(x[i-1]/10) - yi
		grad[i] += 1e-5 * f * math.Exp(x[i]/10) / 5
		grad[i-1] += 1e-5 * f * math.Exp(x[i-1]/10) / 5
	}
	for i := 1; i < dim; i++ {
		f := math.Exp(x[i]/10) - math.Exp(-1.0/10)
		grad[i] += 1e-5 * f * math.Exp(x[i]/10) / 5
	}
	grad[0] += 2 * (x[0] - 0.2)
}

func (PenaltyII) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{0.19999933335, 0.19131670128566283, 0.4801014860897, 0.5188454026659},
			F:      9.376293007355449e-06,
			Global: true,
		},
		{
			X: []float64{0.19998360520892217, 0.010350644318663525,
				0.01960493546891094, 0.03208906550305253, 0.04993267593895693,
				0.07651399534454084, 0.11862407118600789, 0.1921448731780023,
				0.3473205862372022, 0.36916437893066273},
			F:      0.00029366053745674594,
			Global: true,
		},
	}
}

// PowellBadlyScaled implements the Powell's badly scaled function.
// The function is very flat near the minimum. A satisfactory solution is one
// that gives f(x) ≅ 1e-13.
//
// Standard starting point:
//  [0, 1]
//
// References:
//  - Powell, M.J.D.: A Hybrid Method for Nonlinear Equations. Numerical
//    Methods for Nonlinear Algebraic Equations, P. Rabinowitz (ed.), Gordon
//    and Breach (1970)
//  - More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained
//    optimization software. ACM Trans Math Softw 7 (1981), 17-41
type PowellBadlyScaled struct{}

func (PowellBadlyScaled) F(x []float64) float64 {
	if len(x) != 2 {
		panic("dimension of the problem must be 2")
	}

	f1 := 1e4*x[0]*x[1] - 1
	f2 := math.Exp(-x[0]) + math.Exp(-x[1]) - 1.0001
	return f1*f1 + f2*f2
}

func (PowellBadlyScaled) Df(x, grad []float64) {
	if len(x) != 2 {
		panic("dimension of the problem must be 2")
	}
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	f1 := 1e4*x[0]*x[1] - 1
	f2 := math.Exp(-x[0]) + math.Exp(-x[1]) - 1.0001
	grad[0] = 2 * (1e4*f1*x[1] - f2*math.Exp(-x[0]))
	grad[1] = 2 * (1e4*f1*x[0] - f2*math.Exp(-x[1]))
}

func (PowellBadlyScaled) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{1.0981593296997149e-05, 9.106146739867375},
			F:      0,
			Global: true,
		},
	}
}

// Trigonometric implements the trigonometric function.
//
// Standard starting point:
//  [1/dim, ..., 1/dim]
//
// References:
//  - Spedicato E.: Computational experience with quasi-Newton algorithms for
//    minimization problems of moderatly large size. Towards Global
//    Optimization 2 (1978), 209-219
//  - More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained
//    optimization software. ACM Trans Math Softw 7 (1981), 17-41
type Trigonometric struct{}

func (Trigonometric) F(x []float64) (sum float64) {
	var s1 float64
	for _, v := range x {
		s1 += math.Cos(v)
	}
	for i, v := range x {
		f := float64(len(x)+i+1) - float64(i+1)*math.Cos(v) - math.Sin(v) - s1
		sum += f * f
	}
	return sum
}

func (Trigonometric) Df(x, grad []float64) {
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	var s1 float64
	for _, v := range x {
		s1 += math.Cos(v)
	}

	var s2 float64
	for i, v := range x {
		f := float64(len(x)+i+1) - float64(i+1)*math.Cos(v) - math.Sin(v) - s1
		s2 += f
		grad[i] = 2 * f * (float64(i+1)*math.Sin(v) - math.Cos(v))
	}

	for i, v := range x {
		grad[i] += 2 * s2 * math.Sin(v)
	}
}

func (Trigonometric) Minima() []Minimum {
	return []Minimum{
		{
			X: []float64{0.04296456438227447, 0.043976287478192246,
				0.045093397949095684, 0.04633891624617569, 0.047744381782831,
				0.04935473251330618, 0.05123734850076505, 0.19520946391410446,
				0.1649776652761741, 0.06014857783799575},
			F:      0,
			Global: true,
		},
		{
			// TODO(vladimir-ch): If we knew the location of this minimum more
			// accurately, we could decrease defaultGradTol.
			X: []float64{0.05515090434047145, 0.05684061730812344,
				0.05876400231100774, 0.060990608903034337, 0.06362621381044778,
				0.06684318087364617, 0.2081615177172172, 0.16436309604419047,
				0.08500689695564931, 0.09143145386293675},
			F:      2.795056121876575e-05,
			Global: false,
		},
	}
}

// VariablyDimensioned implements a variably dimensioned function.
//
// Standard starting point:
//  [..., (dim-i)/dim, ...], i=1,...,dim
//
// References:
//  More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained optimization
//  software. ACM Trans Math Softw 7 (1981), 17-41
type VariablyDimensioned struct{}

func (VariablyDimensioned) F(x []float64) (sum float64) {
	for _, v := range x {
		t := v - 1
		sum += t * t
	}

	var s float64
	for i, v := range x {
		s += float64(i+1) * (v - 1)
	}
	s *= s
	sum += s
	s *= s
	sum += s
	return sum
}

func (VariablyDimensioned) Df(x, grad []float64) {
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	var s float64
	for i, v := range x {
		s += float64(i+1) * (v - 1)
	}
	for i, v := range x {
		grad[i] = 2 * (v - 1 + s*float64(i+1)*(1+2*s*s))
	}
}

func (VariablyDimensioned) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{1, 1},
			F:      0,
			Global: true,
		},
		{
			X:      []float64{1, 1, 1},
			F:      0,
			Global: true,
		},
		{
			X:      []float64{1, 1, 1, 1},
			F:      0,
			Global: true,
		},
		{
			X:      []float64{1, 1, 1, 1, 1},
			F:      0,
			Global: true,
		},
		{
			X:      []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			F:      0,
			Global: true,
		},
	}
}

// Watson implements the Watson's function.
// Dimension of the problem should be 2 <= dim <= 31. For dim == 9, the problem
// of minimizing the function is very ill conditioned.
//
// Standard starting point:
//  [0, ..., 0]
//
// References:
//  - Kowalik, J.S., Osborne, M.R.: Methods for Unconstrained Optimization
//    Problems. Elsevier North-Holland, New York, 1968
//  - More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained
//    optimization software. ACM Trans Math Softw 7 (1981), 17-41
type Watson struct{}

func (Watson) F(x []float64) (sum float64) {
	for i := 1; i <= 29; i++ {
		c := float64(i) / 29

		var s1 float64
		for j := 1; j < len(x); j++ {
			s1 += float64(j) * x[j] * math.Pow(c, float64(j-1))
		}

		var s2 float64
		for j, v := range x {
			s2 += v * math.Pow(c, float64(j))
		}
		sum += (s1 - s2*s2 - 1) * (s1 - s2*s2 - 1)
	}
	sum += x[0] * x[0]
	t := x[1] - x[0]*x[0] - 1
	sum += t * t
	return sum
}

func (Watson) Df(x, grad []float64) {
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	for i := range grad {
		grad[i] = 0
	}
	for i := 1; i <= 29; i++ {
		c := float64(i) / 29

		var s1 float64
		for j := 1; j < len(x); j++ {
			s1 += float64(j) * x[j] * math.Pow(c, float64(j-1))
		}

		var s2 float64
		for j, v := range x {
			s2 += v * math.Pow(c, float64(j))
		}

		t := s1 - s2*s2 - 1
		for j := range x {
			grad[j] += 2 * t * math.Pow(c, float64(j-1)) * (float64(j) - 2*s2*c)
		}
	}
	t := x[1] - x[0]*x[0] - 1
	grad[0] += 2*x[0] - 4*t*x[0]
	grad[1] += 2 * t
}

func (Watson) Minima() []Minimum {
	return []Minimum{
		{
			X: []float64{-0.01572508639445209, 1.0124348693647496, -0.23299162596578743,
				1.2604300879396417, -1.5137289229557926, 0.9929964325532394},
			F:      0.002287670053552403,
			Global: true,
		},
		{
			X: []float64{-1.5307036746695823e-05, 0.9997897039329611,
				0.014763963657024618, 0.146342328644901, 1.0008211015487485,
				-2.6177311373274232, 4.104403160682486, -3.1436122762314374,
				1.0526264074353293},
			F:      1.3997601380906803e-06,
			Global: true,
		},
		// TODO(vladimir-ch): More, Garbow, Hillstrom list just the value, but
		// not the location. Our minimizers find a minimum, but the value is
		// different.
		// {
		// 	// For dim == 12
		// 	F:      4.72238e-10,
		// 	Global: true,
		// },
		// TODO(vladimir-ch): netlib/uncon report a value of 2.48631d-20 for dim == 20.
	}
}

// Wood implements the Wood's function.
//
// Standard starting point:
//  [-3, -1, -3, -1]
//
// References:
//  - Colville, A.R.: A comparative study of nonlinear programming codes.
//    Report 320-2949, IBM New York Scientific Center (1968)
//  - More, J., Garbow, B.S., Hillstrom, K.E.: Testing unconstrained
//    optimization software. ACM Trans Math Softw 7 (1981), 17-41
type Wood struct{}

func (Wood) F(x []float64) (sum float64) {
	if len(x) != 4 {
		panic("dimension of the problem must be 4")
	}

	f1 := x[1] - x[0]*x[0]
	f2 := 1 - x[0]
	f3 := x[3] - x[2]*x[2]
	f4 := 1 - x[2]
	f5 := x[1] + x[3] - 2
	f6 := x[1] - x[3]
	return 100*f1*f1 + f2*f2 + 90*f3*f3 + f4*f4 + 10*f5*f5 + 0.1*f6*f6
}

func (Wood) Df(x, grad []float64) {
	if len(x) != 4 {
		panic("dimension of the problem must be 4")
	}
	if len(x) != len(grad) {
		panic("incorrect size of the gradient")
	}

	f1 := x[1] - x[0]*x[0]
	f2 := 1 - x[0]
	f3 := x[3] - x[2]*x[2]
	f4 := 1 - x[2]
	f5 := x[1] + x[3] - 2
	f6 := x[1] - x[3]
	grad[0] = -2 * (200*f1*x[0] + f2)
	grad[1] = 2 * (100*f1 + 10*f5)
	grad[2] = 2 * (-180*f3*x[2] - f4 + 0.1*f6)
	grad[3] = 2 * (90*f3 + 10*f5 - 0.1*f6)
}

func (Wood) Minima() []Minimum {
	return []Minimum{
		{
			X:      []float64{1, 1, 1, 1},
			F:      0,
			Global: true,
		},
	}
}