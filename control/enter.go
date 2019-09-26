package control

// Enter : 创建窗口并初始化控制中心
func Enter(scenarioMap ScenarioMap, fn *Functions, sData interface{}) {

	// 创建新的控制中心
	ctrl := CreateControl(scenarioMap, sData)
	ctrl.SetFunctions(fn)
	ctrl.SetDebugLogger()
	ctrl.Init()
	ctrl.Run()
	ctrl.BeforeExit()
}
