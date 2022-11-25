package sim

type Nation struct {
	provinces []*Province
	market    *Market
}

var nation Nation
