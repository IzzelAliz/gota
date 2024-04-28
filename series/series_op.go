package series

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
)

func (s Series) Add(val interface{}) Series {
	switch v := val.(type) {
	case int:
		return s.addVal(float64(v))
	case float64:
		return s.addVal(v)
	case Series:
		return s.addCol(v)
	default:
		return s.errorf("unknown add type: %T", val)
	}
}

func (s Series) addVal(val float64) Series {
	return s.Copy().Map(func(e Element) Element {
		e.Set(e.Float() + val)
		return e
	})
}

func (s Series) addCol(other Series) Series {
	if s.Len() != other.Len() {
		return s.errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
	}
	if s.Type() == Int && other.Type() == Int {
		a, _ := s.Int()
		b, _ := other.Int()
		for i := range a {
			a[i] = a[i] + b[i]
		}
		return New(a, Int, s.Name+"+"+other.Name)
	} else {
		return New(floats.AddTo(make([]float64, s.Len()), s.Float(), other.Float()), Float, s.Name+"+"+other.Name)
	}
}

func (s Series) Sub(val interface{}) Series {
	switch v := val.(type) {
	case int:
		return s.subVal(float64(v))
	case float64:
		return s.subVal(v)
	case Series:
		return s.subCol(v)
	default:
		return s.errorf("unknown sub type: %T", val)
	}
}

func (s Series) subVal(val float64) Series {
	return s.Copy().Map(func(e Element) Element {
		e.Set(e.Float() - val)
		return e
	})
}

func (s Series) subCol(other Series) Series {
	if s.Len() != other.Len() {
		return s.errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
	}
	if s.Type() == Int && other.Type() == Int {
		a, _ := s.Int()
		b, _ := other.Int()
		for i := range a {
			a[i] = a[i] - b[i]
		}
		return New(a, Int, s.Name+"-"+other.Name)
	} else {
		return New(floats.SubTo(make([]float64, s.Len()), s.Float(), other.Float()), Float, s.Name+"-"+other.Name)
	}
}

func (s Series) Mul(val interface{}) Series {
	switch v := val.(type) {
	case int:
		return s.mulVal(float64(v))
	case float64:
		return s.mulVal(v)
	case Series:
		return s.mulCol(v)
	default:
		return s.errorf("unknown mul type: %T", val)
	}
}

func (s Series) mulVal(val float64) Series {
	return s.Copy().Map(func(e Element) Element {
		e.Set(e.Float() * val)
		return e
	})
}

func (s Series) mulCol(other Series) Series {
	if s.Len() != other.Len() {
		return s.errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
	}
	if s.Type() == Int && other.Type() == Int {
		a, _ := s.Int()
		b, _ := other.Int()
		for i := range a {
			a[i] = a[i] * b[i]
		}
		return New(a, Int, s.Name+"*"+other.Name)
	} else {
		return New(floats.MulTo(make([]float64, s.Len()), s.Float(), other.Float()), Float, s.Name+"*"+other.Name)
	}
}
func (s Series) Div(val interface{}) Series {
	switch v := val.(type) {
	case int:
		return s.divVal(float64(v))
	case float64:
		return s.divVal(v)
	case Series:
		return s.divCol(v)
	default:
		return s.errorf("unknown mul type: %T", val)
	}
}

func (s Series) divVal(val float64) Series {
	return s.Copy().Map(func(e Element) Element {
		e.Set(e.Float() / val)
		return e
	})
}

func (s Series) divCol(other Series) Series {
	if s.Len() != other.Len() {
		return s.errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
	}
	if s.Type() == Int && other.Type() == Int {
		a, _ := s.Int()
		b, _ := other.Int()
		for i := range a {
			a[i] = a[i] / b[i]
		}
		return New(a, Int, s.Name+"/"+other.Name)
	} else {
		return New(floats.DivTo(make([]float64, s.Len()), s.Float(), other.Float()), Float, s.Name+"/"+other.Name)
	}
}

func (s Series) Eq(val interface{}) Series {
	return s.Compare(Eq, val)
}

func (s Series) Neq(val interface{}) Series {
	return s.Compare(Neq, val)
}

func (s Series) Greater(val interface{}) Series {
	return s.Compare(Greater, val)
}

func (s Series) GreaterEq(val interface{}) Series {
	return s.Compare(GreaterEq, val)
}

func (s Series) Less(val interface{}) Series {
	return s.Compare(Less, val)
}

func (s Series) LessEq(val interface{}) Series {
	return s.Compare(LessEq, val)
}

func (s Series) In(val interface{}) Series {
	return s.Compare(In, val)
}

func (s Series) Cast(t Type) Series {
	if t == s.t {
		return s
	}
	var (
		newval interface{}
		err    error
	)
	switch t {
	case String:
		newval = s.Records()
	case Float:
		newval = s.Float()
	case Int:
		newval, err = s.Int()
	case Bool:
		newval, err = s.Bool()
	}
	r := New(newval, t, fmt.Sprintf("%v(%v)", t, s.Name))
	r.Err = err
	return r
}

// ZeroNA convert NA values to zero
func (s Series) ZeroNA(t Type) Series {
	var zero interface{} = 0
	if t == String {
		zero = ""
	}
	return s.Copy().Map(func(e Element) Element {
		if e.IsNA() {
			e.Set(zero)
		}
		return e
	})
}

func (s Series) And(other Series) Series {
	return s.Zip(other, Bool, func(a, b Element) interface{} {
		ba, err := a.Bool()
		if err != nil {
			return nil
		}
		bb, err := b.Bool()
		if err != nil {
			return nil
		}
		return ba && bb
	}).As(fmt.Sprintf("(%v && %v)", s.Name, other.Name))
}

func (s Series) Or(other Series) Series {
	return s.Zip(other, Bool, func(a, b Element) interface{} {
		ba, err := a.Bool()
		if err != nil {
			return nil
		}
		bb, err := b.Bool()
		if err != nil {
			return nil
		}
		return ba || bb
	}).As(fmt.Sprintf("(%v || %v)", s.Name, other.Name))
}
