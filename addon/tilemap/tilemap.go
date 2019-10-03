package tilemap

import (
	"github.com/conest/github.com/faiface/pixel"
)

// TileMap :
type TileMap struct {
	x         int
	y         int
	tilePixel int
	tileTab   [][]int
	frameList []pixel.Rect
	pic       *pixel.Picture
	batch     *pixel.Batch
}

// NewTileMap : 创建新的 TileMap
func NewTileMap(tilePixel, x, y int, pic *pixel.Picture, fl []pixel.Rect) *TileMap {

	gt := make([][]int, y)
	for i := range gt {
		gt[i] = make([]int, x)
	}

	return &TileMap{
		x:         x,
		y:         y,
		tilePixel: tilePixel,
		tileTab:   gt,
		frameList: fl,
		pic:       pic,
		batch:     pixel.NewBatch(&pixel.TrianglesData{}, *pic),
	}
}

// SetPic : 设置图像
func (g *TileMap) SetPic(p *pixel.Picture) {
	g.pic = p
}

// SetFrameList : 设置Frame列表
func (g *TileMap) SetFrameList(fl []pixel.Rect) {
	g.frameList = fl
}

// SetTileTabValue : 设置网格地图值
func (g *TileMap) SetTileTabValue(x, y, value int) {
	g.tileTab[y][x] = value
}

// TileTabValue : 获取网格地图值
func (g *TileMap) TileTabValue(x, y int) int {
	return g.tileTab[y][x]
}

// Batch : 获取Batch
func (g *TileMap) Batch() *pixel.Batch {
	return g.batch
}

// Draw : Draw Batch
func (g *TileMap) Draw(target pixel.Target) {
	g.batch.Draw(target)
}

// DrawSprite : 单独绘制Sprite到Batch上
func (g *TileMap) DrawSprite(x, y, v int) {
	sprite := pixel.NewSprite(*g.pic, g.frameList[v])
	l := NewLocation(x, y)
	vec := l.ToVec(true, g.tilePixel)
	sprite.Draw(g.batch, pixel.IM.Moved(vec))
}

// BatchGen : 根据 TileTab 生成 Batch
func (g *TileMap) BatchGen() {
	g.batch.Clear()
	sprite := pixel.NewSprite(nil, pixel.Rect{})
	for y, row := range g.tileTab {
		for x, v := range row {
			sprite.Set(*g.pic, g.frameList[v])
			l := Location{X: x, Y: y}
			vec := l.ToVec(true, g.tilePixel)
			sprite.Draw(g.batch, pixel.IM.Moved(vec))
		}
	}
}
