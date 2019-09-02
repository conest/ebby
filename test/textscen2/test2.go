package textscen2

import (
	"campaign/control/datastruct"
	"campaign/control/order"
	"campaign/scenario"
	"fmt"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

const (
	// rps : 执行部分刷新速率
	rps = 1
)

// Scenario : 返回该实例（一般不需要修改）
func Scenario() *scenario.Scenario {
	i := scenario.Instance(&instance{})
	s := scenario.New(rps, i)
	return s
}

// instance : 数据实例，可以自定义
type instance struct {
	sdata *datastruct.ShareData
}

// SetSData : 设置共享数据指针
func (i *instance) SetSData(sdata *datastruct.ShareData) {
	i.sdata = sdata
}

func (i *instance) Initial(w *pixelgl.Window) {
	logger := i.sdata.Tool.DebugLogger
	fmt.Fprintln(logger, "Now, print Latin only")
}

func (i *instance) Excuter(dt float64) *order.Request {
	logger := i.sdata.Tool.DebugLogger
	fmt.Fprintf(logger, "%v\n", time.Now().String())
	r := &order.Request{
		Continue: true,
	}
	i.sdata.Tool.Logger.Info("hehe")
	return r
}

func (i *instance) Drawer(w *pixelgl.Window) {
}
