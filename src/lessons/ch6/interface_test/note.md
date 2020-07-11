### func(xxx type) methodName() (return type){} 这类函数在接口所实现的方法中被调用
type House interface {
	size() float32
}
type Villa struct {
	x, y float32
}
func (villa Villa) size() float32 {
	return float32(villa.x) * float32(villa.y)
}
villa := Villa{x: 11.0, y: 12.0}
t.Log(getSize(villa))