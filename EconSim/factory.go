package sim

import (
	"fmt"
	"math"
)

// FactoryType
type FactoryType struct {
	inputGoods []Good
	amount     []float64
	outputGood Good
}

func (f FactoryType) String() string {
	var inNames []string

	for i, I := range f.inputGoods {
		if i < len(f.inputGoods)-1 {
			I.name += ","
		}
		inNames = append(inNames, I.name)
	}

	return fmt.Sprintf("%v | %v", inNames, f.outputGood.name)
}

func NewFactoryType(inGoods []int, amounts []float64, outGood int) {
	var f FactoryType
	for _, I := range inGoods {
		f.inputGoods = append(f.inputGoods, goods[I])
	}
	f.outputGood = goods[outGood]
	f.amount = amounts
	factoryTypes = append(factoryTypes, f)
}

var factoryTypes []FactoryType

// Factory
type Factory struct {
	Type           FactoryType
	producedGood   float64
	funds          float64
	fulfilledNeeds bool
	employees      int
	maxEmployees   int
}

func (f *Factory) FactoryTick(province *Province) {
	var employed = int(math.Min(math.Min(f.funds, float64(f.maxEmployees)), float64(numEmployable)))
	numEmployable -= employed
	f.employees = employed
	f.funds -= float64(f.employees) * 0.01
	province.pop.funds += float64(f.employees) * 0.01
	var idx int
	for i, v := range goods {
		if v == f.Type.outputGood {
			idx = i
			break
		}
	}
	f.producedGood = float64(f.employees) / 100
	nation.market.availableGoods[idx] += f.producedGood
	producedGoods[idx] += f.producedGood
	f.funds += nation.market.goodPrices[idx]

	if f.funds > 10 && f.employees == f.maxEmployees {
		f.maxEmployees += 100
	}
}

func (f Factory) String() string {
	return fmt.Sprintf("Input/Output:%v\nProduced:%v\nFunds:%f\nFulfilled Needs:%v\nEmployed:%d", f.Type, f.producedGood, f.funds, f.fulfilledNeeds, f.employees)
}

func NewFactory(Type int) *Factory {
	var f Factory
	f.Type = factoryTypes[Type]
	f.producedGood = 0
	f.funds = 100
	f.fulfilledNeeds = true
	f.employees = 0
	f.maxEmployees = 50
	return &f
}
