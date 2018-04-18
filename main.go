package main

import (
	"fmt"
)

const (
	ClearScreenCMD = "\033[2J\033[1;1H"
)

var (
	choice      rune
	buffer      []int
	count       = 1
	emptyBuffer Semaphore
	fullBuffer  Semaphore
	mutex       Semaphore
	readP       = 0
	writeP      = 0
	readyQueue  = NewQueue()
)

func Clear() {
	fmt.Print(ClearScreenCMD)
	return
}

func Producer() {
	process := NewProcess(PRODUCER, count)
	count++

	if !emptyBuffer.P(*process) {
		return
	}

	if !mutex.P(*process) {
		return
	}

	readyQueue.Push(*process)
}

func Consumer() {
	process := NewProcess(CONSUMER, 0)

	if !fullBuffer.P(*process) {
		return
	}

	if !mutex.P(*process) {
		return
	}

	readyQueue.Push(*process)
}

func Continue() {
	runningProcess := readyQueue.Pop().(Process)
	switch runningProcess.PType {
	case PRODUCER:
		WriteBuffer(runningProcess.Item)

		resumedProcess := mutex.V()
		if resumedProcess.PType != -1 {
			readyQueue.Push(resumedProcess)
		}

		resumedProcess = fullBuffer.V()
		if resumedProcess.PType == CONSUMER {
			//ReadBuffer()
			//emptyBuffer.V()
			mutex.P(resumedProcess)
		}
	case CONSUMER:
		ReadBuffer()

		resumedProcess := mutex.V()
		if resumedProcess.PType != -1 {
			readyQueue.Push(resumedProcess)
		}

		resumedProcess = emptyBuffer.V()
		if resumedProcess.PType == PRODUCER {
			//WriteBuffer(resumedProcess.Item)
			//fullBuffer.V()
			mutex.P(resumedProcess)
		}
	}
}

func PrintProcess(p interface{}) {
	if p.(Process).PType == PRODUCER {
		fmt.Printf("(%v, P) ", p.(Process).Item)
	} else {
		fmt.Printf("(null, C) ")
	}
}

func PrintAll() {
	fmt.Printf("================Buffer(%v)==============\n", len(buffer))
	for i := 0; i < len(buffer); i++ {
		fmt.Printf("%v ", buffer[i])
	}
	fmt.Printf("\n")
	fmt.Printf("================Mutex(%v)===============\n", mutex.Count)
	if mutex.Q.IsEmpty() {
		fmt.Println("Empty")
	} else {
		for i := 0; i < len(mutex.Q.Queue); i++ {
			PrintProcess(mutex.Q.Queue[i])
		}
		fmt.Print("\n")
	}
	fmt.Printf("=============EmptyBuffer(%v)============\n", emptyBuffer.Count)
	if emptyBuffer.Q.IsEmpty() {
		fmt.Println("Empty")
	} else {
		for i := 0; i < len(emptyBuffer.Q.Queue); i++ {
			PrintProcess(emptyBuffer.Q.Queue[i])
		}
		fmt.Print("\n")
	}
	fmt.Printf("=============FullBuffer(%v)=============\n", fullBuffer.Count)
	if fullBuffer.Q.IsEmpty() {
		fmt.Println("Empty")
	} else {
		for i := 0; i < len(fullBuffer.Q.Queue); i++ {
			PrintProcess(fullBuffer.Q.Queue[i])
		}
		fmt.Print("\n")
	}
	fmt.Printf("=============ReadyQueue(%v)==============\n", readyQueue.Count())
	if readyQueue.IsEmpty() {
		fmt.Println("Empty")
	} else {
		for i := 0; i < len(readyQueue.Queue); i++ {
			PrintProcess(readyQueue.Queue[i])
		}
		fmt.Print("\n")
	}
}

func WriteBuffer(num int) {
	buffer[writeP] = num
	writeP = (writeP + 1) % len(buffer)
}

func ReadBuffer() (num int) {
	num = buffer[readP]
	buffer[readP] = 0
	readP = (readP + 1) % len(buffer)
	return
}

func main() {
	fmt.Println("Simulation start")
	// Init buffer
	fmt.Println("Please input buffer size: ")
	var buffSize int
	fmt.Scanf("%v", &buffSize)
	buffer = make([]int, buffSize)
	// Init semaphores
	emptyBuffer.Count = buffSize
	fullBuffer.Count = 0
	mutex.Count = 1
	// Start main loop
	for choice != 'q' {
		//Clear()
		PrintAll()
		fmt.Println("p: Call Producer, c: Call Consumer, v: Continue, q: Exit")
		fmt.Scanf("%c\n", &choice)
		switch choice {
		case 'p':
			Producer()
		case 'c':
			Consumer()
		case 'v':
			if readyQueue.IsEmpty() {
				fmt.Println("Error: No running process!")
				continue
			}
			Continue()
		case 'q':
			break
		default:
			continue
		}
	}
	fmt.Println("Simulation end")
}
