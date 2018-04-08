package graph

type Weight interface {
	GetMin()		uint64
	GetMax()		uint64
	GetDelta()		uint64
	GetDispersion()	uint64

	Recompute(previous uint64, ts uint64)
}

type weight struct {
	min			uint64
	max			uint64
	delta		uint64
	dispersion 	uint64
}

func (w *weight) GetMin() uint64 {
	return w.min
}
func (w *weight) GetMax() uint64 {
	return w.max
}
func (w *weight) GetDelta() uint64 {
	return w.delta
}
func (w *weight) GetDispersion() uint64 {
	return w.dispersion
}
func (w *weight) Recompute(previous uint64, ts uint64) {
	diff := ts - previous
	if diff < w.min {
		w.min = diff
	} else if diff > w.max {
		w.max = diff
	}
	w.delta = w.max - w.min
}

// this factory will be used in tests,
// for access to fields and assertion
func createWeight(previous uint64, ts uint64) Weight {
	diff := ts - previous
	return &weight {
		min: diff,
		max: diff,
		delta: 0,
		dispersion: 0,
	}
}
