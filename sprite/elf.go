package sprite

import (
	"github.com/faiface/pixel"
)

type aniFrame struct {
	frame pixel.Rect
	last  int64
}

// Elf : 强化型 Sprite，支持动画等
type Elf struct {
	s       *pixel.Sprite
	picture *pixel.Picture
	fTable  [][]aniFrame
	fID     int
	animate bool
}

// NewElf creates a Elf from the supplied frame of a Picture and more.
func NewElf(p *pixel.Picture) *Elf {
	e := Elf{
		s:       pixel.NewSprite(*p, pixel.Rect{}),
		picture: p,
	}

	return &e
}
