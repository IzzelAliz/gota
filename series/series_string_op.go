package series

import (
	"fmt"
	"regexp"
	"strings"
)

func (s Series) SubstringIndex(delim string, count int) Series {
	return s.MapAs(String, func(e Element) interface{} {
		str := e.String()
		n := count
		if n > 0 {
			idx := -1
			for n > 0 {
				i := strings.Index(str[idx+1:], delim)
				if i >= 0 {
					n--
					idx = idx + i + 1
				} else {
					return str
				}
			}
			return str[:idx]
		} else {
			idx := len(str) - len(delim) + 1
			n = -n
			for n > 0 {
				idx = strings.LastIndex(str[:idx-1], delim)
				if idx >= 0 {
					n--
				} else {
					return str
				}
			}
			return str[idx+len(delim):]
		}
	}).As(fmt.Sprintf("SubstringIndex(%v, %v, %v)", s.Name, delim, count))
}

func (s Series) Contains(substr interface{}) Series {
	switch v := substr.(type) {
	case string:
		return s.MapAs(Bool, func(e Element) interface{} {
			return strings.Contains(e.String(), v)
		}).As(fmt.Sprintf("Contains(%v, %v)", s.Name, substr))
	case Series:
		return s.Zip(v, Bool, func(a, b Element) interface{} {
			return strings.Contains(a.String(), b.String())
		}).As(fmt.Sprintf("Contains(%v, %v)", s.Name, substr))
	default:
		return s.errorf("unknown substr type: %T", substr)
	}
}

func (s Series) RegexMatch(regex string) Series {
	re := regexp.MustCompile(regex)
	return s.MapAs(Bool, func(e Element) interface{} {
		return re.MatchString(e.String())
	}).As(fmt.Sprintf("RegexMatch(%v, %v)", s.Name, regex))
}

func (s Series) RegexReplace(regex string, replacement string) Series {
	re := regexp.MustCompile(regex)
	return s.MapAs(String, func(e Element) interface{} {
		return re.ReplaceAllString(e.String(), replacement)
	}).As(fmt.Sprintf("RegexReplace(%v, %v, %v)", s.Name, regex, replacement))
}
