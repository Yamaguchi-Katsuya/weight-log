package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

const THEME = types.ThemeChalk

type WeightLogs struct {
	WeightLogs []WeightLog `json:"weight_logs"`
}

type WeightLog struct {
	Date   string  `json:"date"`
	Weight float64 `json:"weight"`
}

func generateLineItems(weightLogs WeightLogs) []opts.LineData {
	items := make([]opts.LineData, 0)
	for _, weightLog := range weightLogs.WeightLogs {
		items = append(items, opts.LineData{Value: weightLog.Weight})
	}
	return items
}

func handler(w http.ResponseWriter, _ *http.Request) {
	page := components.NewPage()
	jsonFile, err := os.Open("weight-log.json")
	if err != nil {
		return
	}
	defer jsonFile.Close()
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var weightLogs WeightLogs
	err = json.Unmarshal(jsonData, &weightLogs)
	if err != nil {
		panic(err)
	}
	page.AddCharts(
		lineExampleSmooth(weightLogs),
	)
	page.Render(w)
}

func lineExampleSmooth(weightLogs WeightLogs) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme:  THEME,
			Width:  "90%",
			Height: "800px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Weight Log",
			Subtitle: "Weight Log",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Type: "category",
			AxisLabel: &opts.AxisLabel{
				Formatter: "{value}",
				Interval:  "0",
				Rotate:    45,
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Type:        "value",
			Min:         60,
			Max:         70,
			SplitNumber: 20,
			AxisLabel: &opts.AxisLabel{
				Formatter: "{value} kg",
			},
			SplitLine: &opts.SplitLine{
				Show: opts.Bool(true),
				LineStyle: &opts.LineStyle{
					Color: "#e6e6e6",
					Type:  "dashed",
				},
			},
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      opts.Bool(true),
			Trigger:   "axis",
			TriggerOn: "mousemove|click",
			Formatter: "{b}: {c} kg",
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:       "slider",
			Start:      0,
			End:        100,
			XAxisIndex: []int{0},
		}),
		charts.WithGridOpts(opts.Grid{
			Show:   opts.Bool(true),
			Left:   "10%",
			Right:  "10%",
			Bottom: "15%",
			Top:    "10%",
		}),
	)

	var xAxis []string
	for _, weightLog := range weightLogs.WeightLogs {
		xAxis = append(xAxis, weightLog.Date)
	}
	line.SetXAxis(xAxis).
		AddSeries("Weight", generateLineItems(weightLogs)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(true)}))
	return line
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8082", nil)
}
