package ebby

import (
	"github.com/conest/ebby/game"
)

// Ebby : Ebby结构定义
type Ebby struct {
	sceneMap   game.SceneMap
	fn         game.ExFunctions
	publicData interface{}
}

// New : 创建新的Ebby实例
func New(sceneMap game.SceneMap) *Ebby {
	return &Ebby{
		sceneMap: sceneMap,
		fn:       game.ExFunctions{},
	}
}

// SetInitialFunc : 设定Game初始化函数，将在加载整个游戏时触发
func (e *Ebby) SetInitialFunc(fn func(*game.Game)) {
	e.fn.Ini = fn
}

// SetBeforeExitFunc : 设定Game结束函数，将在游戏退出前
func (e *Ebby) SetBeforeExitFunc(fn func(*game.Game)) {
	e.fn.Exi = fn
}

// SetPublicData : 设定公共数据，为Game的全局数据
func (e *Ebby) SetPublicData(i interface{}) {
	e.publicData = i
}

// Run :
func (e *Ebby) Run() {
	if e.fn.Ini == nil {
		e.fn.Ini = defaultInitialFunc
	}
	if e.fn.Exi == nil {
		e.fn.Exi = defaultBeforeExitFunc
	}

	game.Enter(e.sceneMap, &e.fn, e.publicData)
}
