package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Repository struct {
	Lang    string          `json:"lang"`
	Result  [][]interface{} `json:"result"`
	Time    int64           `json:"time"`
	TypeMap [][]string      `json:"typeMap"`
	UID     string          `json:"uid"`
}

var data map[string]int

func main() {
	filePath := "./gacha-list-146544092.json"
	fObj, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fObj.Close()

	fReader := make([]byte, 1024)
	fContext := []byte{}

	for {
		f, err := fObj.Read(fReader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fContext = append(fContext, fReader[:f]...)
	}

	var repo Repository

	err = json.Unmarshal(fContext, &repo)
	if err != nil {
		fmt.Println(err)
		return
	}

	data = make(map[string]int)

	// repo.Result[0] 角色祈愿
	// repo.Result[1] 武器祈愿
	// repo.Result[2] 常驻祈愿
	// repo.Result[3] 新手祈愿
	// fmt.Printf("%T\n", repo.Result[0][1])
	var title string
	for _, value := range repo.Result[0][1].([]interface{}) {
		// value.([]interface{})[0] = "2020-11-07 10:20:02",
		// value.([]interface{})[1] = "黎明神剑",
		// value.([]interface{})[2] = "武器",
		// value.([]interface{})[3] =  3
		// fmt.Println(strconv.FormatFloat(value.([]interface{})[3].(float64), 'f', 0, 64))
		// fmt.Println(value.([]interface{})[3].(float64))
		switch strconv.FormatFloat(value.([]interface{})[3].(float64), 'f', 0, 64) {
		case "3":
			title = "三星"
		case "4":
			title = "四星"
		case "5":
			title = "五星"
		}
		data[title]++
	}

	fmt.Println(data)

	http.HandleFunc("/", PieShowLabel)
	http.ListenAndServe(":8090", nil)
}

func generatePieItems() []opts.PieData {
	items := make([]opts.PieData, 0)
	for key, value := range data {
		items = append(items, opts.PieData{Name: key, Value: value})
		// items = append(items, opts.PieData{Name: key, Value: value})
	}
	return items
}

func PieShowLabel(w http.ResponseWriter, _ *http.Request) {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "原神",
			Subtitle: "角色祈愿",
		}),
	)

	pie.AddSeries("pie", generatePieItems()).SetSeriesOptions(charts.WithLabelOpts(
		opts.Label{
			Show:      true,
			Formatter: "{b}: {c}",
		}),
	)
	pie.Render(w)
}
