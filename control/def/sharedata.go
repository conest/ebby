package def

import (
	"github.com/faiface/pixel/text"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ShareData : 控制中心共享数据
type ShareData struct {
	ebbyData
	UserData interface{}
}

// ShareData : 控制中心共享数据
type ebbyData struct {
	Resource resource
	Tool     tool
}

// 公共资源库，如 Atlas、Sprite
type resource struct {
	DebugAtlas *text.Atlas
}

// 公共工具库，如 logger
type tool struct {
	Display     DisplayController
	Config      *viper.Viper
	DebugLogger *text.Text
	Logger      *logrus.Logger
}
