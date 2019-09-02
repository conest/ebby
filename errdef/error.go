package errdef

type errStruct struct {
	ID  string
	Str string
}

var (
	// FontLoadFile :
	FontLoadFile = errStruct{
		ID:  "FontLoadFile",
		Str: "字体初始化失败",
	}
)
