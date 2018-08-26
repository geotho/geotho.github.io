package smplr

import (
	"sort"
	"testing"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func TestWeightedSampler(t *testing.T) {
	type testcase struct {
		name    string
		weights map[string]int
	}

	const (
		quantile = 0.99
		N        = 1000000
	)

	testcases := []testcase{
		testcase{
			name:    "fourValsEqualWeights",
			weights: map[string]int{"a": 1, "b": 1, "c": 1, "d": 1},
		},
		testcase{
			name:    "twoValEqualWeights",
			weights: map[string]int{"a": 1, "b": 1},
		},
		testcase{
			name:    "twoValDifferentWeights",
			weights: map[string]int{"a": 7, "b": 1},
		},
		testcase{
			name:    "twoValVeryDifferentWeights",
			weights: map[string]int{"a": 99, "b": 1},
		},
		testcase{
			name:    "fourValDifferentWeights",
			weights: map[string]int{"a": 1, "b": 1, "c": 2, "d": 4},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			ws := NewWeightedSampler()
			total := 0
			for k, v := range tt.weights {
				total += v
				ws.Set(k, v)
			}

			expected := make(map[string]int, len(tt.weights))
			counts := make(map[string]int, len(tt.weights))

			for k, v := range tt.weights {
				expected[k] = (N * v) / total
				counts[k] = 0
			}

			for i := 0; i < N; i++ {
				k := ws.Pick()
				counts[k]++
			}

			obs := toSortedArray(counts)
			exp := toSortedArray(expected)

			t.Logf("\nobs: %v\nexp: %v\n", obs, exp)

			distance := stat.ChiSquare(obs, exp)
			distribution := distuv.ChiSquared{
				K: float64(len(tt.weights) - 1),
			}

			criticalValue := distribution.Quantile(quantile)

			if criticalValue < distance {
				t.Errorf("critical value (%.1f percentile) = %f < chi-squared distance = %f", quantile*100, criticalValue, distance)
			}
		})
	}
}

func toSortedArray(counts map[string]int) []float64 {
	type pair struct {
		key string
		v   float64
	}

	pairs := make([]pair, 0, len(counts))

	for k, v := range counts {
		pairs = append(pairs, pair{k, float64(v)})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].key < pairs[j].key
	})

	vs := make([]float64, 0, len(pairs))
	for _, p := range pairs {
		vs = append(vs, p.v)
	}

	return vs
}
