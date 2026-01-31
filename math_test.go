package math

import (
	"math"
	"testing"

	"github.com/franela/goblin"
)

const precision = 8

var mat3x3 = [][]float64{
	{0.617988, 0.27225, 0.398392},
	{0.0552414, 0.258802, 0.800991},
	{0.60112, 0.921029, 0.413371},
}
var mat2x2 = [][]float64{
	{0.916756, 0.712766},
	{0.546127, 0.498242},
}

func TestFloat(t *testing.T) {
	g := goblin.Goblin(t)
	x := 0.1
	y := 0.2
	z := x + y
	g.Describe("FloatEqual", func() {
		g.It("consts", func() {
			g.Assert(PRECISION).Equal(12)
			g.Assert(EPSILON).Equal(1e-12)
		})
		g.It("strfloat", func() {
			g.Assert(FloatToString(2740278.8823960135)).Equal("2740278.8823960135")
		})
		g.It("float equality", func() {
			a, b := 0.1, 0.2
			g.Assert((a + b) == 0.3).IsFalse()
			g.Assert(FloatEqual(0.1+0.2, 0.3)).IsTrue()
			g.Assert(FloatEqual(math.Inf(1), math.Inf(1))).IsTrue()
			g.Assert(FloatEqual(math.Inf(-1), math.Inf(-1))).IsTrue()
			g.Assert(FloatEqual(math.MaxFloat64, math.MaxFloat64-1e100)).IsTrue()
			g.Assert(FloatEqual(0.1+0.2, 0.3)).IsTrue()
			g.Assert(FloatEqual(0.3, 0.3)).IsTrue()
			g.Assert(FloatEqual(0., -0.)).IsTrue()
			g.Assert(FloatEqual(-0., 0.)).IsTrue()
			g.Assert(FloatEqual(-0., -0.)).IsTrue()
			g.Assert(FloatEqual(0.3, 0.31)).IsFalse()
			g.Assert(FloatEqual(x+y, z)).IsTrue()
			g.Assert(FloatEqual(z-0.3, 0.0)).IsTrue()
		})
		g.It("float comparison", func() {
			g.Assert(Min(0.3, 0.3)).Equal(0.3)
			g.Assert(Min(0.2, 0.3)).Equal(0.2)
			g.Assert(Min(0.2, -0.3)).Equal(-0.3)

			g.Assert(Max(0.3, 0.3)).Equal(0.3)
			g.Assert(Max(0.2, 0.3)).Equal(0.3)
			g.Assert(Max(0.2, -0.3)).Equal(0.2)

			g.Assert(Min[int64](3, 3)).Equal(int64(3))
			g.Assert(Min[int64](2, 3)).Equal(int64(2))
			g.Assert(Min[int64](2, -3)).Equal(int64(-3))

			g.Assert(Max[int64](3, 3)).Equal(int64(3))
			g.Assert(Max[int64](2, 3)).Equal(int64(3))
			g.Assert(Max[int64](2, -3)).Equal(int64(2))

			g.Assert(Min(3, 3)).Equal(3)
			g.Assert(Min(2, 3)).Equal(2)
			g.Assert(Min(2, -3)).Equal(-3)

			g.Assert(Max(3, 3)).Equal(3)
			g.Assert(Max(2, 3)).Equal(3)
			g.Assert(Max(2, -3)).Equal(2)

		})

		g.It("float nan inf floor ceil trunc", func() {
			g.Assert(IsNaN(NaN())).IsTrue()
			g.Assert(IsNaN(1.4)).IsFalse()
			inf := Inf(1)
			g.Assert(IsInf(1.4, 1)).IsFalse()
			g.Assert(IsInf(inf, -1)).IsFalse()
			g.Assert(IsInf(inf, 1)).IsTrue()

			g.Assert(Floor(1.00011) == 1).IsTrue()
			g.Assert(Ceil(1.00011) == 2).IsTrue()
			g.Assert(Trunc(1.00011) == 1).IsTrue()
		})
		g.It("log", func() {
			g.Assert(Log2(0.5) == -1).IsTrue()
			g.Assert(Log2(1000000000) == 29.897352853986263).IsTrue()
			g.Assert(Log10(1000000000) == 9.0).IsTrue()

		})
		g.It("pow", func() {
			g.Assert(Pow(2, 3) == 8).IsTrue()
			g.Assert(Pow10(3) == 1000).IsTrue()
			g.Assert(Pow10(-2) == 0.01).IsTrue()
			g.Assert(Pow10(2) == 1024).IsFalse()
		})
		g.It("trig", func() {
			g.Assert(Cos(Pi) == -1).IsTrue()
			g.Assert(Cos(Pi/2.0) < 1.e-16).IsTrue()
			g.Assert(Cos(2*Pi) == 1).IsTrue()

			g.Assert(Sin(Pi) < 1.e-12).IsTrue()
			g.Assert(Sin(Pi/2.0) == 1).IsTrue()
			g.Assert(Sin(2*Pi) < -2.44e-16).IsTrue()

			g.Assert(Hypot(3, 4) == 5.0).IsTrue()
			g.Assert(Hypot2(3, 4) == 25.0).IsTrue()
			g.Assert(math.IsInf(Hypot2(3, math.Inf(1)), 1)).IsTrue()
			g.Assert(math.IsInf(Hypot2(3, math.Inf(-1)), 1)).IsTrue()
			g.Assert(math.IsNaN(Hypot2(3, math.NaN()))).IsTrue()
			g.Assert(Atan2(90, 15)-1.4056476493802699 < 1e-12).IsTrue()
			g.Assert(Sqrt(9) == 3).IsTrue()
			g.Assert(Sqrt(2) == Sqrt2).IsTrue()
		})

		g.It("avg", func() {
			var vals = []float64{2, 3, 1, 5.6, 2, 8, 9.8, 6.7, 9, 3, 5, 7}
			var v = Sum(vals) / float64(len(vals))
			g.Assert(Average(vals...)).Equal(v)
			g.Assert(Average[float64]()).Equal(0.0)
		})

	})

}

