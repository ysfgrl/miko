package core

import "slices"

type Queue[Type any] struct {
	list []Type
}

func NewQueue[Type any](size int) *Queue[Type] {
	return &Queue[Type]{}
}

func (q *Queue[Type]) IsEmpty() bool {
	return len(q.list) == 0
}

func (q *Queue[Type]) Length() int {
	return len(q.list)
}
func (q *Queue[Type]) Push(value Type) {
	q.list = slices.Insert(q.list, 0, value)
}
func (q *Queue[Type]) Pop() Type {
	length := q.Length()
	res := q.list[length-1]
	q.list = q.list[:length-1]
	return res
}
func (q *Queue[Type]) For(f func(item Type)) {
	for _, t := range q.list {
		f(t)
	}
}
func (q *Queue[Type]) First() Type {
	return q.list[q.Length()-1]
}
func (q *Queue[Type]) Last() Type {
	return q.list[0]
}
