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
		{weight: 13, id: 1, amount: 0},
		{weight: 7, id: 2, amount: 0},
		{weight: 8.6, id: 3, amount: 0},
		{weight: 12, id: 4, amount: 0},
		{weight: 14, id: 5, amount: 0},
		{weight: 15, id: 6, amount: 0},
		{weight: 1.1, id: 7, amount: 0},
		{weight: 1.99, id: 8, amount: 0},
		{weight: 1.17, id: 9, amount: 0},
		{weight: 0.9, id: 10, amount: 0},
	}
}
