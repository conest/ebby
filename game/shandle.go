package game

import (
	"ebby/game/def"
	"ebby/game/scene"
	"ebby/system"
	"errors"
	"fmt"

	"github.com/faiface/pixel/pixelgl"
	"github.com/spf13/viper"
)

// SceneMap : Scene 列表映射 map
type SceneMap map[string]*scene.Scene

// loadScenes : 加载 Scene 列表
func loadScenes(sm SceneMap, sdata *def.ShareData, config *viper.Viper) SceneMap {
	for _, ss := range sm {
		ss.SetConfig(config)
		ss.SetData(sdata)
	}
	return sm
}

// initScene : 初始化场景
func initScene(s *scene.Scene, r def.Request, win *pixelgl.Window) {
	if r.ResetData {
		s.ResetData()
		s.Initial(win)
	}
}

// checkScene : 检测对应的scene是否存在
func checkScene(ok bool, id string) {
	if !ok {
		str := fmt.Sprintf("id: [%s]", id)
		err := errors.New(str)
		system.CheckErr(err, "game", system.ErrorTable["NoScene"])
	}
}
