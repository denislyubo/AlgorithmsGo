package dailytemperatures

type stack struct {
	data []interface{}
}

func (s *stack) Empty() bool {
	return len(s.data) == 0
}

func (s *stack) Pop() {
	s.data = s.data[:len(s.data)-1]
}

func (s *stack) Push(item interface{}) {
	s.data = append(s.data, item)
}

func (s *stack) Top() interface{} {
	return s.data[len(s.data)-1]
}

type Pair struct {
	key   int
	value int
}

func dailyTemperatures(temperatures []int) []int {
	l := len(temperatures)

	if l == 0 || l == 1 {
		return []int{0}
	}

	ret := make([]int, l)

	s := stack{}

	for i := l - 1; i >= 0; i-- {
		cur := temperatures[i]
		if i == l-1 {
			ret[i] = 0
			s.Push(Pair{key: i, value: cur})

			continue
		}

		for !s.Empty() {
			top := s.Top()
			if top.(Pair).value > cur {
				ret[i] = top.(Pair).key - i
				s.Push(Pair{key: i, value: cur})
				break
			}

			s.Pop()
		}

		if s.Empty() {
			ret[i] = 0
			s.Push(Pair{key: i, value: cur})
		}
	}

	return ret
}
