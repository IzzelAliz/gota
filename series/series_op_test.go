package series_test

import (
	"reflect"
	"testing"

	"github.com/IzzelAliz/gota/series"
)

func TestArith(t *testing.T) {
	table := []struct {
		expect series.Series
		actual series.Series
		reason string
	}{
		{
			series.Ints([]int{2, 3, 4}),
			series.Ints([]int{1, 2, 3}).Add(series.Ints([]int{1, 1, 1})),
			"Add",
		},
		{
			series.Ints([]int{2, 3, 4}),
			series.Ints([]int{1, 2, 3}).AddVal(1),
			"Add",
		},
		{
			series.Ints([]int{0, 1, 2}),
			series.Ints([]int{1, 2, 3}).Sub(series.Ints([]int{1, 1, 1})),
			"Sub",
		},
		{
			series.Ints([]int{2, 4, 6}),
			series.Ints([]int{1, 2, 3}).Mul(series.Ints([]int{2, 2, 2})),
			"Mul",
		},
		{
			series.Ints([]int{0, 1, 1}),
			series.Ints([]int{1, 2, 3}).Div(series.Ints([]int{2, 2, 2})),
			"Div",
		},
		{
			series.Floats([]float64{2.5, 3.5, 4.5}),
			series.Floats([]float64{1.5, 2.5, 3.5}).Add(series.Floats([]float64{1.0, 1.0, 1.0})),
			"Add",
		},
		{
			series.Floats([]float64{0.5, 1.5, 2.5}),
			series.Floats([]float64{1.5, 2.5, 3.5}).Sub(series.Floats([]float64{1.0, 1.0, 1.0})),
			"Sub",
		},
		{
			series.Floats([]float64{3.0, 5.0, 7.0}),
			series.Floats([]float64{1.5, 2.5, 3.5}).Mul(series.Floats([]float64{2.0, 2.0, 2.0})),
			"Mul",
		},
		{
			series.Floats([]float64{0.75, 1.25, 1.75}),
			series.Floats([]float64{1.5, 2.5, 3.5}).Div(series.Floats([]float64{2.0, 2.0, 2.0})),
			"Div",
		},
	}
	for _, ent := range table {
		if !reflect.DeepEqual(ent.expect.As(""), ent.actual.As("")) {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				ent.reason, ent.expect, ent.actual,
			)
		}
	}
}

func TestMisc(t *testing.T) {
	table := []struct {
		expect series.Series
		actual series.Series
		reason string
	}{
		{
			series.Floats([]float64{1, 2, 3}),
			series.Ints([]int{1, 2, 3}).Cast(series.Float),
			"Cast to float",
		},
		{
			series.Strings([]string{"1", "2", "3"}),
			series.Ints([]int{1, 2, 3}).Cast(series.String),
			"Cast to string",
		},
		{
			series.Floats([]float64{1, 2, 3}),
			series.Strings([]string{"1", "2", "3"}).Cast(series.Float),
			"Cast string to flaot",
		},
		{
			series.Strings([]string{"1", "2", "3"}),
			series.Ints([]int{1, 2, 3}).MapAs(series.String, func(e series.Element) interface{} { return e.String() }),
			"MapAs string",
		},
	}
	for _, ent := range table {
		if !reflect.DeepEqual(ent.expect.As(""), ent.actual.As("")) {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				ent.reason, ent.expect, ent.actual,
			)
		}
	}
}

func TestBool(t *testing.T) {
	table := []struct {
		expect series.Series
		actual series.Series
		reason string
	}{
		{
			series.Bools([]bool{false, false, true}),
			series.Ints([]int{1, 2, 3}).Greater(series.Lit(2, 3)),
			"Greater",
		},
		{
			series.Bools([]bool{false, true, true}),
			series.Ints([]int{1, 2, 3}).GreaterEq(series.Lit(2, 3)),
			"GreaterEq",
		},
		{
			series.Bools([]bool{true, true, false}),
			series.Ints([]int{1, 2, 3}).Less(series.Lit(3, 3)),
			"Less",
		},
		{
			series.Bools([]bool{true, true, true}),
			series.Ints([]int{1, 2, 3}).LessEq(series.Lit(3, 3)),
			"LessEq",
		},
	}
	for _, ent := range table {
		if !reflect.DeepEqual(ent.expect.As(""), ent.actual.As("")) {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				ent.reason, ent.expect, ent.actual,
			)
		}
	}
}
