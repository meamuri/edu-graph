package graph

const COUNT = 5

type Manager interface {

}

type manager struct {
	graphArray [COUNT]graph
	currentElement int
}

func CreateManager() Manager {
	return &manager{
		currentElement: -1,
	}
}
