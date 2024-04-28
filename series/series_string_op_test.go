package series_test

import (
	"reflect"
	"testing"

	"github.com/IzzelAliz/gota/series"
)

func TestString(t *testing.T) {
	table := []struct {
		expect series.Series
		actual series.Series
		reason string
	}{
		{
			series.Strings([]string{"a", "b"}),
			series.Strings([]string{"a.b", "b.c"}).SubstringIndex(".", 1),
			"substring index left",
		},
		{
			series.Strings([]string{"b", "c"}),
			series.Strings([]string{"a.b", "b.c"}).SubstringIndex(".", -1),
			"substring index right",
		},
		{
			series.Bools([]bool{true, false}),
			series.Strings([]string{"a.b", "ab"}).Contains("."),
			"string contains",
		},
		{
			series.Bools([]bool{false, true}),
			series.Strings([]string{"a.b", "ab"}).RegexMatch("^\\w+$"),
			"regex match",
		},
		{
			series.Strings([]string{"1(a)", "2(b)"}),
			series.Strings([]string{"a(1)", "b(2)"}).RegexReplace("^(\\w)\\((\\w)\\)$", "$2($1)"),
			"regex replace",
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
