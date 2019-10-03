package ebby

import (
	"github.com/conest/ebby/game"
)

// Ebby : 实例定义
type Ebby struct {
	sm game.SceneMap
	fn game.ExFunctions
	sd interface{}
}

// New :
func New(sm game.SceneMap) *Ebby {
	return &Ebby{
		sm: sm,
		fn: game.ExFunctions{},
	}
}

// SetInitialFunc :
func (e *Ebby) SetInitialFunc(fn func(*game.Game)) {
	e.fn.Ini = fn
}

// SetBeforeExitFunc :
func (e *Ebby) SetBeforeExitFunc(fn func(*game.Game)) {
	e.fn.Exi = fn
}

// SetShareData :
func (e *Ebby) SetShareData(i interface{}) {
	e.sd = i
}

// Run :
func (e *Ebby) Run() {
	if e.fn.Ini == nil {
		e.fn.Ini = defaultInitialFunc
	}
	if e.fn.Exi == nil {
		e.fn.Exi = defaultBeforeExitFunc
	}

	game.Enter(e.sm, &e.fn, e.sd)
}
