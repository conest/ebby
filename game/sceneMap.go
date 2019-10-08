package game

import (
	"errors"
	"fmt"

	"github.com/conest/ebby/game/def"
	"github.com/conest/ebby/game/scene"
	"github.com/conest/ebby/system"
)

// SceneMap : Scene 列表映射 map
type SceneMap map[string]*scene.Scene

// loadScenes : 加载 Scene 列表
func loadScenes(sceneMap SceneMap, gamedata *def.GameData) SceneMap {
	for _, s := range sceneMap {
		s.SetGameData(gamedata)
		s.ResetInstanceData()
	}
	return sceneMap
}

// initScene : 初始化场景
func initScene(s *scene.Scene, r def.Request) {
	if r.ResetData {
		s.ResetInstanceData()
		s.InitInstance()
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
