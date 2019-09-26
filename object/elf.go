package object

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
)

// AniFrame : 动画帧
type AniFrame struct {
	Frame pixel.Rect
	Last  float64
}

// Animate : 动画结构
type Animate struct {
	frames  []AniFrame
	sumLast float64
	overDt  float64 // 余出delta time
	Repeat  int     // 重复播放次数，0 为无限循环
	P       int     // 当前播放的动画帧号
	Playing bool    // 是否播放动画
}

// Add : 根据frame和last创建新动画帧
func (a *Animate) Add(frame pixel.Rect, last float64) {
	af := AniFrame{
		Frame: frame,
		Last:  last,
	}
	a.frames = append(a.frames, af)
	a.sumLast += last
}

// Push : 增加新动画帧
func (a *Animate) Push(newFrame AniFrame) {
	a.frames = append(a.frames, newFrame)
	a.sumLast += newFrame.Last
}

// PushList : 列表方式增加新动画帧
func (a *Animate) PushList(newFrames []AniFrame) {
	for _, newFrame := range newFrames {
		a.frames = append(a.frames, newFrame)
		a.sumLast += newFrame.Last
	}
}

// Get : 获取当前动画帧
func (a *Animate) Get(dt float64) pixel.Rect {
	if !a.Playing {
		return a.frames[a.P].Frame
	}
	plusDt := math.Mod(dt, a.sumLast) + a.overDt
	for {
		if plusDt > a.frames[a.P].Last {
			plusDt -= a.frames[a.P].Last
			a.nextFrame()
		} else {
			a.overDt = plusDt
			return a.frames[a.P].Frame
		}
	}
}

// Play : 播放动画
func (a *Animate) Play(repeat int) {
	a.Repeat = repeat
	a.Playing = true
	a.P = 0
	a.overDt = 0
}

// Stop : 停止播放动画
func (a *Animate) Stop() {
	a.Repeat = 0
	a.Playing = false
}

// nextFrame : 动画帧号前进，返回前进后的帧号
func (a *Animate) nextFrame() int {
	if a.P >= len(a.frames)-1 {
		// 判断无限循环
		if a.Repeat == 0 {
			a.P = 0
			return a.P
		}
		a.Repeat--
		// 判断停止
		if a.Repeat == 0 {
			a.Playing = false
		} else {
			a.P = 0
		}
	} else {
		a.P++
	}
	return a.P
}

// Elf : 强化型 Sprite，支持动画等
type Elf struct {
	s        *pixel.Sprite
	picture  *pixel.Picture
	sTable   []pixel.Rect
	sID      int
	Animate  bool
	aniTable []Animate
	aniID    int
	z        int
	show     bool
}

// NewElf creates a Elf from the supplied frame of a Picture and more.
func NewElf(p *pixel.Picture) *Elf {
	e := Elf{
		s:       pixel.NewSprite(*p, pixel.Rect{}),
		picture: p,
		show:    true,
	}
	return &e
}

// Z : 返回Z值
func (e *Elf) Z() int {
	return e.z
}

// SetZ : 设置Z
func (e *Elf) SetZ(z int) {
	e.z = z
}

// Show : 返回Show值
func (e *Elf) Show() bool {
	return e.show
}

// SetShow : 设置Show
func (e *Elf) SetShow(s bool) {
	e.show = s
}

// AddAnimate : 添加新动画，返回动画序列编号
func (e *Elf) AddAnimate(animate Animate) int {
	e.aniTable = append(e.aniTable, animate)
	return len(e.aniTable)
}

// SetAnimate : 配置动画播放
func (e *Elf) SetAnimate(playID int, repeat int) {
	if playID >= len(e.aniTable) {
		playID = len(e.aniTable) - 1
	}
	e.Animate = true
	e.aniID = playID
	e.aniTable[playID].Play(repeat)
}

// AddStatic : 添加新静态图，返回静态图序列编号
func (e *Elf) AddStatic(staticPic pixel.Rect) int {
	e.sTable = append(e.sTable, staticPic)
	return len(e.sTable)
}

// SetStatic : 配置静态图
func (e *Elf) SetStatic(id int) {
	if id >= len(e.sTable) {
		id = len(e.sTable) - 1
	}
	e.Animate = false
	e.sID = id
}

// Update : 更新 sprite
func (e *Elf) Update(dt float64) {
	f := pixel.Rect{}
	if !e.Animate {
		f = e.sTable[e.sID]
	} else {
		f = e.aniTable[e.aniID].Get(dt)
	}
	e.s.Set(*e.picture, f)
}

// Draw : 绘制当前 sprite
func (e *Elf) Draw(target pixel.Target, m pixel.Matrix) {
	e.s.Draw(target, m)
}

// DrawColorMask : 绘制当前 sprite，并应用色彩蒙版
func (e *Elf) DrawColorMask(target pixel.Target, m pixel.Matrix, mask color.Color) {
	e.s.DrawColorMask(target, m, mask)
}

// UDraw : 更新并绘制当前 sprite
func (e *Elf) UDraw(dt float64, target pixel.Target, m pixel.Matrix) {
	e.Update(dt)
	e.s.Draw(target, m)
}
