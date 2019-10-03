package system

import (
	"fmt"
)

type errStruct struct {
	ID  string
	Str string
}

// CheckErr : 检查错误并输出日志
func CheckErr(err error, where string, es errStruct) {
	if err != nil {
		log := NewLogger()
		es := fmt.Sprintf("[%s] <%s> %v", where, es.Str, err)
		log.Error(es)
		panic(es)
	}
}

// Err : 直接报出错误并输出日志
func Err(err string, where string) {
	log := NewLogger()
	es := fmt.Sprintf("[%s] %v", where, err)
	log.Error(es)
	panic(es)
}

// ErrorTable : 默认错误类型map
var ErrorTable = map[string]errStruct{
	"CreateWindow": errStruct{
		ID:  "CreateWindow",
		Str: "Window创建失败",
	},
	"ViperErr": errStruct{
		ID:  "ViperErr",
		Str: "Viper设置错误",
	},

	// FontLoadFile :
	"FontLoadFile": errStruct{
		ID:  "FontLoadFile",
		Str: "字体初始化失败",
	},

	// PictureLoadFile :
	"PictureLoadFile": errStruct{
		ID:  "PictureLoadFile",
		Str: "图片文件读取失败",
	},

	// NoScene :
	"NoScene": errStruct{
		ID:  "NoScene",
		Str: "无法取得对应ID的Scene",
	},

	// PictureGet :
	"PictureGet": errStruct{
		ID:  "PictureGet",
		Str: "Picture get error. No picture named ",
	},

	// SetBatch :
	"SetBatch": errStruct{
		ID:  "SetBatch",
		Str: "Set batch error. No picture named ",
	},

	// GetBatch :
	"GetBatch": errStruct{
		ID:  "GetBatch",
		Str: "Get batch error. No batch named ",
	},

	// GetFrame :
	"GetFrame": errStruct{
		ID:  "GetBatch",
		Str: "Get frame error. No frame named ",
	},

	// GLSLLoad :
	"GLSLLoad": errStruct{
		ID:  "GLSLLoad",
		Str: "GLSL load resource err",
	},
}
