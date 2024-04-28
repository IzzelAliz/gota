package series

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
)

func (s Series) AddVal(val float64) Series {
	return s.Copy().Map(func(e Element) Element {
		e.Set(e.Float() + val)
		return e
	})
}

func (s Series) Add(other Series) Series {
	if s.Len() != other.Len() {
		r := s.Empty()
		r.Err = fmt.Errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
		return r
	}
	if s.Type() == Int && other.Type() == Int {
		a, _ := s.Int()
		b, _ := other.Int()
		return New(addToInt(make([]int, s.Len()), a, b), Int, s.Name+"+"+other.Name)
	} else {
		return New(floats.AddTo(make([]float64, s.Len()), s.Float(), other.Float()), Float, s.Name+"+"+other.Name)
	}
}

func addToInt(dst, a, b []int) []int {
	for i := range dst {
		dst[i] = a[i] + b[i]
	}
	return dst
}

func (s Series) SubVal(val float64) Series {
	return s.Copy().Map(func(e Element) Element {
		e.Set(e.Float() - val)
		return e
	})
}

func (s Series) Sub(other Series) Series {
	if s.Len() != other.Len() {
		r := s.Empty()
		r.Err = fmt.Errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
		return r
	}
	if s.Type() == Int && other.Type() == Int {
		a, _ := s.Int()
		b, _ := other.Int()
		return New(subToInt(make([]int, s.Len()), a, b), Int, s.Name+"-"+other.Name)
	} else {
		return New(floats.SubTo(make([]float64, s.Len()), s.Float(), other.Float()), Float, s.Name+"-"+other.Name)
	}
}

func subToInt(dst, a, b []int) []int {
	for i := range dst {
		dst[i] = a[i] - b[i]
	}
	return dst
}

func (s Series) MulVal(val float64) Series {
	return s.Copy().Map(func(e Element) Element {
		e.Set(e.Float() * val)
		return e
	})
}

func (s Series) Mul(other Series) Series {
	if s.Len() != other.Len() {
		r := s.Empty()
		r.Err = fmt.Errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
		return r
	}
	if s.Type() == Int && other.Type() == Int {
		a, _ := s.Int()
		b, _ := other.Int()
		return New(mulToInt(make([]int, s.Len()), a, b), Int, s.Name+"*"+other.Name)
	} else {
		return New(floats.MulTo(make([]float64, s.Len()), s.Float(), other.Float()), Float, s.Name+"*"+other.Name)
	}
}

func mulToInt(dst, a, b []int) []int {
	for i := range dst {
		dst[i] = a[i] * b[i]
	}
	return dst
}

func (s Series) DivVal(val float64) Series {
	return s.Copy().Map(func(e Element) Element {
		e.Set(e.Float() / val)
		return e
	})
}

func (s Series) Div(other Series) Series {
	if s.Len() != other.Len() {
		r := s.Empty()
		r.Err = fmt.Errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
		return r
	}
	if s.Type() == Int && other.Type() == Int {
		a, _ := s.Int()
		b, _ := other.Int()
		return New(divToInt(make([]int, s.Len()), a, b), Int, s.Name+"/"+other.Name)
	} else {
		return New(floats.DivTo(make([]float64, s.Len()), s.Float(), other.Float()), Float, s.Name+"/"+other.Name)
	}
}

func divToInt(dst, a, b []int) []int {
	for i := range dst {
		dst[i] = a[i] / b[i]
	}
	return dst
}

func (s Series) Eq(other Series) Series {
	if s.Len() != other.Len() {
		r := s.Empty()
		r.Err = fmt.Errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
		return r
	}
	r := New(make([]bool, s.Len()), Bool, s.Name+"=="+other.Name)
	for i := 0; i < s.Len(); i++ {
		r.Elem(i).Set(s.Elem(i).Eq(other.Elem(i)))
	}
	return r
}

func (s Series) Neq(other Series) Series {
	if s.Len() != other.Len() {
		r := s.Empty()
		r.Err = fmt.Errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
		return r
	}
	r := New(make([]bool, s.Len()), Bool, s.Name+"!="+other.Name)
	for i := 0; i < s.Len(); i++ {
		r.Elem(i).Set(s.Elem(i).Neq(other.Elem(i)))
	}
	return r
}

func (s Series) Greater(other Series) Series {
	if s.Len() != other.Len() {
		r := s.Empty()
		r.Err = fmt.Errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
		return r
	}
	r := New(make([]bool, s.Len()), Bool, s.Name+">"+other.Name)
	for i := 0; i < s.Len(); i++ {
		r.Elem(i).Set(s.Elem(i).Greater(other.Elem(i)))
	}
	return r
}

func (s Series) GreaterEq(other Series) Series {
	if s.Len() != other.Len() {
		r := s.Empty()
		r.Err = fmt.Errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
		return r
	}
	r := New(make([]bool, s.Len()), Bool, s.Name+">="+other.Name)
	for i := 0; i < s.Len(); i++ {
		r.Elem(i).Set(s.Elem(i).GreaterEq(other.Elem(i)))
	}
	return r
}

func (s Series) Less(other Series) Series {
	if s.Len() != other.Len() {
		r := s.Empty()
		r.Err = fmt.Errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
		return r
	}
	r := New(make([]bool, s.Len()), Bool, s.Name+"<"+other.Name)
	for i := 0; i < s.Len(); i++ {
		r.Elem(i).Set(s.Elem(i).Less(other.Elem(i)))
	}
	return r
}

func (s Series) LessEq(other Series) Series {
	if s.Len() != other.Len() {
		r := s.Empty()
		r.Err = fmt.Errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
		return r
	}
	r := New(make([]bool, s.Len()), Bool, s.Name+"<="+other.Name)
	for i := 0; i < s.Len(); i++ {
		r.Elem(i).Set(s.Elem(i).LessEq(other.Elem(i)))
	}
	return r
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
	if s.Len() != other.Len() {
		r := s.Empty()
		r.Err = fmt.Errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
		return r
	}
	if s.t != Bool || other.t != Bool {
		r := s.Empty()
		r.Err = fmt.Errorf("type is not bool: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
		return r
	}
	a, err := s.Bool()
	if err != nil {
		r := s.Empty()
		r.Err = err
		return r
	}
	b, err := other.Bool()
	if err != nil {
		r := s.Empty()
		r.Err = err
		return r
	}
	result := make([]bool, s.Len())
	for i := range result {
		result[i] = a[i] && b[i]
	}
	return New(result, Bool, fmt.Sprintf("(%v && %v)", s.Name, other.Name))
}

func (s Series) Or(other Series) Series {
	if s.Len() != other.Len() {
		r := s.Empty()
		r.Err = fmt.Errorf("index dimensions mismatch: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
		return r
	}
	if s.t != Bool || other.t != Bool {
		r := s.Empty()
		r.Err = fmt.Errorf("type is not bool: %s %v, %s %v", s.Name, s.Len(), other.Name, other.Len())
		return r
	}
	a, err := s.Bool()
	if err != nil {
		r := s.Empty()
		r.Err = err
		return r
	}
	b, err := other.Bool()
	if err != nil {
		r := s.Empty()
		r.Err = err
		return r
	}
	result := make([]bool, s.Len())
	for i := range result {
		result[i] = a[i] || b[i]
	}
	return New(result, Bool, fmt.Sprintf("(%v || %v)", s.Name, other.Name))
}
