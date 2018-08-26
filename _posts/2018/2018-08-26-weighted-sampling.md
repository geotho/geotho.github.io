---
layout: "post"
title: "Random weighted sampling"
date: "2018-08-26 14:10"
---

Suppose you have some mapping from keys to weights and you want to pick a key
at random from the map proportional to its weight. How do you do it?

Build a map that stores (key, weight) pairs.
Remember the total weight of the map.
When picking, pick a random number `r`, then iterate through the map.
If the value `v` for key `k` is greater than `r`, return `k`.
Otherwise, subtract `v` from `r` and continue.

```go
package smplr

import (
	"math/rand"
	"time"
)

type WeightedSampler struct {
	total   int
	weights map[string]int
	src     *rand.Rand
}

func NewWeightedSampler() *WeightedSampler {
	return &WeightedSampler{
		weights: make(map[string]int),
		src:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (w *WeightedSampler) Set(key string, weight int) {
	if weight < 0 {
		panic("weight cannot be less than 0")
	}

	w.total -= w.weights[key]

	if weight == 0 {
		delete(w.weights, key)
		return
	}

	w.total += weight
	w.weights[key] = weight
}

func (w *WeightedSampler) Pick() string {
	if w.total == 0 {
		return ""
	}

	r := w.src.Intn(w.total)

	for k, v := range w.weights {
		if r < v {
			return k
		}
		r -= v
	}

	return ""
}
```

Once you've done all that hard work, how can you test it?

When unit testing deterministic behaviour, we specify expected outputs for given inputs,
and assert they are equal.

When unit testing non-deterministic behaviour, like our weighted sampler here,
it's unwise to do this for two reasons:

1. Even if it is possible to make the behaviour deterministic, e.g. by seeding a random generator
or mocking its behaviour, doing so ties the test to the implementation rather than the interface.
2. Deterministic tests would still require robust statistical analysis that proper non-deterministic tests require anyway, so you're actually creating more work for yourself.

So let's assume we want to test our weighted random sampler with a statistical test.

One method might be to just try it out and eyeball the results. Here's an example unit test:

```go
package smplr

import (
	"testing"
)

func TestWeightedSampler(t *testing.T) {
	ws := NewWeightedSampler()
	ws.Set("a", 1)
	ws.Set("b", 1)
	ws.Set("c", 1)
	ws.Set("d", 1)

	N := 10000
	counts := make(map[string]int, 4)
	for i := 0; i < N; i++ {
		k := ws.Pick()
		counts[k]++
	}

	t.Logf("counts: %v\n", counts)
}
```

I ran this, and got a log line that looked like:

```
counts: map[b:2556 a:2460 c:2513 d:2471]
```

which looks plausible: the expected value for each of these is 2500. But statistically how confident should I be?

Well one could incorporate a chi-squared test here to find out.

We can use chi-squared to test whether the difference between our observed frequencies and our expected frequencies is significant. All that's left to do now is to pick an appropriate
quantile and number of trials. Picking a higher quantile is essentially saying "this test should
fail by random chance 1 in every X times". For a unit test in your CI pipeline, you want X to be quite high: maybe 1000 or 10000. Or maybe a million or more if you're Google.

Another thing to bear in mind: ensure you're writing tests for your algorithm, rather
than for the source of randomness itself. It's easy to go off the deep end and test too much.
Ensure you assume your source of randomness is infact random and that you aren't bothering to test that part.

Here's a Go table-driven test which tests my random sampler:

```go
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
```

Great! I'm missing a unit test for the single-key case, and also for the zero-key case (which should panic). But other than that I'd be happy to make a PR for this code.

I'll add a quick note on time and space complexity. The space complexity here is O(n) in the number of keys added (though some say that's O(1) because we aren't storing anything additional to the input.)

The time complexity for insertion is O(1), and for picking it is O(n), scaling linearly with number of keys. There are ways you can use more memory and increase insertion time in order to make picking faster (O(logN) or even O(1)), but of course before doing any performance enhancements it's wise to benchmark first.




