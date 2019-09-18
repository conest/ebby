package object

import (
	"fmt"

	"github.com/faiface/pixel"
)

// Easel : Picture、Batch 综合管理与控制实例
type Easel struct {
	pictures map[string]*pixel.Picture
	batchs   map[string]*pixel.Batch
	frames   map[string][]pixel.Rect
}

// NewEasel : 新Easel
func NewEasel() Easel {
	return Easel{
		pictures: make(map[string]*pixel.Picture, 0),
		batchs:   make(map[string]*pixel.Batch, 0),
		frames:   make(map[string][]pixel.Rect, 0),
	}
}

// SetPicture : 设置图片
func (e *Easel) SetPicture(name string, pic *pixel.Picture) {
	e.pictures[name] = pic
}

// Picture : 获取图片
func (e *Easel) Picture(name string) *pixel.Picture {
	pic, exist := e.pictures[name]
	if !exist {
		panic(fmt.Errorf("[Easel] Picture get error. No picture named: %s", name))
	}
	return pic
}

// SetBatch : 设置Batch
func (e *Easel) SetBatch(name string, picName string) {
	pic, exist := e.pictures[picName]
	if !exist {
		panic(fmt.Errorf("[Easel] Set Batch error. No picture named: %s", picName))
	}
	e.batchs[name] = pixel.NewBatch(&pixel.TrianglesData{}, *pic)
}

// Batch : 获取Batch
func (e *Easel) Batch(name string) *pixel.Batch {
	batch, exist := e.batchs[name]
	if !exist {
		panic(fmt.Errorf("[Easel] Batch get error. No Batch named: %s", name))
	}
	return batch
}

// SetFrame : 设置Frame
func (e *Easel) SetFrame(name string, frame []pixel.Rect) {
	e.frames[name] = frame
}

// Frame : 获取Frame
func (e *Easel) Frame(name string) []pixel.Rect {
	frame, exist := e.frames[name]
	if !exist {
		panic(fmt.Errorf("[Easel] Frame get error. No Frame named: %s", name))
	}
	return frame
}
