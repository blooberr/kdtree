package point

import(
  "testing"
)

func TestPointDistance(t *testing.T) {
  p1 := New(3,4,5)
  t.Logf("p1 %#v \n", p1)

  if p1.Size() != 3 {
    t.Errorf("Dimension should be 3 for (3,4,5). \n")
  }

  p2 := New(8,3,3.2)

  eDistance := Distance(p1, p2)
  t.Logf("euclidean distance is: %#v \n", eDistance)
}

func TestPointEquality(t *testing.T) {
  p1 := New(5,3,1)
  p2 := New(5,3,1)

  if Equal(p1, p2) != true {
    t.Errorf("Should be equal.\n")
  }

  p3 := New(10,2,9)
  if Equal(p1, p3) != false {
    t.Errorf("both points should not equal.\n")
  }
}
