package ebby

import (
	"ebby/common/logger"
	"ebby/control"
)

// Init : 初始化
func defaultInitialFunc(c *control.Control) {
	c.SData().Tool.Logger = logger.New()
	c.SData().Tool.Logger.Info("[Control Center] Started")
}

// BeforeExit : 关闭前行为（保存数据等）
func defaultBeforeExitFunc(c *control.Control) {
	c.SData().Tool.Logger.Info("[control] Ternimated")
}

// emptyFunc : 空函数
func emptyFunc(c *control.Control) {

}
