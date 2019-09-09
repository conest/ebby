# Ebby

一个使用Go编写的2D游戏引擎，基于faiface/pixel开发。提供了基础的框架引擎，以场景为单位的分离和调度，并提供了实用的工具集。

项目还在处于猛烈开发中，内容随时可能改动

```
go get github.com/conest/ebby
```

## 范例

### 入口

``` go
// 配置入口函数
func run() {
    // 加载你的场景列表
	sm := control.ScenarioMap{
		"dungeon": dungeon.Scenario(),
	}
    // 新建ebby实例
	eb := ebby.New(sm)
    // 设置全局共享数据格式
	eb.SetShareData(&udef.ShareData{})
	eb.Run()
}

func main() {
	pixelgl.Run(run)
}
```

### Hello World 场景配置
```go
const (
	// rps : 执行部分刷新速率
	rps = 60
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
	Hello string
}

// SetSData : 设置共享数据
func (i *instance) SetSData(sdata *def.ShareData) {
	i.sdata = sdata
}

// ResetData : 重置自定义数据
func (i *instance) ResetData() {
	i.data = &customData{}
}

// 初始化场景
func (i *instance) Initial(w *pixelgl.Window) {
    i.data.Hello = "Hello World!"
    // 使用 Debug 工具中的 Logger 在屏幕上显示字符串
    fmt.Fprintf(i.sdata.Tool.DebugLogger, "%v", i.data.Hello)
}

// 执行计算阶段调用函数
func (i *instance) Excuter(dt float64) def.Request {
	return def.DefaultRequest
}

// 绘图阶段调用函数
func (i *instance) Drawer(w *pixelgl.Window, dt float64) {

}

// 输出读取调用函数
func (i *instance) InputHandle(w *pixelgl.Window, dt float64) {
	
}
```

## TODO
- [x] 场景(Scenario)调度
- [x] 基础Debug工具
- [x] Logger(使用logrus)
- [x] 基础字体加载
- [x] Elf - 支持动画的高级精灵对象
- [ ] 精灵图处理工具集
- [ ] 高级实用工具集
- [ ] 完善中文字体加载性能问题