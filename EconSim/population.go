package sim

import "fmt"

// PopType
type PopType struct {
	needs  []Good
	amount []float64
}

var popTypes []PopType

// Pop
type Pop struct {
	Type           PopType
	size           int
	funds          float64
	fulfilledNeeds bool
}

func NewPop(Type int) {
	var p Pop
	p.Type = popTypes[0]
	p.funds = 100
	p.fulfilledNeeds = true
	p.size = 1000
	pop = p
}

func (p Pop) String() string {
	return fmt.Sprintf("%v\nFunds:%f\nFulfilled Needs:%v\nSize:%d", p.Type, p.funds, p.fulfilledNeeds, p.size)
}

var pop Pop
