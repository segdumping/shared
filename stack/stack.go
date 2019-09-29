package stack

type Stack []interface{}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *Stack) Pop() interface{} {
	l := len(*s)
	if l == 0 {
		return nil
	}

	val := (*s)[l-1]
	*s = (*s)[0 : l-1]
	return val
}

func (s Stack) Top() interface{} {
	l := len(s)
	if l == 0 {
		return nil
	}

	return s[l-1]
}

func (s Stack) Empty() bool {
	return len(s) == 0
}
