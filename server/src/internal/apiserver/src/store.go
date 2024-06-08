package apiserver

type good struct {
	weight float64
	id     int
	amount int
}

type order struct {
	item  []good
	ready bool
	id    int
}

func StartStore() []good {
	return []good{
		{weight: 1, id: 1, amount: 0},
		{weight: 1.87, id: 2, amount: 0},
		{weight: 1.7, id: 3, amount: 0},
		{weight: 1.2, id: 4, amount: 0},
		{weight: 1.5, id: 5, amount: 0},
		{weight: 1.9, id: 6, amount: 0},
	}
}
