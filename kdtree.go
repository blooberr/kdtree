package kdtree

import (
	"log"
	"math"

	"github.com/blooberr/kdtree/bpqueue"
	"github.com/blooberr/kdtree/point"
)

type KDTree struct {
	Tree *Node
	Root *Node

	Dimension int
	Size      int
}

type Node struct {
	Point *point.Point
	Value string
	Level int

	Left  *Node
	Right *Node
}

func New() *KDTree {
	return &KDTree{
		Size: 0,
	}
}

// Inserts a point into the KDTree with an associated value. If the point
// already exists, we overwrite the existing value
func (kdt *KDTree) Insert(p *point.Point, value string) {

	n := &Node{Point: p, Value: value}
	if kdt.Tree == nil {
		kdt.Tree = n
		kdt.Root = n

		kdt.Dimension = p.Dimension
		kdt.Size = kdt.Size + 1
		return
	}

	current := kdt.Tree // start at root
	dimension := current.Point.Dimension
	for level := 0; true; level++ {
		// get n'th level component % dimension
		//log.Printf("level - %d \n", level)

		if point.Equal(current.Point, p) {
			// no dupes
			return
		}

		l := current.Left
		if l != nil {
			if point.Equal(l.Point, p) {
				return
			}
		}

		r := current.Right
		if r != nil {
			if point.Equal(r.Point, p) {
				return
			}
		}

		component := level % dimension
		//log.Printf("component - %d \n", component)

		if p.Coords[component] < current.Point.Coords[component] {
			// insert into left side
			if current.Left == nil {
				current.Left = n
				log.Printf("current.Left - %#v | %#v \n", current.Left, n)
				kdt.Size = kdt.Size + 1
				n.Level = level
				break
			} else {
				current = current.Left
			}
		} else {
			if current.Right == nil {
				current.Right = n
				log.Printf("current.Right - %#v | %#v \n", current.Right, n)
				kdt.Size = kdt.Size + 1
				n.Level = level
				break
			} else {
				current = current.Right
			}
		}
	}
	//
}

func (kdt *KDTree) Contains(p *point.Point) bool {
	currentNode := kdt.Root

	for {
		if currentNode == nil {
			break
		}

		if point.Equal(p, currentNode.Point) {
			return true
		}

		level := currentNode.Level % currentNode.Point.Dimension
		if currentNode.Point.Coords[level] < p.Coords[level] {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}

	return false
}

func (kdt *KDTree) PrintTree() {
	kdt.PrintAll(kdt.Tree)
}

func (kdt *KDTree) PrintAll(n *Node) {
	if n != nil {
		log.Printf("-> %d] %#v \n", n.Level, n.Point.Coords)
		kdt.PrintAll(n.Left)
		kdt.PrintAll(n.Right)
	}
}

// KNNValue takes a point p and integer k, finds the k points in the KDTree
// nearest to p and returns the most common value associated with the points.
func (kdt *KDTree) KNNValue(p *point.Point, k int) {
	bpq := bpqueue.New(k)
	root := kdt.Tree

	kdt.Search(root, bpq, p)
	log.Printf("bpq - %#v \n", bpq)
}

// Search function for KDTree
//
func (kdt *KDTree) Search(curr *Node, bpq *bpqueue.BPQueue, lookup *point.Point) {
	if curr == nil {
		return
	}

	bpq.Enqueue(curr, point.Distance(lookup, curr.Point))

	level := curr.Level % curr.Point.Dimension
	if lookup.Coords[level] < curr.Point.Coords[level] {
		kdt.Search(curr.Left, bpq, lookup)
		// if hypersphere crosses the plane, check the other branch
		currPriority := bpq.Queue[0].Priority
		if len(bpq.Queue) < bpq.MaxSize || math.Abs(curr.Point.Coords[level]-lookup.Coords[level]) < currPriority {
			kdt.Search(curr.Right, bpq, lookup)
		}
	} else {
		kdt.Search(curr.Right, bpq, lookup)
		currPriority := bpq.Queue[0].Priority
		if len(bpq.Queue) < bpq.MaxSize || math.Abs(curr.Point.Coords[level]-lookup.Coords[level]) < currPriority {
			kdt.Search(curr.Left, bpq, lookup)
		}
	}

}

/*
func (kdt *KDTree) Search(curr *Node, other *Node, lookup *point.Point, bpq *bpqueue.BPQueue, level int) {
	if curr == nil {
		return
	}

	bpq.Enqueue(curr, point.Distance(lookup, curr.Point))
	if lookup.Coords[level] < curr.Point.Coords[level] {
		kdt.Search(curr.Left, curr.Right, lookup, bpq, (level+1)%curr.Point.Dimension)
	} else {
		kdt.Search(curr.Right, curr.Left, lookup, bpq, (level+1)%curr.Point.Dimension)
	}

	currPriority := bpq.Queue[0].Priority
	if len(bpq.Queue) < bpq.MaxSize || math.Abs(curr.Point.Coords[level]-lookup.Coords[level]) < currPriority {
		kdt.Search(other, nil, lookup, bpq, level)
	}
}
*/
