package point

import(
  "log"
  "math"
)

type Point struct {
  Coords []float64
}

func New(dimension ...float64) *Point {
  log.Printf("dimension - %#v \n", dimension)
  return &Point{Coords: dimension}
}

func (p *Point) Size() int {
  return len(p.Coords)
}

func Distance(one *Point, two *Point) (distance float64) {
  // calculate euclidean distance
  result := 0.0

  for i := 0; i < one.Size(); i++ {
    result += (one.Coords[i] - two.Coords[i]) * (one.Coords[i] - two.Coords[i])
  }

  return math.Sqrt(result)
}

func Equal(one *Point, two *Point) bool {
  for i := 0; i < one.Size(); i++ {
    if one.Coords[i] != two.Coords[i] {
      return false
    }
  }

  return true
}

