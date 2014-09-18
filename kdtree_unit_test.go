package kdtree

import (
	"github.com/blooberr/kdtree/point"
	"testing"
)

func TestInsert(t *testing.T) {

	root := point.New(3, 1, 4)

	kdt := New()
	kdt.Insert(root, "this is the root")

	t.Logf("%#v \n", kdt.Tree)

	kdt.Insert(point.New(2, 3, 7), "second one")
	kdt.PrintTree()

	kdt.Insert(point.New(4, 3, 4), "third point")
	kdt.PrintTree()

	kdt.Insert(point.New(2, 1, 3), "fourth point")
	kdt.PrintTree()

	kdt.Insert(point.New(6, 1, 4), "")
	kdt.Insert(point.New(2, 4, 5), "fifth point")
	kdt.Insert(point.New(1, 4, 4), "----")
	kdt.PrintTree()

	kdt.KNNValue(point.New(1, 1, 1), 5)
}

func TestBasicKDTree(t *testing.T) {
	kdt := New()

	if kdt.Size != 0 {
		t.Errorf("initial KDTree should be 0!\n")
	}

	kdt.Insert(point.New(1, 0, 0), "Y")
	kdt.Insert(point.New(0, 1, 0), "Y")
	kdt.Insert(point.New(0, 0, 1), "N")

	if kdt.Size != 3 {
		t.Errorf("we just inserted 3 points \n")
	}

	if kdt.Dimension != 3 {
		t.Errorf("should be dimension 3\n")
	}
}

func TestModerateKDTree(t *testing.T) {
	kdt := New()
	kdt.Insert(point.New(0, 0, 0, 0), "Y")
	kdt.Insert(point.New(0, 0, 0, 1), "Y")
	kdt.Insert(point.New(0, 0, 1, 0), "Y")
	kdt.Insert(point.New(0, 0, 1, 1), "Y")
	kdt.Insert(point.New(0, 1, 0, 0), "A")
	kdt.Insert(point.New(0, 1, 0, 1), "B")
	kdt.Insert(point.New(0, 1, 1, 0), "C")
	kdt.Insert(point.New(0, 1, 1, 1), "Y")
	kdt.Insert(point.New(1, 0, 0, 0), "Y")

	if kdt.Size != 9 {
		t.Errorf("we just inserted 9 points \n")
	}

	if kdt.Dimension != 4 {
		t.Errorf("should be dimension 4\n")
	}

	exist := kdt.Contains(point.New(1.0, 1.0, 1.0, 0.5))
	if exist != false {
		t.Errorf("point should not exist \n")
	}

	exist = kdt.Contains(point.New(0.0, 0.0, 0.0, -0.5))
	if exist != false {
		t.Errorf("point should not exist \n")
	}

}

func TestHarderKDTree(t *testing.T) {
	kdt := New()
	kdt.Insert(point.New(0, 0, 0, 0), "Y")
	kdt.Insert(point.New(0, 1, 0, 1), "Y")
	kdt.Insert(point.New(0, 0, 0, 0), "Y") // dupe
	kdt.Insert(point.New(0, 1, 0, 1), "Y") // dupe
	kdt.Insert(point.New(0, 1, 1, 0), "C")
	kdt.Insert(point.New(1, 0, 1, 0), "C")

  if kdt.Dimension != 4 {
    t.Errorf("should be dimension 4 \n")
  }

  if kdt.Size != 4 {
    t.Errorf("we should have only inserted 4 (no dupes). We have %d inserted \n", kdt.Size)
    kdt.PrintTree()
  }
}
