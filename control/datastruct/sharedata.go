package datastruct

import (
	"github.com/faiface/pixel/text"
	"github.com/sirupsen/logrus"
)

// ShareData : 控制中心共享数据
type ShareData struct {
	Resource resource
	Tool     tool
}

// 公共资源库，如 Atlas、Sprite
type resource struct {
	DebugAtlas *text.Atlas
}

// 公共工具库，如 logger
type tool struct {
	DebugLogger *text.Text
	Logger      *logrus.Logger
}
