package interface_pkg

import "testing"

type House interface {
	size() float32
}

type Pool struct {
	depth float32
	wide  float32
	lens  float32
}
type Villa struct {
	pool Pool
	x, y float32
}

func (villa Villa) size() float32 {
	return float32(villa.x) * float32(villa.y)
}

type Mansion struct {
	pool Pool
	x, y float32
}

func (mansion Mansion) size() float32 {
	return float32(mansion.x) + float32(mansion.y)
}
func getSize(house House) float32 {
	return house.size()
}

func TestInterface(t *testing.T) {
	pool := Pool{depth: 3.2, wide: 10.0, lens: 20.0}
	villa := Villa{pool: pool, x: 11.0, y: 12.0}
	t.Log(getSize(villa))

	mansion := Mansion{pool: pool, x: 11.0, y: 12.0}
	t.Log(getSize(mansion))
}
