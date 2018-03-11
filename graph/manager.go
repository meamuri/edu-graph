package graph

const COUNT = 5

type Manager interface {
	// choose graph for record processing
	// clustering graphs if we can't create new graph for new TID
	ManageRecord(record Record)
}

type manager struct {
	graphArray 		[COUNT]Graph
	currentElement 	int

	// that is simple hack:
	// tid -> graphArray[index]
	// we get index of graph which contain info about TID
	graphToTheirTIDs map[int]int
}

func CreateManager() Manager {
	return &manager{
		// when new tid will registered, increment this value:
		// that's mean that last used graph of array: graphArray[currentElement]
		currentElement: -1,
	}
}

func (m *manager) clusteringGraphs() {
	// some steps
	m.currentElement = 0		// we have only first graph, may be
}

func (m *manager) ManageRecord(record Record) {
	if arrIndex, ok := m.graphToTheirTIDs[record.Tid]; ok {
		m.graphArray[arrIndex].RegisterRecord(record)
		return
	}
	// if record.TID is unknown tid for manager, register new graph
	// or clustering if array is full

	m.currentElement += 1
	if m.currentElement == COUNT {
		// clustering
	} else {
		m.graphArray[m.currentElement] = NewGraph(record)
		m.graphToTheirTIDs[record.Tid] = m.currentElement
	}

	// m.graphArray[m.currentElement].RegisterRecord(record)

}
