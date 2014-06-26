package bpqueue

import (
	"log"
)

// implementation of bounded priority-queue. Whenever a new element is added to
// the queue, if the queue is at capactiy - then the element with the highest
// priority value is ejected from the queue.

type BPQueue struct {
	MaxSize int
	Queue   []*Node
}

type Node struct {
	Priority float64
	Inner    interface{}
}

func New(size int) *BPQueue {
	bpq := BPQueue{MaxSize: size}
	//bpq.Queue = make(*Node, size)
	return &bpq
}

func GetParentIndex(index int) int {
	parentIndex := (index - 1) / 2

	if parentIndex < 0 {
		return 0
	} else {
		return parentIndex
	}
}

func GetLeftChildren(index int) int {
	return 2*index + 1
}

func GetRightChildren(index int) int {
	return 2*index + 2
}

func (bpq *BPQueue) Enqueue(node Node) (*Node) {
	log.Printf("enqueue - %#v \n", node)

  bpq.Insert(node)

	if len(bpq.Queue) > bpq.MaxSize {
		// time to dequeue
		log.Printf("time to eject - reached max %d \n", bpq.MaxSize)
    return bpq.Delete()
	} else {
    return nil
  }
}

func (bpq *BPQueue) Insert(node Node) {
	// add to bottom of heap.
	bpq.Queue = append(bpq.Queue, &node)
	log.Printf("insert - %#v \n", bpq.Queue)

	currentNodeIndex := len(bpq.Queue) - 1

	for {
		if !(currentNodeIndex > 0) {
			break
		}

		parentIndex := GetParentIndex(currentNodeIndex)
		log.Printf("c: %d and parent: %d \n", currentNodeIndex, parentIndex)

		parentNode := bpq.Queue[parentIndex]
		if parentNode.Priority < node.Priority {
			bpq.Swap(parentIndex, currentNodeIndex)
			currentNodeIndex = parentIndex
		} else {
      break
    }
	}

  for _, node := range bpq.Queue {
    log.Printf("[%f] ", node.Priority)
  }
  log.Printf("********* \n")
}

func (bpq *BPQueue) Delete() *Node {
  head := bpq.Queue[0]

  log.Printf("delete() \n")
  for _, node := range bpq.Queue {
    log.Printf("[%f] ", node.Priority)
  }

  // replace root with last element
  bpq.Queue[0] = bpq.Queue[len(bpq.Queue) - 1]
  bpq.Queue = bpq.Queue[0:len(bpq.Queue)-1]

  log.Printf("delete() \n")
  for _, node := range bpq.Queue {
    log.Printf("[%f] ", node.Priority)
  }

  // check if the root is correct
  index := 0
  for {
    lIndex := GetLeftChildren(index)
    rIndex := GetRightChildren(index)
    largestIndex := index

    // check if there's children. if none, quit
    if lIndex < len(bpq.Queue) {
      if bpq.Queue[lIndex].Priority > bpq.Queue[largestIndex].Priority {
        largestIndex = lIndex
      }
    }

    if rIndex < len(bpq.Queue) {
      if bpq.Queue[rIndex].Priority > bpq.Queue[largestIndex].Priority {
        largestIndex = rIndex
      }
    }

    if largestIndex != index {
      bpq.Swap(index, largestIndex)
      index = largestIndex
    } else {
      break
    }
  }

  log.Printf("after max-heapify \n")
  for _, node := range bpq.Queue {
    log.Printf("[%f] ", node.Priority)
  }

  return head
}

func (bpq *BPQueue) Swap(one int, two int) {
  log.Printf("swapping position %d with %d \n", one, two)
  log.Printf("node at one %#v \n", bpq.Queue[one])
  log.Printf("node at two %#v \n", bpq.Queue[two])

  bpq.Queue[one], bpq.Queue[two] = bpq.Queue[two], bpq.Queue[one]
}

