package main

import (
	"container/heap"
	"math"
)

type QueueNode struct {
	X int
	Y int
	distance int
	index int

}

type PriorityQueue []*QueueNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].distance > pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	point := x.(*QueueNode)
	point.index = n
	*pq = append(*pq, point)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	point := old[n-1]
	old[n-1] = nil   // avoid memory leak
	point.index = -1 // for safety
	*pq = old[0 : n-1]
	return point
}

func (pq *PriorityQueue) update(queueNode *QueueNode, distance int) {
	queueNode.distance = distance
	heap.Fix(pq, queueNode.index)
}

type VertexQueue struct {
	priorityQueue *PriorityQueue
	vertexMap map[Point]*QueueNode
}

func NewVertexQueue(points []Point) *VertexQueue {
	priorityQueue := make(PriorityQueue, 0)
	vertexMap := make(map[Point]*QueueNode)

	heap.Init(&priorityQueue)
	for i, point := range points {
		distance := math.MaxInt
		if point.X == 0 && point.Y == 0 {
			distance = 0
		}
		node := &QueueNode{
			X: point.X,
			Y: point.Y,
			distance: distance * -1,
			index: i,
		}
		vertexMap[point] = node
		heap.Push(&priorityQueue, node)
	}

	return &VertexQueue{
		priorityQueue: &priorityQueue,
		vertexMap: vertexMap,
	}
}

func (vq *VertexQueue) Push(point Point, distance int) {
	queueNode := &QueueNode{
		X: point.Y,
		Y: point.Y,
		distance: distance * -1,
	}
	heap.Push(vq.priorityQueue, queueNode)
}

func (vq *VertexQueue) Pop() *Point {
	if len(*vq.priorityQueue) == 0 {
		return nil
	}
	entry := heap.Pop(vq.priorityQueue)
	if entry == nil {
		return nil
	}
	node := entry.(*QueueNode)
	return &Point{
		X: node.X,
		Y: node.Y,
	}
}

func (vq *VertexQueue) Update(point Point, distance int) {
	queueNode := vq.vertexMap[point]
	vq.priorityQueue.update(queueNode, distance * -1)
}