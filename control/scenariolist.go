package control

import (
	"campaign/control/datastruct"
	"campaign/scenario"
	"campaign/test/textscen"

	"github.com/spf13/viper"
)

// scenario 列表，手动添加
var sl = []sListStruct{
	{
		name:     "test",
		scenario: textscen.Scenario(),
	},
}

type sListStruct struct {
	name     string
	scenario *scenario.Scenario
}

// ScenarioListMap : Scenario 列表映射 map
type ScenarioListMap map[string]*scenario.Scenario

// ScenarioList : 加载 Scenario 列表
func ScenarioList(sdata *datastruct.ShareData, config *viper.Viper) ScenarioListMap {

	sm := make(ScenarioListMap)

	for _, ss := range sl {
		ss.scenario.SetConfig(config)
		ss.scenario.SetSData(sdata)
		sm[ss.name] = ss.scenario
	}

	return sm
}
