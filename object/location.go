package object

import "github.com/faiface/pixel"

// Location : 位置
type Location struct {
	X, Y int
}

// NewLocation : 输入 (x, y) 创建Location
func NewLocation(x, y int) Location {
	return Location{
		X: x,
		Y: y,
	}
}

// NewLocationFromVec : 输入 Vec 创建Location
func NewLocationFromVec(vec pixel.Vec, gridPixel int) Location {
	return Location{
		X: int(vec.X) / gridPixel,
		Y: int(vec.Y) / gridPixel,
	}
}

// Vec : 直接转换为Vec
func (l *Location) Vec() pixel.Vec {
	return pixel.Vec{
		X: float64(l.X),
		Y: float64(l.Y),
	}
}

// ToVec : 根据gridPixel转换为 Vec (central：返回的为中心点，否则为(0,0)点)
func (l *Location) ToVec(central bool, gridPixel int) pixel.Vec {
	var c float64
	if central {
		c = float64(gridPixel / 2)
	}
	return pixel.Vec{
		X: float64(l.X*gridPixel) + c,
		Y: float64(l.Y*gridPixel) + c,
	}
}

// ToVecOffset : 根据偏移距离转换为 Vec (offset: 偏移距离)
func (l *Location) ToVecOffset(offset float64, gridPixel int) pixel.Vec {
	return pixel.Vec{
		X: float64(l.X*gridPixel) + offset,
		Y: float64(l.Y*gridPixel) + offset,
	}
}

// SetFromVec : 从Vec设置Location
func (l *Location) SetFromVec(vec pixel.Vec, gridPixel int) {
	l.X = int(vec.X) / gridPixel
	l.Y = int(vec.Y) / gridPixel
}

// Eq : 比较是否相等
func (l *Location) Eq(cl Location) bool {
	if l.X == cl.X && l.Y == cl.Y {
		return true
	}
	return false
}

// Add : 相加，返回结果，不改变原值
func (l *Location) Add(cl Location) Location {
	return Location{
		X: l.X + cl.X,
		Y: l.Y + cl.Y,
	}
}

// Sub : 相减，返回结果，不改变原值
func (l *Location) Sub(cl Location) Location {
	return Location{
		X: l.X - cl.X,
		Y: l.Y - cl.Y,
	}
}

// Moved : 根据输入xy改变量，返回新 Location
func (l *Location) Moved(x, y int) Location {
	return Location{
		X: l.X + x,
		Y: l.Y + y,
	}
}

// Distance : 返回两点之间的方格距离
func (l *Location) Distance(cl Location) int {
	x := l.X - cl.X
	y := l.Y - cl.Y
	return Abs(x) + Abs(y)
}

// ToSlice : 返回一个以当前点为0元素的slice
func (l *Location) ToSlice() []Location {
	return []Location{*l}
}
