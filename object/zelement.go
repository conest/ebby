package object

import (
	"sort"

	"github.com/faiface/pixel"
)

// Zelement : 用于根据Z值来排序与批量Draw的接口
type Zelement interface {
	Name() string
	Show() bool
	Z() int
	Draw(pixel.Target)
	UDraw(float64, pixel.Target)
}

// Zcluster : Element 聚合，用于根据Z值来排序与批量Draw
type Zcluster []Zelement

// Push : Push 新的 Element
func (zc *Zcluster) Push(z Zelement) int {
	*zc = append(*zc, z)
	return len(*zc)
}

// Delete : 删除对应name的 Zelement
func (zc Zcluster) Delete(name string) {
	for i, ze := range zc {
		if ze.Name() == name {
			copy(zc[i:], zc[i+1:])
			zc[len(zc)-1] = nil
			zc = zc[:len(zc)-1]
		}
	}
}

func (zc Zcluster) Len() int {
	return len(zc)
}

func (zc Zcluster) Swap(i, j int) {
	zc[i], zc[j] = zc[j], zc[i]
}

func (zc Zcluster) Less(i, j int) bool {
	return zc[i].Z() < zc[j].Z()
}

// Sort : 根据Z值来排序
func (zc Zcluster) Sort() {
	sort.Sort(zc)
}

// Draw : 批量Draw
func (zc Zcluster) Draw(target pixel.Target) {
	for _, e := range zc {
		if e.Show() {
			e.Draw(target)
		}
	}
}

// UDraw : 批量UDraw
func (zc Zcluster) UDraw(dt float64, target pixel.Target) {
	for _, e := range zc {
		if e.Show() {
			e.UDraw(dt, target)
		}
	}
}

// SUDraw : sort后批量UDraw
func (zc Zcluster) SUDraw(dt float64, target pixel.Target) {
	zc.Sort()
	zc.UDraw(dt, target)
}
