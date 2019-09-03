package control

import (
	"ebby/common/logger"
	"ebby/control/scenario"
	"ebby/control/strdef"
	"ebby/errdef"
	"fmt"

	"github.com/faiface/pixel/pixelgl"
	"github.com/spf13/viper"
)

// ScenarioMap : Scenario 列表映射 map
type ScenarioMap map[string]*scenario.Scenario

// loadScenarios : 加载 Scenario 列表
func loadScenarios(sm ScenarioMap, sdata *strdef.ShareData, config *viper.Viper) ScenarioMap {
	for _, ss := range sm {
		ss.SetConfig(config)
		ss.SetData(sdata)
	}
	return sm
}

// initScenario : 初始化场景
func initScenario(s *scenario.Scenario, r strdef.Request, win *pixelgl.Window) {
	if r.ResetData {
		s.ResetData()
		s.Initial(win)
	}
}

// checkScenario : 检测对应的scenario是否存在
func checkScenario(ok bool, id string) {
	if !ok {
		log := logger.New()
		errString := fmt.Sprintf("[control]<%s> %v", errdef.NoScenario.Str, id)
		log.Error(errString)
		panic(errString)
	}
}
