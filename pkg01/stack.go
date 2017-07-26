package pkg01

type Stack struct {
	data []int
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}

func (s Stack) Top() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	return s.data[len(s.data)-1], true
}
