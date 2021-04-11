package queue

import "errors"

type ElementType interface{}
type Array []ElementType

func New(cap int) *Queue {
	data := make(Array, cap)
	return &Queue{
		Capacity: cap,
		Front:    1,
		Rear:     0,
		Size:     0,
		Data:     &data,
	}

}

type QueueInterface interface {
	IsEmpty() bool
	IsFull() bool
	MakeEmpty()
	EnQueue(x ElementType) error
	DeQueue() (ElementType, error)
}

/* Queue implementation with Slice */
type Queue struct {
	Capacity int
	Front    int
	Rear     int
	Size     int
	Data     *Array
}

/* Return true if queue is empty */
func (s *Queue) IsEmpty() bool {
	return s.Size == 0
}

/* Return true if the queue is full*/
func (s *Queue) IsFull() bool {
	return s.Size == s.Capacity
}

/* Make queue Empty */
func (s *Queue) MakeEmpty() {
	s.Front = 1
	s.Rear = 0
	s.Size = 0
	empty_slice := (*s.Data)[:0]
	s.Data = &empty_slice
}

/* Push x enter queue at rear*/
func (s *Queue) EnQueue(x ElementType) error {
	if s.IsFull() {
		return errors.New("full queue")
	}
	s.Rear = s.succ(s.Rear)
	(*s.Data)[s.Rear] = x
	s.Size++
	return nil

}

/*Get and delete x from queue at front */
func (s *Queue) DeQueue() (ElementType, error) {
	if s.IsEmpty() {
		return nil, errors.New("empty queue")
	}
	x := (*s.Data)[s.Front]
	s.Front = s.succ(s.Front)
	s.Size--
	return x, nil
}

/* Enhance value circulate  */
func (s *Queue) succ(value int) int {
	if value++; s.Capacity == value {
		value = 0
	}
	return value
}
