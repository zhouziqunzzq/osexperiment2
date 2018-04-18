package main

const (
	PRODUCER = 1
	CONSUMER = 2
)

type Process struct {
	PType int
	Item  int
}

func CopyProcess(p Process) *Process {
	return &Process{
		PType: p.PType,
		Item:  p.Item,
	}
}

func NewProcess(ptype int, item int) *Process {
	return &Process{
		PType: ptype,
		Item:  item,
	}
}
