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
