type Task struct {
	Task byte
	Freq int
}

type MaxHeap []Task

func (mh MaxHeap) Len() int           { return len(mh) }
func (mh MaxHeap) Less(i, j int) bool { return mh[i].Freq > mh[j].Freq }
func (mh MaxHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }

func (mh *MaxHeap) Push(item interface{}) {
	*mh = append(*mh, item.(Task))
}

func (mh *MaxHeap) Pop() interface{} {
	old := *mh
	lastIndex := len(*mh) - 1

	last := old[lastIndex]
	*mh = old[:lastIndex]
	return last
}

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	tasksFreq := make(map[byte]int)
	for _, task := range tasks {
		tasksFreq[task]++
	}

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for task, freq := range tasksFreq {
		heap.Push(maxHeap, Task{Task: task, Freq: freq})
	}

	var time int

	O(N)

	for maxHeap.Len() > 0 {
		tasksLeft := []Task{}
		cycle := n + 1

		for cycle > 0 && maxHeap.Len() > 0 {
			task := heap.Pop(maxHeap).(Task)
			time++
			cycle--
			task.Freq--

			if task.Freq > 0 {
				tasksLeft = append(tasksLeft, task)
			}
		}

		for _, task := range tasksLeft {
			heap.Push(maxHeap, task)
		}

		if maxHeap.Len() > 0 {
			time += cycle
		}
	}

	return time
}

// T = number of unique tasks
// time: O(TlogT)
// space: O(T)