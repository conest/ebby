package object

import (
	"github.com/faiface/pixel"
)

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
func NewGridMap(gridPixel, x, y int, pic *pixel.Picture, fl []pixel.Rect) *GridMap {

	gt := make([][]int, y)
	for i := range gt {
		gt[i] = make([]int, x)
	}

	return &GridMap{
		x:         x,
		y:         y,
		gridPixel: gridPixel,
		gridTab:   gt,
		frameList: fl,
		pic:       pic,
		batch:     pixel.NewBatch(&pixel.TrianglesData{}, *pic),
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

// Batch : 获取Batch
func (g *GridMap) Batch() *pixel.Batch {
	return g.batch
}

// BatchGen : 根据 GridTab 生成 Batch
func (g *GridMap) BatchGen() {
	g.batch.Clear()
	sprite := pixel.NewSprite(nil, pixel.Rect{})
	for y, row := range g.gridTab {
		for x, v := range row {
			sprite.Set(*g.pic, g.frameList[v])
			l := Location{X: x, Y: y}
			vec := l.ToVec(true, g.gridPixel)
			sprite.Draw(g.batch, pixel.IM.Moved(vec))
		}
	}
}
