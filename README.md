# Ebby

一个使用Go编写的2D游戏引擎，基于faiface/pixel开发。提供了基础的框架引擎，以场景为单位的分离和调度，并提供了实用的工具集。

项目还在处于猛烈开发中，内容随时可能改动

```
go get github.com/conest/ebby
```

你还需要安装 pixel

```
go get github.com/faiface/pixel
```

## 范例

### 入口

``` go
// 配置入口函数
func run() {
	// 加载你的场景列表
	sm := control.ScenarioMap{
		"hello": hello.Scenario(),
	}
	// 新建ebby实例并运行
	ebby.New(sm).Run()
}

func main() {
	// 使用pixelgl的入口函数
	pixelgl.Run(run)
}
```

### 配置文件

文件 config.yaml 放在项目根目录下

```yaml
mode: debug

screen:
  rX: 1024
  rY: 800
  title: Hello World
  VSync:  False

scenario:
  # 最高循环速率限制，避免一些情况(关闭垂直同步；窗口最小化等)高速loop导致消耗大量cpu
  # 该功能可能会使垂直同步失效
  # 单位: frames / second
  maxRate: 60
  # 入口场景名称
  entry: hello

logger:
  level: debug
  logToFile: False
  filePath: log/1.log
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
func (i *instance) Drawer(w *pixelgl.Window, dt float64) {}

// 输出读取调用函数
func (i *instance) InputHandle(w *pixelgl.Window, dt float64) {}
```

## TODO
- [x] 场景(Scenario)调度
- [x] 基础Debug工具
- [x] Logger(使用logrus)
- [x] 基础字体加载
- [x] 支持动画的高级sprite对象
- [-] 按钮对象以及监听工具
- [x] 视角工具
- [x] GLSL控制支持
- [-] sprite图片处理工具集
- [ ] 高级实用工具集
- [ ] 完善中文字体加载性能问题