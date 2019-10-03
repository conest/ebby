package game

// Enter : 创建窗口并初始化控制中心
<<<<<<< HEAD:control/enter.go
func Enter(scenarioMap ScenarioMap, fn *Functions, sData interface{}) {
=======
func Enter(sm SceneMap, fn *ExFunctions, sd interface{}) {
>>>>>>> fd8bd623f768a2abd9bd135df650099ef4017922:game/enter.go

	// 创建新的控制中心
	ctrl := CreateControl(scenarioMap, sData)
	ctrl.SetFunctions(fn)
	ctrl.SetDebugLogger()
	ctrl.Init()
	ctrl.Run()
	ctrl.BeforeExit()
}
