package control

import (
	"ebby/control/scenario"
	"ebby/control/strdef"
	"ebby/test/textscen"
	"ebby/test/textscen2"

	"github.com/spf13/viper"
)

// ScenarioListMap : Scenario 列表映射 map
type ScenarioListMap map[string]*scenario.Scenario

type sListStruct struct {
	name     string
	scenario *scenario.Scenario
}

// scenario 列表，手动添加
var sl = []sListStruct{
	{
		name:     "test",
		scenario: textscen.Scenario(),
	},
	{
		name:     "test2",
		scenario: textscen2.Scenario(),
	},
}

// ScenarioList : 加载 Scenario 列表
func ScenarioList(sdata *strdef.ShareData, config *viper.Viper) ScenarioListMap {

	sm := make(ScenarioListMap)

	for _, ss := range sl {
		ss.scenario.SetConfig(config)
		ss.scenario.SetData(sdata)
		sm[ss.name] = ss.scenario
	}

	return sm
}
