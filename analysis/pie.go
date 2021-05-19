package pie

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func generatePieItems(data map[interface{}]nt) []opts.PieData {
	items := make([]opts.PieData, 0)
	for key, value := range data {
		items = append(items, opts.PieData{Name: key, Value: value})
	}
	return items
}

func PieShowLabel(w http.ResponseWriter, _ *http.Request, data map[interface{}]int) {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "demo",
			Subtitle: "demo",
		}),
	)

	pie.AddSeries("pie", generatePieItems(data)).SetSeriesOptions(charts.WithLabelOpts(
		opts.Label{
			Show:      true,
			Formatter: "{b}: {c}",
		}),
	)
	pie.Render(w)
}
