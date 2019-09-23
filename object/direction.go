package object

import (
	"math"
)

// Direct4 : 4方向方位
//   0
// 3 X 1
//   2
type Direct4 int

// Direct4tab : 4方向定位表
var Direct4tab = []Location{
	Location{X: 0, Y: 1},
	Location{X: 1, Y: 0},
	Location{X: 0, Y: -1},
	Location{X: -1, Y: 0},
}

// GenDirect4 : 根据2点生成4方位方向 (lo: 原点，de: 指向点)
func GenDirect4(lo, de Location) Direct4 {
	var s Direct4
	if lo.Eq(de) {
		return s
	}

	disl := de.Sub(lo)
	if Abs(disl.X) > Abs(disl.Y) {
		s = Direct4(2 - Sign(disl.X))
	} else {
		s = Direct4(1 - Sign(disl.Y))
	}
	return s
}

// Location : 返回Location格式的方位
func (d Direct4) Location() Location {
	if !(d >= 0 && int(d) <= len(Direct4tab)-1) {
		return Location{}
	}
	return Direct4tab[d]
}

// NextLocation : 根据输入地点返回方位所指向的接邻点
func (d Direct4) NextLocation(lo Location) Location {
	return lo.Add(d.Location())
}

// Direct8 : 8方向方位
// 7 0 1
// 6 X 2
// 5 4 3
type Direct8 int

// Direct8tab : 8方向定位表
var Direct8tab = []Location{
	Location{X: 0, Y: 1},
	Location{X: 1, Y: 1},
	Location{X: 1, Y: 0},
	Location{X: 1, Y: -1},
	Location{X: 0, Y: -1},
	Location{X: -1, Y: -1},
	Location{X: -1, Y: 0},
	Location{X: -1, Y: 1},
}

// GenDirect8 : 根据2点生成8方位方向 (lo: 原点，de: 指向点)
func GenDirect8(lo, de Location) Direct8 {

	if lo.Eq(de) {
		return 0
	}

	disl := de.Sub(lo)
	disl.X = Sign(disl.X)
	disl.Y = Sign(disl.Y)

	for i, v := range Direct8tab {
		if disl == v {
			return Direct8(i)
		}
	}
	return 0
}

// GenDirect8s : 根据2点生成8方位方向，斜度版本，性能相比GenDirect8会慢 (lo: 原点，de: 指向点)
func GenDirect8s(lo, de Location) Direct8 {

	var s Direct8
	if lo.Eq(de) {
		return s
	}

	disl := de.Sub(lo)
	// 避开除数为0
	if disl.Y == 0 || disl.X == 0 {
		return GenDirect8(lo, de)
	}

	slope := float64(disl.X) / float64(disl.Y)
	radian := math.Atan(slope)
	if disl.Y < 0 {
		radian += math.Pi
	}
	if disl.X < 0 && disl.Y > 0 {
		radian += math.Pi * 2
	}
	cr := math.Mod(radian+math.Pi/8, math.Pi*2)

	return Direct8(cr / (math.Pi / 4))
}

// Location : 返回Location格式的方位
func (d Direct8) Location() Location {
	if !(d >= 0 && int(d) <= len(Direct8tab)-1) {
		return Location{}
	}
	return Direct8tab[d]
}

// NextLocation : 根据输入地点返回方位所指向的接邻点
func (d Direct8) NextLocation(lo Location) Location {
	return lo.Add(d.Location())
}
