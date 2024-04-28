package dataframe

type Window struct {
	Partition []string
	Order     []Order
	frame     *windowFrame
}

type windowFrame struct {
	start int
	end   int
}

func (w Window) OrderBy(order ...Order) Window {
	return Window{w.Partition, order, w.frame}
}

func (w Window) PartitionBy(cols ...string) Window {
	return Window{cols, w.Order, w.frame}
}

func (w Window) RowsBetween(start, end int) Window {
	return Window{w.Partition, w.Order, &windowFrame{start, end}}
}


