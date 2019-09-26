package errdef

import (
	"ebby/common/logger"
	"fmt"
)

type errStruct struct {
	ID  string
	Msg string
}

// CheckErr : 检查错误并输出日志
func CheckErr(err error, where string, es errStruct) {
	if err != nil {
		log := logger.New()
		log.Error(fmt.Sprintf("[%s] <%s> %v", where, es.Msg, err))
		panic(err)
	}
}

var (
	// CreateWindow :
	CreateWindow = errStruct{
		ID:  "CreateWindow",
		Msg: "Window创建失败",
	}

	// ViperErr :
	ViperErr = errStruct{
		ID:  "ViperErr",
		Msg: "Viper设置错误",
	}

	// FontLoadFile :
	FontLoadFile = errStruct{
		ID:  "FontLoadFile",
		Msg: "字体初始化失败",
	}

	// PictureLoadFile :
	PictureLoadFile = errStruct{
		ID:  "PictureLoadFile",
		Msg: "图片文件读取失败",
	}

	// NoScenario :
	NoScenario = errStruct{
		ID:  "NoScenario",
		Msg: "无法取得对应ID的Scenario",
	}
)
