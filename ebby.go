package ebby

import (
	"ebby/control"
)

// Ebby : 实例定义
type Ebby struct {
	sm control.SceneMap
	fn control.Functions
	sd interface{}
}

// New :
func New(sm control.SceneMap) *Ebby {
	return &Ebby{
		sm: sm,
		fn: control.Functions{},
	}
}

// SetInitialFunc :
func (e *Ebby) SetInitialFunc(fn func(*control.Control)) {
	e.fn.Ini = fn
}

// SetBeforeExitFunc :
func (e *Ebby) SetBeforeExitFunc(fn func(*control.Control)) {
	e.fn.Bfex = fn
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
	if e.fn.Bfex == nil {
		e.fn.Bfex = defaultBeforeExitFunc
	}

	control.Enter(e.sm, &e.fn, e.sd)
}
