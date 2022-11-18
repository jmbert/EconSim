package sim

var numEmployable int

func Tick() {
	numEmployable = pop.size

	pop.PopTick()

	for _, f := range factories {
		f.FactoryTick()
		//fmt.Printf("%v\n\n", f)
	}

	MarketTick()
	//fmt.Printf("%v\n", market)

	ironPrice = append(ironPrice, market.goodPrices[0])
	coalPrice = append(coalPrice, market.goodPrices[1])
	steelPrice = append(steelPrice, market.goodPrices[2])
	ironSupply = append(ironSupply, float64(producedGoods[0]))
	coalSupply = append(coalSupply, float64(producedGoods[1]))
	steelSupply = append(steelSupply, float64(producedGoods[2]))
	ironDemand = append(ironDemand, float64(market.demand[0]))
	coalDemand = append(coalDemand, float64(market.demand[1]))
	steelDemand = append(steelDemand, float64(market.demand[2]))
	populationSize = append(populationSize, float64(pop.size)/1000)
	populationFunds = append(populationFunds, pop.funds/100)

	for i := range market.demand {
		market.demand[i] = 0
	}
	for i := range market.availableGoods {
		market.availableGoods[i] = 0
	}
	for i := range producedGoods {
		producedGoods[i] = 0
	}
	numEmployable = 0
}

func Run() {
	Setup()
	for i := 0; i < 1000; i++ {
		Tick()
		//fmt.Printf("\n---------------------------------------------------------\n")
	}
	Graph()
}

func Setup() {
	goods = []Good{{"Iron", 5}, {"Coal", 10}, {"Steel", 20}}
	popTypes = []PopType{{[]Good{goods[2]}, []float64{1}}}
	NewFactoryType([]int{0, 1}, []float64{2, 1}, 2)
	NewFactoryType([]int{}, []float64{}, 0)
	NewFactoryType([]int{}, []float64{}, 1)
	market.availableGoods, market.goodPrices, market.demand, producedGoods = make([]float64, len(goods)), make([]float64, len(goods)), make([]float64, len(goods)), make([]float64, len(goods))
	for i := range market.goodPrices {
		market.goodPrices[i] = 1
	}
	NewFactory(0)
	NewFactory(1)
	NewFactory(2)
	NewPop(0)
}
