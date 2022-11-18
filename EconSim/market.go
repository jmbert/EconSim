package sim

import (
	"fmt"
	"math"
)

var market Market

// Market
type Market struct {
	availableGoods []float64
	goodPrices     []float64
	demand         []float64
}

func (m Market) String() string {
	var availableGoodsNames []string
	var goodPricesNames []string
	var demandNames []string
	for i, good := range goods {
		var name string
		name = fmt.Sprintf("%s | %f", good.name, m.availableGoods[i])
		if i < len(goods)-1 {
			name += ","
		}
		availableGoodsNames = append(availableGoodsNames, name)
	}
	for i, good := range goods {
		var name string
		name = fmt.Sprintf("%s | %f", good.name, m.goodPrices[i])
		if i < len(goods)-1 {
			name += ","
		}
		goodPricesNames = append(goodPricesNames, name)
	}
	for i, good := range goods {
		var name string
		name = fmt.Sprintf("%s | %f", good.name, m.demand[i])
		if i < len(goods)-1 {
			name += ","
		}
		demandNames = append(demandNames, name)
	}
	return fmt.Sprintf("Available Goods: %v \nGood Prices: %v\nDemand: %v", availableGoodsNames, goodPricesNames, demandNames)
}

func MarketTick() {
	for _, f := range factories {
		var idx int
		for i, v := range goods {
			if v == f.Type.outputGood {
				idx = i
				break
			}
		}
		market.availableGoods[idx] += f.producedGood
		f.fulfilledNeeds = true
		var didx []int
		for i, v := range goods {
			for _, I := range f.Type.inputGoods {
				if v == I {
					didx = append(didx, i)
				}
			}
		}
		for _, d := range didx {
			market.demand[d] += f.Type.amount[d]
			if market.availableGoods[d] > 0 {
				if f.funds >= f.Type.amount[d]*market.goodPrices[d] {
					market.availableGoods[d] -= f.Type.amount[d]
					if market.goodPrices[d] < 0.1 {
						f.funds -= 0.1
					} else {
						f.funds -= market.goodPrices[d]
					}
				} else {
					f.fulfilledNeeds = false
				}
			} else {
				f.fulfilledNeeds = false
			}
		}

	}

	p := pop
	t := pop.Type
	for i, n := range t.needs {
		var idx int
		for i, v := range goods {
			if v == n {
				idx = i
				break
			}
		}
		var demand = t.amount[i] * float64(pop.size) / 1000

		market.demand[idx] += math.Max(demand, 0)

		if market.availableGoods[idx] >= t.amount[i] {
			if pop.funds >= float64(demand*market.goodPrices[idx]) {
				pop.funds -= demand * market.goodPrices[idx]
				market.availableGoods[idx] -= demand
			} else {
				p.fulfilledNeeds = false
			}
		} else {
			pop.fulfilledNeeds = false
		}
	}

	for i := range goods {
		if producedGoods[i] > market.demand[i] {
			market.goodPrices[i] -= 0.1 * (producedGoods[i] - market.demand[i])
			if market.goodPrices[i] <= 0 {
				market.goodPrices[i] = 0
			}
		} else if producedGoods[i] < market.demand[i] {
			market.goodPrices[i] += 0.1 * (market.demand[i] - producedGoods[i])
		}
	}
}
