package control

import (
	"ebby/control/scenario"
	"ebby/control/def"
	"ebby/errdef"
	"errors"
	"fmt"

	"github.com/faiface/pixel/pixelgl"
	"github.com/spf13/viper"
)

// ScenarioMap : Scenario 列表映射 map
type ScenarioMap map[string]*scenario.Scenario

// loadScenarios : 加载 Scenario 列表
func loadScenarios(sm ScenarioMap, sdata *def.ShareData, config *viper.Viper) ScenarioMap {
	for _, ss := range sm {
		ss.SetConfig(config)
		ss.SetData(sdata)
	}
	return sm
}

// initScenario : 初始化场景
func initScenario(s *scenario.Scenario, r def.Request, win *pixelgl.Window) {
	if r.ResetData {
		s.ResetData()
		s.Initial(win)
	}
}

// checkScenario : 检测对应的scenario是否存在
func checkScenario(ok bool, id string) {
	if !ok {
		str := fmt.Sprintf("id: [%s]", id)
		err := errors.New(str)
		errdef.CheckErr(err, "control", errdef.NoScenario)
	}
}
