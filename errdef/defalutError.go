package errdef

type errStruct struct {
	ID  string
	Str string
}

var (
	// CreateWindow :
	CreateWindow = errStruct{
		ID:  "CreateWindow",
		Str: "Window创建失败",
	}
	// FontLoadFile :
	FontLoadFile = errStruct{
		ID:  "FontLoadFile",
		Str: "字体初始化失败",
	}

	// NoScenario :
	NoScenario = errStruct{
		ID:  "NoScenario",
		Str: "无法取得对应ID的Scenario",
	}
)
