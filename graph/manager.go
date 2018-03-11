package graph

const COUNT = 5

type Manager interface {
	// choose graph for record processing
	// clustering graphs if we can't create new graph for new TID
	ManageRecord(record Record)
}

type manager struct {
	graphArray 		[COUNT]graph
	currentElement 	int

	// that is simple hack:
	// we map COUNT of graphs to their TIDs:
	// for example, ``graphArray[0]`` manages next set of TIDs: 10, 11, 95.
	// and ``graphArray[1]`` manages TIDs: 3, 15, 200
	// so, this field will contain:
	//	0	-> 10	true
	//		-> 11	true
	//		-> 95	true
	//	1	-> 3	true
	//		-> 15	true
	//		-> 200	true
	//	2	-> empty
	//	3	-> empty
	// 	..
	graphToTheirTIDs map[int]map[uint64]bool
}

func CreateManager() Manager {
	return &manager{
		// when new tid will registered, increment this value:
		// that's mean that last used graph of array: graphArray[currentElement]
		currentElement: -1,
	}
}

func (m *manager) ManageRecord(record Record) {

}
