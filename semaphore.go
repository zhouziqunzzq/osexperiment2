package main

type Semaphore struct {
	Count int
	Q     Queue
}

// If blocked, return false. Else return true
func (s *Semaphore) P(process Process) bool {
	s.Count--
	if s.Count < 0 {
		s.Q.Push(process)
		return false
	}
	return true
}

func (s *Semaphore) V() Process {
	s.Count++
	if s.Count <= 0 {
		return s.Q.Pop().(Process)
	}
	return Process{PType: -1}
}
