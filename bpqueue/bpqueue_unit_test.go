package bpqueue

import(
  "testing"
)

type DummyStruct struct {
  Hello string
}

func TestInsert(t *testing.T) {
  bpq := New(5)
  t.Logf("bpq New - %#v \n", bpq)

  bpq.Enqueue(DummyStruct{Hello: "world"}, 2.4)
  bpq.Enqueue(DummyStruct{Hello: "bear"}, 5.3)
  bpq.Enqueue(DummyStruct{Hello: "tiger"}, 2.1)
  bpq.Enqueue(DummyStruct{Hello: "chain"}, 7.3)
}

