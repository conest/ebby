package errdef

import (
	"ebby/common/logger"
	"fmt"
)

type errStruct struct {
	ID  string
	Str string
}

// CheckErr : 检查错误并输出日志
func CheckErr(err error, where string, es errStruct) {
	if err != nil {
		log := logger.New()
		log.Error(fmt.Sprintf("[%s] <%s> %v", where, es.Str, err))
		panic(err)
	}
}

var (
	// CreateWindow :
	CreateWindow = errStruct{
		ID:  "CreateWindow",
		Str: "Window创建失败",
	}

	// ViperErr :
	ViperErr = errStruct{
		ID:  "ViperErr",
		Str: "Viper设置错误",
	}

	// FontLoadFile :
	FontLoadFile = errStruct{
		ID:  "FontLoadFile",
		Str: "字体初始化失败",
	}

	// PictureLoadFile :
	PictureLoadFile = errStruct{
		ID:  "PictureLoadFile",
		Str: "图片文件读取失败",
	}

	// NoScenario :
	NoScenario = errStruct{
		ID:  "NoScenario",
		Str: "无法取得对应ID的Scenario",
	}
)
