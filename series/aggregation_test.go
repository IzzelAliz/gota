package series_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/IzzelAliz/gota/series"
)

func TestAgg(t *testing.T) {
	table := []struct {
		expect any
		actual any
		reason string
	}{
		{
			1.0,
			series.Aggregation_Avg(series.Ints([]int{0, 1, 2})),
			"avg",
		},
		{
			0.0,
			series.Aggregation_Min(series.Ints([]int{0, 1, 2})),
			"min",
		},
		{
			2.0,
			series.Aggregation_Max(series.Ints([]int{0, 1, 2})),
			"max",
		},
		{
			3.0,
			series.Aggregation_Sum(series.Ints([]int{0, 1, 2})),
			"sum",
		},
		{
			3,
			series.Aggregation_Count(series.Ints([]int{0, 1, 2})),
			"count",
		},
		{
			2,
			series.Aggregation_CountDistinct(series.Ints([]int{0, 1, 1})),
			"countdistinct",
		},
		{
			2,
			series.Aggregation_CountDistinct(series.Strings([]string{"0", "NaN", "NaN"})),
			"countdistinct",
		},
		{
			3,
			series.Aggregation_CountDistinct(series.Floats([]float64{0, 1, math.NaN()})),
			"countdistinct",
		},
	}
	for _, ent := range table {
		if !reflect.DeepEqual(ent.expect, ent.actual) {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				ent.reason, ent.expect, ent.actual,
			)
		}
	}
}
