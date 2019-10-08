package game

// Enter : 创建窗口并初始化控制中心
func Enter(sceneMap SceneMap, fn *ExFunctions, publicData interface{}) {

	game := New(sceneMap, publicData)
	game.SetFunctions(fn)
	game.Init()
	game.Run()
	game.BeforeExit()
}
