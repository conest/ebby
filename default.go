package ebby

import (
	"ebby/control"
)

// Init : 初始化
func defaultInitialFunc(c *control.Control) {
	c.SData().Tool.Logger.Info("[Control Center] Started")
}

// BeforeExit : 关闭前行为（保存数据等）
func defaultBeforeExitFunc(c *control.Control) {
	c.SData().Tool.Logger.Info("[control] Ternimated")
}

// emptyFunc : 空函数
func emptyFunc(c *control.Control) {}

// emptyData : 空数据
var emptyData = new(struct{})