func TestSum(t *testing.T) {
	g := goblin.Goblin(t)
	var l1 [100]int
	var l2 [100]int64
	var l3 [100]float64

	for i := 1; i <= 100; i++ {
		idx := i - 1
		l1[idx] = i
	}

	for i := 1; i <= 100; i++ {
		idx := i - 1
		l2[idx] = int64(i)
	}

	for i := 1; i <= 100; i++ {
		idx := i - 1
		l3[idx] = float64(i)
	}

	s := 5050
	g.Describe("Sum", func() {
		g.It("sum  from int 1..100 ", func() {
			g.Assert(Sum(l1[:])).Equal(s)
		})
		g.It("sum  from int64 1..100 ", func() {
			g.Assert(Sum(l2[:])).Equal(int64(s))
		})
		g.It("sum  from float64 1..100 ", func() {
			g.Assert(Sum(l3[:])).Equal(float64(s))
		})
	})
}

func TestDet2(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Determinant 2x2 , 3x3", func() {
		g.It("3 x 3 matrix determinant", func() {
			g.Assert(Round(-0.30663810342819653, precision)).Equal(
				Round(Det3(mat3x3), precision))
		})
		g.It("2 x 2 matrix determinant ", func() {

			g.Assert(Round(0.06750558567000006, precision)).Equal(
				Round(Det2(mat2x2), precision))

			g.Assert(SignOfDet2(0.916756, 0.712766, 0.546127, 0.498242)).Equal(1)
			g.Assert(SignOfDet2(0., 0., 0., 0.)).Equal(0)
		})
	})
}

func TestValConversion(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Distance", func() {
		g.It("round float ", func() {
			g.Assert(Round(3.1415926535897, precision)).Equal(Round(Pi, precision))
		})
		g.It("rad 2 deg float ", func() {
			g.Assert(Round(Rad2deg(0.523598775598299), precision)).Equal(30.0)
		})
		g.It("rad 2 deg float ", func() {
			g.Assert(Round(Deg2rad(70.736380255432223), precision)).Equal(Round(1.2345827364, precision))
		})
		g.It("sign + ", func() {
			g.Assert(Sign(3.1415926535897)).Equal(1.0)
		})
		g.It("sign - ", func() {
			g.Assert(Sign(-3.1415926535897)).Equal(-1.0)
		})
		g.It("sign -0 ", func() {
			g.Assert(Sign(-0)).Equal(0.0)
		})
		g.It("sign +0 ", func() {
			g.Assert(Sign(+0)).Equal(0.0)
		})
	})
}

func TestMid(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Mid2D coordinates", func() {
		g.It("Mid 0 ", func() {
			g.Assert(Mid2D([]float64{0, 0}, []float64{100, 100})).Equal([]float64{50, 50})
		})
		g.It("Mid 1 ", func() {
			g.Assert(Mid2D([]float64{3, 6}, []float64{7, 9})).Equal([]float64{5.0, 7.5})
		})
		g.It("Mid 2 ", func() {
			g.Assert(Mid2D([]float64{-3, -6}, []float64{7, 9})).Equal([]float64{2.0, 1.5})
		})
	})
}
