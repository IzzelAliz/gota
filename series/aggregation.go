package series

type AggregationFunc func(Series) interface{}

var (
	Aggregation_Max    AggregationFunc = func(s Series) interface{} { return s.Max() }
	Aggregation_Min    AggregationFunc = func(s Series) interface{} { return s.Min() }
	Aggregation_Avg    AggregationFunc = func(s Series) interface{} { return s.Mean() }
	Aggregation_Sum    AggregationFunc = func(s Series) interface{} { return s.Sum() }
	Aggregation_StdDev AggregationFunc = func(s Series) interface{} { return s.StdDev() }

	Aggregation_Count         AggregationFunc = func(s Series) interface{} { return s.Len() }
	Aggregation_CountDistinct AggregationFunc = func(s Series) interface{} {
		distinct := make(map[interface{}]struct{})
		for i := 0; i < s.Len(); i++ {
			distinct[s.Elem(i).Val()] = struct{}{}
		}
		return len(distinct)
	}

	Aggregation_First AggregationFunc = func(s Series) interface{} {
		if s.Len() > 0 {
			return s.Val(0)
		}
		return nil
	}
)

func AggregationQuantile(p float64) AggregationFunc {
	return func(s Series) interface{} {
		return s.Quantile(p)
	}
}
