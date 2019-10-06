package def

import (
	"github.com/conest/ebby/game/sys"

	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// GameData : Game实例功能以及数据接口
type GameData struct {
	Sys
	Tool
	PublicData interface{}
}

// Sys : 系统功能
type Sys struct {
	Win     *pixelgl.Window
	Display sys.DisplayController
	Config  *viper.Viper
	Logger  *logrus.Logger
}

// Tool : 公共工具库，如 logger
type Tool struct {
	DebugAtlas  *text.Atlas
	DebugLogger *text.Text
}
