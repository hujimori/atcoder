package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	n, k := nextInt(), nextInt()
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = nextInt()
	}

	x := make([]int, n*k)
	for i := 0; i < n; i++ {
		x[i] = A[i]
	}

	for i := n; i < n*k; i++ {
		x[i] = x[i-n]
	}

	fmt.Println(x)

	q := NewQueue()

	for i := 0; i < n*k; i++ {

	}
}

type QueueData struct {
	data interface{}
	prev *QueueData
}

type Queue struct {
	front *QueueData
	last  *QueueData
	top   int
}

func NewQueue() *Queue {
	return &Queue{top: 0}
}

func (q *Queue) Enqueue(data interface{}) {
	qd := &QueueData{data: data, prev: nil}
	if q.top == 0 {
		q.front = qd
		q.last = q.front
		q.top++
		return
	}

	q.last.prev = qd
	q.last = qd
	q.top++
}

func (q *Queue) Dequeue() interface{} {
	if q.top > 0 {
		d := q.front.data
		q.front = q.front.prev
		q.top--

		return d
	}
	return nil
}

func (q *Queue) Front() interface{} {
	return q.front.data
}

func (q *Queue) Last() interface{} {
	return q.last.data
}

func (q *Queue) IsEmpty() bool {
	if q.top == 0 {
		return true
	}
	return false
}

func (q *Queue) Size() int {
	return q.top
}
