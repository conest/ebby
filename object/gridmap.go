package object

import "github.com/faiface/pixel"

// GridMap :
type GridMap struct {
	x         int
	y         int
	gridPixel int
	gridTab   [][]int
	frameList []pixel.Rect
	pic       *pixel.Picture
	batch     *pixel.Batch
}

// NewGridMap : 创建新的 GridMap
func NewGridMap(gridPixel, x, y int) *GridMap {

	gt := make([][]int, y)
	for i := range gt {
		gt[i] = make([]int, x)
	}

	return &GridMap{
		x:         x,
		y:         y,
		gridPixel: gridPixel,
		gridTab:   gt,
	}
}

// SetPic : 设置图像
func (g *GridMap) SetPic(p *pixel.Picture) {
	g.pic = p
}

// SetFrameList : 设置Frame列表
func (g *GridMap) SetFrameList(fl []pixel.Rect) {
	g.frameList = fl
}

// SetGridTabValue : 设置网格地图值
func (g *GridMap) SetGridTabValue(x, y, value int) {
	g.gridTab[y][x] = value
}

// GridTabValue : 获取网格地图值
func (g *GridMap) GridTabValue(x, y int) int {
	return g.gridTab[y][x]
}

// TODO:
// BatchGen : 根据 GridTab 生成 Batch
// func (g *GridMap) BatchGen() {
// 	batch := pixel.NewBatch(&pixel.TrianglesData{}, *g.pic)
// 	sprite := pixel.NewSprite(nil, pixel.Rect{})
// 	for y, l := range m.frameNum {
// 		for x, v := range l {
// 			sprite.Set(*ss, frames[v])

// 			vec := pixel.Vec{
// 				X: float64(x*pSize + 8),
// 				Y: float64(y*pSize + 8),
// 			}
// 			sprite.Draw(batch, pixel.IM.Moved(vec))
// 		}
// 	}

// 	return batch
// }
