package ebby

import (
	"github.com/conest/ebby/game"
)

// Init : 初始化
func defaultInitialFunc(c *game.Game) {
	c.GameData().Sys.Logger.Info("[game] Started")
}

// BeforeExit : 关闭前行为（保存数据等）
func defaultBeforeExitFunc(c *game.Game) {
	c.GameData().Sys.Logger.Info("[game] Ternimated")
}

// emptyFunc : 空函数
func emptyFunc(c *game.Game) {}

// emptyData : 空数据
var emptyData = new(struct{})
