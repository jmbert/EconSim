package sim

import (
	"net/http"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

var ironPrice []float64
var coalPrice []float64
var steelPrice []float64
var ironSupply []float64
var coalSupply []float64
var steelSupply []float64
var ironDemand []float64
var coalDemand []float64
var steelDemand []float64
var producedGoods []float64
var populationSize []float64
var populationFunds []float64

func xLabel(len int) []string {
	var s []string
	for i := 0; i < len; i++ {
		s = append(s, strconv.FormatInt(int64(i+1), 10))
	}
	return s
}

func generateItems(list []float64) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(list); i++ {
		items = append(items, opts.LineData{Value: list[i]})
	}
	return items
}

func marketGraph(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeInfographic}),
		charts.WithTitleOpts(opts.Title{
			Title: "Market",
		}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
	)

	// Put data into instance
	line.SetXAxis(xLabel(len(ironPrice))).
		AddSeries("Iron Prices", generateItems(ironPrice)).
		AddSeries("Coal Prices", generateItems(coalPrice)).
		AddSeries("Steel Prices", generateItems(steelPrice)).
		AddSeries("Iron Supply", generateItems(ironSupply)).
		AddSeries("Coal Supply", generateItems(coalSupply)).
		AddSeries("Steel Supply", generateItems(steelSupply)).
		AddSeries("Iron Demand", generateItems(ironDemand)).
		AddSeries("Coal Demand", generateItems(coalDemand)).
		AddSeries("Steel Demand", generateItems(steelDemand))
	line.Render(w)
}

func pops(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeInfographic}),
		charts.WithTitleOpts(opts.Title{
			Title: "Pops",
		}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
	)

	// Put data into instance
	line.SetXAxis(xLabel(len(ironPrice))).
		AddSeries("Population Size", generateItems(populationSize)).
		AddSeries("Population Funds", generateItems(populationFunds))
	line.Render(w)
}

func Graph() {
	http.HandleFunc("/market", marketGraph)
	http.HandleFunc("/pops", pops)
	http.ListenAndServe(":8080", nil)
}
