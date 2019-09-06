package textscen

import (
	"ebby/control/scenario"
	"ebby/control/def"
	"fmt"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

const (
	// rps : 执行部分刷新速率
	rps = 3
)

// Scenario : 返回该实例（一般不需要修改）
func Scenario() *scenario.Scenario {
	i := scenario.Instance(&instance{})
	s := scenario.New(rps, i)
	return s
}

// instance : 数据实例，可以自定义
type instance struct {
	sdata *def.ShareData
	data  *customData
}

// customData : 自定义数据格式
type customData struct {
	num    int
	ticker *time.Ticker
}

// SetSData : 设置共享数据
func (i *instance) SetSData(sdata *def.ShareData) {
	i.sdata = sdata
}

// ResetData : 重置自定义数据
func (i *instance) ResetData() {
	i.data = &customData{}
}

// Initial : 初始化场景
func (i *instance) Initial(w *pixelgl.Window) {
	logger := i.sdata.Tool.DebugLogger
	fmt.Fprintln(logger, "Here is Test 1")

	i.data.ticker = time.NewTicker(5 * time.Second)

	i.data.num = 1
}

func (i *instance) Excuter(dt float64) def.Request {
	debugLogger := i.sdata.Tool.DebugLogger

	r := def.DefaultRequest
	select {
	case <-i.data.ticker.C:
		debugLogger.Clear()
		r.Continue = false
		r.NextScenario = "test2"
		return r
	default:
	}

	i.data.num++
	fmt.Fprintf(debugLogger, "%v\n", i.data.num)

	return r
}

func (i *instance) Drawer(w *pixelgl.Window) {
}
