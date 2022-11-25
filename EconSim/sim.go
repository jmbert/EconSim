package sim

var numEmployable int

func (province *Province) ProvinceTick(nation *Nation) {
	numEmployable = province.pop.size

	province.pop.PopTick(province)

	for _, f := range province.factories {
		f.FactoryTick(province)
		//fmt.Printf("%v\n\n", f)
	}

	//fmt.Printf("%v\n", market)

	numEmployable = 0
}

func Run() {
	Setup()
	for i := 0; i < 1000; i++ {

		for _, p := range nation.provinces {
			p.ProvinceTick(&nation)
		}

		nation.market.MarketTick(&nation)
		ironPrice = append(ironPrice, nation.market.goodPrices[0])
		coalPrice = append(coalPrice, nation.market.goodPrices[1])
		steelPrice = append(steelPrice, nation.market.goodPrices[2])
		ironSupply = append(ironSupply, float64(producedGoods[0]))
		coalSupply = append(coalSupply, float64(producedGoods[1]))
		steelSupply = append(steelSupply, float64(producedGoods[2]))
		ironDemand = append(ironDemand, float64(nation.market.demand[0]))
		coalDemand = append(coalDemand, float64(nation.market.demand[1]))
		steelDemand = append(steelDemand, float64(nation.market.demand[2]))
		for i := range nation.market.demand {
			nation.market.demand[i] = 0
		}
		for i := range nation.market.availableGoods {
			nation.market.availableGoods[i] = 0
		}
		for i := range producedGoods {
			producedGoods[i] = 0
		}

		//fmt.Printf("\n---------------------------------------------------------\n")
	}
	Graph()
}

func Setup() {
	goods = []Good{{"Iron", 5}, {"Coal", 10}, {"Steel", 20}}
	popTypes = []PopType{{[]Good{goods[2]}, []float64{1}}}
	NewFactoryType([]int{0, 1}, []float64{1, 0.5}, 2)
	NewFactoryType([]int{}, []float64{}, 0)
	NewFactoryType([]int{}, []float64{}, 1)
	var m Market
	nation.market = &m
	nation.market.availableGoods, nation.market.goodPrices, nation.market.demand, producedGoods = make([]float64, len(goods)), make([]float64, len(goods)), make([]float64, len(goods)), make([]float64, len(goods))
	for i := range nation.market.goodPrices {
		nation.market.goodPrices[i] = 1
	}

	nation.provinces = append(nation.provinces, &Province{})
	nation.provinces = append(nation.provinces, &Province{})
	nation.provinces = append(nation.provinces, &Province{})

	for _, p := range nation.provinces {
		p.ProvinceSetup()
	}
}
