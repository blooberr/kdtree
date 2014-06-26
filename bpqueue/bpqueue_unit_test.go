package bpqueue

import(
  "testing"
)

func TestInsert(t *testing.T) {
  bpq := New(5)
  t.Logf("bpq New - %#v \n", bpq)

  n1 := Node{Priority: 2.4}
  bpq.Enqueue(n1)

  n2 := Node{Priority: 5.3}
  bpq.Enqueue(n2)

  n3 := Node{Priority: 8.3}
  bpq.Enqueue(n3)

  n4 := Node{Priority: 2.1}
  bpq.Enqueue(n4)

  n5 := Node{Priority: 3.7}
  bpq.Enqueue(n5)

  n6 := Node{Priority: 9.9}
  res := bpq.Enqueue(n6)

  if res.Priority != n6.Priority {
    t.Errorf("this priority should have been dequeued \n")
  }

  n7 := Node{Priority: 7.5}
  res = bpq.Enqueue(n7)
  if res.Priority != n3.Priority {
    t.Errorf("this priority should have been 8.3\n")
  }
}

