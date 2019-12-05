package intro

import "testing"

type testpair struct {
	values  []float64
	average float64
}

func Average(data []float64) float64 {
	var sum float64
	for _, v := range data {
		sum = sum + v
	}
	return sum / float64(len(data))
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

// http://www.golang-book.com/books/intro/12
func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
