package control

// Enter : 创建窗口并初始化控制中心
func Enter(sm ScenarioMap, fn *Functions, sd interface{}) {

	// 创建新的控制中心
	c := New(sm, sd)
	c.SetFunctions(fn)
	c.SetDebugLogger()
	c.Init()
	c.Run()
	c.BeforeExit()
}
