package stack

type MyQueue struct {
	stIn  *Stack
	stOut *Stack
}

func (q *MyQueue) Push(val int) {
	q.stIn.Push(val)
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		stIn:  NewStack(),
		stOut: NewStack(),
	}
}

func (q *MyQueue) Pop() int {

	if q.stOut.Len() == 0 {
		for q.stIn.Len() != 0 {
			q.stOut.Push(q.stIn.Pop())
		}
	}

	return q.stOut.Pop().(int)
}

func (q *MyQueue) Peek() int {
	if q.stOut.Len() == 0 {
		for q.stIn.Len() != 0 {
			q.stOut.Push(q.stIn.Pop())
		}
	}

	return q.stOut.Peek().(int)
}

func (q *MyQueue) Empty() bool {
	if q.stIn.Len() == 0 && q.stOut.Len() == 0 {
		return true
	}

	return false
}
