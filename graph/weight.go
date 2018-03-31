package graph

type Weight interface {
	GetMin()		int
	GetMax()		int
	GetDelta()		int
	GetDispersion()	int

	Recompute(ts uint64)
}

type weight struct {
	min			int
	max			int
	delta		int
	dispersion 	int
}

func (w *weight) GetMin() int {
	return w.min
}
func (w *weight) GetMax() int {
	return w.max
}
func (w *weight) GetDelta() int {
	return w.delta
}
func (w *weight) GetDispersion() int {
	return w.dispersion
}
func (w *weight) Recompute(ts uint64) {
	panic("TODO: should be implement!!")
}

// this factory will be used in tests,
// for access to fields and assertion
func createWeight(ts uint64) *weight {
	// TODO: timestamp param should be used
	return &weight {
		min: 0,
		max: 0,
		delta: 0,
		dispersion: 0,
	}
}

func CreateWeight(ts uint64) Weight {
	return createWeight(ts)
}
